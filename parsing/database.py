import sqlalchemy
from sqlalchemy import Column, ForeignKey, Integer, String, Text, Date, DateTime
from sqlalchemy.orm import relationship
from sqlalchemy.ext.declarative import declarative_base
import sqlalchemy_utils
import json
from sqlalchemy.orm import Session


class DataBase:
    CONFIG_FILE = 'config.json'
    DATABASE_NAME = 'euandvk'

    def __init__(self):
        self.password = None
        self.user = None
        self.server = None

        self.get_config_data()

        self.engine = sqlalchemy.create_engine(f'mysql+pymysql://{self.user}:{self.password}@{self.server}/{self.DATABASE_NAME}')

        self.init_db()
        self.session = Session(bind=self.engine)

    def init_db(self):
        if not sqlalchemy_utils.database_exists(self.engine.url):
            sqlalchemy_utils.create_database(self.engine.url)
            print(f"New Database Created at/with {self.engine}")
        else:
            print("Database was not created as it exists")

    def get_config_data(self):
        try:
            with open(self.CONFIG_FILE) as json_data_file:
                data = json.load(json_data_file)
                self.user = data['mysql']['user']
                self.password = data['mysql']['password']
                self.server = data['mysql']['server']
        except Exception as e:
            raise Exception("Unable to get config data", e)


Base = declarative_base()


class Student(Base):
    # Configuration
    __tablename__ = 'students'

    # Attributes
    id = Column(Integer,
                nullable=False,
                unique=True,
                primary_key=True,
                autoincrement=True
                )
    name = Column(Text, nullable=False)
    vk_link = Column(Text, nullable=True)
    student_group = Column(Text, nullable=False)

    # Relations
    marks = relationship("Marks", back_populates="student")
    vkgroups = relationship("GroupsStudents", back_populates="students")

    def __repr__(self):
        return f'{self.name} | {self.student_group} | {self.vk_link}'


class Marks(Base):
    # Configuration
    __tablename__ = 'marks'

    # Attributes
    student_id = Column(ForeignKey("students.id"), primary_key=True, nullable=False, unique=True)

    credit_1 = Column(Integer, nullable=True)
    credit_2 = Column(Integer, nullable=True)
    credit_3 = Column(Integer, nullable=True)
    credit_4 = Column(Integer, nullable=True)
    credit_5 = Column(Integer, nullable=True)
    credit_6 = Column(Integer, nullable=True)
    credit_7 = Column(Integer, nullable=True)
    credit_8 = Column(Integer, nullable=True)
    credit_9 = Column(Integer, nullable=True)
    credit_10 = Column(Integer, nullable=True)

    exam_1 = Column(Integer, nullable=True)
    exam_2 = Column(Integer, nullable=True)
    exam_3 = Column(Integer, nullable=True)
    exam_4 = Column(Integer, nullable=True)
    exam_5 = Column(Integer, nullable=True)
    exam_6 = Column(Integer, nullable=True)
    exam_7 = Column(Integer, nullable=True)
    exam_8 = Column(Integer, nullable=True)


    # Relations
    student = relationship("Student", back_populates="marks")


class VKGroup(Base):
    # Configuration
    __tablename__ = 'vkgroups'

    # Attributes
    id = Column(Integer,
                      nullable=False,
                      unique=True,
                      primary_key=True)
    name = Column(Text, nullable=False)
    category = Column(Text, nullable=True)
    link = Column(Text, nullable=False)

    # Relations
    students = relationship("GroupsStudents", back_populates="vkgroups")


class GroupsStudents(Base):
    # Configuration
    __tablename__ = 'GroupsStudents'

    group_id = Column(ForeignKey("vkgroups.id"), primary_key=True)
    student_id = Column(ForeignKey("students.id"), primary_key=True)

    vkgroups = relationship("VKGroup", back_populates="students")
    students = relationship("Student", back_populates="vkgroups")


if __name__ == "__main__":
    db = DataBase()
    Base.metadata.create_all(db.engine)
    print("Python classes moved to mysql Entities")

    # Tests, only executed from if __main__
    from sqlalchemy.orm import Session
    session = Session(bind=db.engine)
    session.add(Student(name='Вася Пупукин', student_group='СГН3-31Б'))
    session.add(VKGroup(id=1, name='Первая вк группа', link='ссылка на группу'))
    session.add(GroupsStudents(group_id=1, student_id=1))
    session.add(Marks(student_id=1, exam_1=5))
    session.commit()


