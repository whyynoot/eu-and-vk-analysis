import time
import vk
from vk import exceptions
import json
from database import Student, VKGroup, GroupsStudents

class VKParser:
    CONFIG_LINK = '../config.json'
    LANG = 'ru'
    API_VERSION = '5.131'

    def __init__(self, db_session):
        # Getting access token as API Key
        self.access_token = None
        self.groups = []

        self.get_config_info()

        # Authorization
        self.session = vk.Session(access_token=self.access_token)
        self.api = vk.API(self.session, v=self.API_VERSION, lang=self.LANG)

        self.db_session = db_session

    def get_config_info(self):
        try:
            with open(self.CONFIG_LINK) as json_data_file:
                data = json.load(json_data_file)
                self.access_token = data['vk']['access_token']
                for group in data['vk']['groupsids']:
                    self.groups.append(group)
        except Exception as e:
            raise Exception("Unable to get config data", e)

    def get_bmstu_groups_members(self):
        for group in self.groups:
            counter = 0
            response = dict(count=0)
            while counter <= int(response['count']):
                response = self.api.groups.getMembers(group_id=group, offset=counter, count=1000, fields="lists")
                for person in response['items']:
                    name = f'{person["last_name"]} {person["first_name"]}'
                    link = f'{person["id"]}'

                    # Searching in db for person with name {name}
                    student = self.check_for_person_in_db(name)
                    if student is not None:
                        print(f"Suggested VK {link} for {student.name} from db")
                        self.update_student_vk_link(student, link)
                self.db_session.commit()
                counter += 1000
                print(f'Students left to check {response["count"] - counter}')

    def get_students_groups(self):
        students = self.db_session.query(Student).filter(Student.vk_link is not None).all()

        for student in students:
            if student.vk_link is not None:
                self.save_student_groups(student)
                self.db_session.commit()
                time.sleep(5)

    def save_student_groups(self, student):
        counter = 0
        response = dict(count=0)
        while counter <= int(response['count']):
            try:
                response = self.api.users.getSubscriptions(user_id=student.vk_link, offset=counter, count=200, fields='activity', extended=False)
            except exceptions.VkAPIError as e:
                # Profile is private
                if str(e)[:2] == '30':
                    break
                else:
                    raise Exception("Unable to parse student groups", e)

            for group in response['items']:
                try:
                    self.save_group(group['id'], group['name'], group['activity'])
                    self.save_group_student_linking(student.id, group['id'])
                except:
                    # Subs for users starts
                    pass
            self.db_session.commit()
            counter += 200

    def save_group_student_linking(self, student_id, group_id):
        self.db_session.merge(GroupsStudents(student_id=student_id, group_id=group_id))

    def save_group(self, id, name, category):
        self.db_session.merge(VKGroup(id=id, name=name, category=category))

    def update_student_vk_link(self, student, link):
        self.db_session.query(Student).filter(Student.name == student.name).\
            update({"vk_link": link}, synchronize_session="fetch")

    def check_for_person_in_db(self, name):
        suggested_student = self.db_session.query(Student).filter(Student.name.contains(name)).first()
        self.db_session.commit()
        return suggested_student

    def parse(self):
        self.get_bmstu_groups_members()
        self.get_students_groups()