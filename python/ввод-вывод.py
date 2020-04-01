import sys

path = '/home/revolman/python/file1.txt'
counter = 1
file = open(path, 'w')
rfile = open(path, 'r')

while True:
    line = sys.stdin.readline()
    if not line:
        break
    file.write("%s: %s" % (counter, line))
    counter += 1
file.close()

print(rfile.read())
