import socket
import sys


def check_server(address, port):
    s = socket.socket()
    print('Attemptiong to connect to %s on port %s.' % (address, port))
    try:
        s.connect((address, port))
        print('Connected to %s on port %s.' % (address, port))
        return True
    except socket.error as e:
        print("Connection to %s on port %s failed: %s" % (address, port, e))
        return False


if __name__ == '__main__':
    from optparse import OptionParser
    parser = OptionParser()

    parser.add_option("-a", "--address", dest="address", default="localhost",
                      help="ADDRESS for server", metavar="ADDRESS")
    parser.add_option("-p", "--port", dest="port", type=int, default=80,
                      help="PORT for server", metavar="PORT")
    parser.add_option("-l", "--log", dest="log_file",
                      help="Specify log file", metavar="/path/to/logfile")

    (options, args) = parser.parse_args()
    print('options: %s, args: %s' % (options, args))
    check = check_server(options.address, options.port)
    print("Health: %s" % check)
    if options.log_file:
        with open(options.log_file, 'a') as file:
            file.write("Health check passed: " + str(check) + '\n')

    sys.exit(not check)
