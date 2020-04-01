from TheBook.sqlAlchemy_lab import session, OperatingSystem

for os in session.query(OperatingSystem).filter(OperatingSystem.name.like('Lin%')):
    print(os)
