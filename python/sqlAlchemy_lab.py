from sqlalchemy import create_engine
from sqlalchemy import Table, Column, Integer, Text, VARCHAR, MetaData
from sqlalchemy.orm import mapper
from sqlalchemy.orm import sessionmaker


engine = create_engine('sqlite:////home/revolman/python/inventory.db')
metadata = MetaData()
os_table = Table('inventory_operatingsystem', metadata,
                 Column('id', Integer, primary_key=True),
                 Column('name', VARCHAR(50)),
                 Column('description', Text()),)


class OperatingSystem(object):
    def __init__(self, name, description):
        self.name = name
        self.description = description

    def __repr__(self):
        return "<OperatingSystem('%s','%s')>" % (self.name, self.description)


mapper(OperatingSystem, os_table)
Session = sessionmaker(bind=engine, autoflush=True)
session = Session()
