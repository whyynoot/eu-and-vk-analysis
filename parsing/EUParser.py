from requests import Session
from bs4 import BeautifulSoup
import json
from database import Student, Marks

class EUParser:

    CONFIG_LINK = '../config.json'

    LOGIN_LINK = "https://proxy.bmstu.ru:8443/cas/login?service=https%3A%2F%2Fproxy." \
                 "bmstu.ru%3A8443%2Fcas%2Foauth2.0%2FcallbackAuthorize%3Fclient_name%3DCasOAuthClient%26client_id%3DEU"

    AUTH_LINK = "https://proxy.bmstu.ru:8443/cas/oauth2.0/authorize?state=nCyI7Zk0WZBpFiO" \
            "MoWVmsP4bt6ItlJ7p&response_type=code&approval_prompt=auto&client_id=EU&" \
           "redirect_uri=https%3A%2F%2Feu.bmstu.ru%2Fportal3%2Flogin1%2Fmail%3Fback%3Dhttps%3A%2F%2Feu.bmstu.ru%2F"

    #CREDIT_VARIANTS = ['НА', '']
    def __init__(self, db_session=None):
        # Getting user and password for proxy.bmstu.ru
        self.username = None
        self.password = None
        self.get_user_password()

        # Setting proxy
        self.proxy = {'https': f"https://{self.username}:{self.password}@proxy.bmstu.ru:8080"}

        # Requesting new requests session
        self.session = Session()

        self.logged_in = False

        self.login()

        self.db_session = db_session

    def get_user_password(self):
        try:
            with open(self.CONFIG_LINK) as json_data_file:
                data = json.load(json_data_file)
                self.username = data['bmstu']['username']
                self.password = data['bmstu']['password']
        except Exception as e:
            raise Exception("Unable to get config data", e)

    def login(self):
        try:
            print("Start logging in...")
            token = self.get_login_token()
            if self.authorize(token):
                print("Successfully logged in to eu.bmstu.ru")
                self.logged_in = True
        except Exception as e:
            raise Exception("Unable to login", e)

    def get_login_token(self):
        try:
            response = self.session.get(self.LOGIN_LINK, allow_redirects=False, timeout=5)
            html = BeautifulSoup(response.text, features="lxml")
            token = html.find("input", {'name': 'execution'})['value']
            return token
        except Exception as e:
            raise Exception("Unable to get execution token", e)

    def authorize(self, token):
        data = {
            'username': self.username,
            'password': self.password,
            'execution': token,
            '_eventId': 'submit',
            'geolocation': ''
        }
        try:
            response = self.session.post(self.LOGIN_LINK, data=data, allow_redirects=True)
        except Exception as e:
            raise Exception("Unable to get login token", e)
        try:
            response = self.session.get(self.AUTH_LINK, timeout=5, allow_redirects=False)
        except Exception as e:
            raise Exception("Unable to auth", e)
        try:
            response = self.session.request("GET", response.headers['Location'], proxies=self.proxy)
            if response.status_code == 200:
                return True
        except Exception as e:
            raise Exception("Unable to enter eu.bmstu.ru", e)

    def parse(self):
        if not self.logged_in:
            self.login()
        print("Start parsing...")
        try:
            response = self.session.get('https://eu.bmstu.ru/modules/session/?session_id=32', proxies=self.proxy)
            response.encoding = 'utf-8'
        except Exception as e:
            raise Exception("Unable to parse groups", e)

        if response.status_code == 200:
            html = BeautifulSoup(response.text, features='lxml')

            groups = html.find_all("a", {'name': 'sdlk'})

            #REMOVE TEST GROUP
            #groups = ['modules/session/group/201a4dfa-8610-11ea-8d72-005056960017/']

            print(f"Total groups found {len(groups)}")

            for group in groups:
                if ('СГН' in group.text):
                    try:
                        # REAL
                        print("Parsing " + group.text)
                        self.parse_students(group['href'])
                        # TEST
                        # self.parse_students(group)
                    except Exception as e:
                        print(f"Error with {group.text}", e)
                        pass
        else:
            raise Exception("Session's status code error")

    def parse_students(self, group_link):
        student_massive = []
        marks_massive = []
        try:
            response = self.session.get(f'https://eu.bmstu.ru/{group_link}', proxies=self.proxy)
            response.encoding = 'utf-8'
            #print(response.text)
            if response.status_code == 200:
                html = BeautifulSoup(response.text, features='lxml')
                table = iter(html.find('table').find_all('tr'))
                next(table)
                for row in table:
                    #student_uuid = row['student-uuid']
                    student_name = row.div.span.text
                    student_group = row.div.find_next('span').find_next('span').text
                    student_marks = row.find_all('td')
                    credits_massive = []
                    exams_massive = []
                    # 3 for skipping list number, name, number of document id
                    for mark in student_marks[3:]:
                        # Getting only credits
                        if int(mark['test-type']) == 2:
                            credits_massive.append(mark.span.text)
                        if int(mark['test-type']) == 1:
                            exams_massive.append(mark.span.text)
                    #print(student_name, student_group, credits_massive, exams_massive)
                    student_massive.append(Student(name=student_name, student_group=student_group))
                    marks_massive.append((self.convert_credits(credits_massive),
                                          self.convert_exams(exams_massive)))
                if self.db_session is not None:
                    self.db_session.add_all(student_massive)
                    self.db_session.commit()
                    marks = self.convert_marks(marks_massive, student_massive)
                    self.db_session.add_all(marks)
                    self.db_session.commit()
            else:
                raise Exception("Unable to parse students from group")
        except Exception as e:
            raise Exception("Unable to parse students:", e)

    @staticmethod
    def convert_marks(marks, students):
        db_marks = []
        for i in range(len(students)):
            db_mark = Marks(student_id=students[i].id)
            (credit, exam) = marks[i]
            for i in range(1, len(credit) + 1):
                setattr(db_mark, f"credit_{i}", credit[i - 1])
            for i in range(1, len(exam) + 1):
                setattr(db_mark, f"exam_{i}", exam[i - 1])
            db_marks.append(db_mark)
        return db_marks

    @staticmethod
    def convert_exams(exams):
        for i in range(len(exams)):
            if exams[i] == 'НА' or exams[i] == 'Я' or exams[i] == 'Неуд':
                exams[i] = 0
            elif exams[i] == 'Удов':
                exams[i] = 3
            elif exams[i] == 'Хор':
                exams[i] = 4
            elif exams[i] == 'Отл':
                exams[i] = 5
            else:
                exams[i] = 0
        return exams

    @staticmethod
    def convert_credits(student_credit):
        for i in range(len(student_credit)):
            if student_credit[i] == 'НА' or student_credit[i] == 'Я' or student_credit[i] == 'Нзч' or student_credit[i] == 'Неуд':
                student_credit[i] = 0
            elif student_credit[i] == 'Зчт':
                student_credit[i] = 1
            else:
                student_credit[i] = 0
        return student_credit