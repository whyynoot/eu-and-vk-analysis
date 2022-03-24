from database import DataBase
from EUParser import EUParser

# TODO: Main algorithm
def main():
    db = DataBase()
    euparser = EUParser(db.session)
    euparser.parse()

if __name__ == "__main__":
    main()
