from database import DataBase, Base
from EUParser import EUParser
from VKParser import VKParser

# TODO: Main algorithm
def main():
    db = DataBase()
    Base.metadata.create_all(db.engine)
    eu_parser = EUParser(db.session)
    eu_parser.parse()
    vk_parser = VKParser(db.session)
    vk_parser.parse()

if __name__ == "__main__":
    main()
