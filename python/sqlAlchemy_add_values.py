from TheBook.sqlAlchemy_lab import session, OperatingSystem
import distro
import subprocess

uname = subprocess.Popen("uname -a", shell=True, stdout=subprocess.PIPE, encoding='utf-8')
out = uname.stdout.read()
linux = distro.linux_distribution()[0]
kernel = out.split()[2]

myos = OperatingSystem(name=linux, description=kernel)
session.add(myos)
session.commit()
