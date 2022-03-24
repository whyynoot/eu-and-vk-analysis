from database import DataBase
from EUParser import EUParser
from VKParser import VKParser

# TODO: Main algorithm
def main():
    db = DataBase()
    #eu_parser = EUParser(db.session)
    #eu_parser.parse()
    vk_parser = VKParser(db.session)
    vk_parser.parse()

if __name__ == "__main__":
    main()
