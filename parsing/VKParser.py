import requests
import vk
import json


class VKParser:

    CONFIG_LINK = 'config.json'

    def __init__(self):
        # Getting access token as API Key
        self.access_token = None
        self.groups = []

        self.get_config_info()

        # Authorization
        self.session = vk.Session(access_token=self.access_token)
        self.api = vk.API(self.session, v='5.131')

    def get_config_info(self):
        try:
            with open(self.CONFIG_LINK) as json_data_file:
                data = json.load(json_data_file)
                self.access_token = data['vk']['access_token']
                for group in data['vk']['groups']:
                    self.groups.append(group)
        except Exception as e:
            raise Exception("Unable to get config data", e)

    def get_groups_members(self):
        # TODO: Update method to get members
        pass

    def get_person_groups(self):
        # TODO: Update method to get members
        pass


if __name__ == "__main__":
    parser = VKParser()
    parser.get_groups_members()
