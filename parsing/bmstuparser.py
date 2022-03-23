from requests import Session
from bs4 import BeautifulSoup
import json


class EUParser:

    CONFIG_LINK = 'config.json'

    LOGIN_LINK = "https://proxy.bmstu.ru:8443/cas/login?service=https%3A%2F%2Fproxy." \
                 "bmstu.ru%3A8443%2Fcas%2Foauth2.0%2FcallbackAuthorize%3Fclient_name%3DCasOAuthClient%26client_id%3DEU"

    AUTH_LINK = "https://proxy.bmstu.ru:8443/cas/oauth2.0/authorize?state=nCyI7Zk0WZBpFiO" \
            "MoWVmsP4bt6ItlJ7p&response_type=code&approval_prompt=auto&client_id=EU&" \
           "redirect_uri=https%3A%2F%2Feu.bmstu.ru%2Fportal3%2Flogin1%2Fmail%3Fback%3Dhttps%3A%2F%2Feu.bmstu.ru%2F"

    def __init__(self):
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
        # TODO: Add parsing

if __name__ == "__main__":
    parser = EUParser()
    parser.parse()