from TheBook.sqlAlchemy_lab import OperatingSystem, session

for os in session.query(OperatingSystem):
    print(os)
