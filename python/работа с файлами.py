
# with open('/home/revolman/httpd.conf') as file:
#    print(file.read())

# try:
#     wfile = open('/home/revolman/httpd.conf', 'a')
#     wfile.write('kyky\n')
# finally:
#     wfile.close()
#
# try:
#     rfile = open('/home/revolman/httpd.conf', 'r')
# finally:
#     print(rfile.read())

path = '/home/revolman/python/test.txt'
# with open(path, 'w') as file:
#     file.writelines('%s\t' % i for i in range(10))
#
# file = open(path, 'r')
# print(file.read())


def myRange(r):
    i = 0
    while i < r:
        yield "%s\t" % i
        i += 1


if __name__ == '__main__':
    with open(path, 'w') as wfile:
        wfile.writelines(myRange(6))

    rfile = open(path, 'r')
    print(rfile.read())
