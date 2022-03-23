from requests import Session
from bs4 import BeautifulSoup
import json

try:
    with open("config.json") as json_data_file:
        data = json.load(json_data_file)
        username = data['bmstu']['username']
        password = data['bmstu']['password']
except Exception as e:
    raise Exception("Unable to get config data", e)

proxy = {'https': f"https://{username}:{password}@proxy.bmstu.ru:8080"}

login_link = "https://proxy.bmstu.ru:8443/cas/login?service=https%3A%2F%2Fproxy.bmstu.ru%3A8443%2Fcas%2Foauth2.0%2FcallbackAuthorize%3Fclient_name%3DCasOAuthClient%26client_id%3DEU"

auth_link= "https://proxy.bmstu.ru:8443/cas/oauth2.0/authorize?state=nCyI7Zk0WZBpFiOMoWVmsP4bt6ItlJ7p&response_type=code&approval_prompt=auto&client_id=EU&redirect_uri=https%3A%2F%2Feu.bmstu.ru%2Fportal3%2Flogin1%2Fmail%3Fback%3Dhttps%3A%2F%2Feu.bmstu.ru%2F"

def get_login_token(session):
    try:
        response = session.get(login_link, allow_redirects=False, timeout=5)
        html = BeautifulSoup(response.text, features="lxml")
        token = html.find("input", {'name': 'execution'})['value']
        return token
    except Exception as e:
        raise Exception("Unable to get execution token", e)

def login(session, token):
    data = {
        'username': username,
        'password': password,
        'execution': token,
        '_eventId': 'submit',
        'geolocation': ''
    }
    try:
        response = session.post(login_link, data=data, allow_redirects=True)
    except Exception as e:
        raise Exception("Unable to get login token", e)
    try:
        response = session.get(auth_link, timeout=5, allow_redirects=False)
    except Exception as e:
        raise Exception("Unable to auth", e)
    try:
        response = session.request("GET", response.headers['Location'], proxies=proxy)
        if response.status_code == 200:
            return True
    except Exception as e:
        raise Exception("Unable to enter eu.bmstu.ru", e)


def authorize(session):
    print("Staring to login...")
    try:
        token = get_login_token(session)
        if login(session, token):
            print("Successfully logged in to eu.bmstu.ru")
    except Exception as e:
        raise e

def main():
    try:
        login_session = Session()
        authorize(login_session)
        parse(login_session)
    except Exception as e:
        print(str(e))

if __name__ == "__main__":
    main()