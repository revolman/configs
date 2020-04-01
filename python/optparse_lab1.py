import optparse


def main():
    p = optparse.OptionParser()
    p.add_option("-s", "--sysadmin", dest="admin", default="Admin",
                 help="setting up admin name", metavar="<name>")
    options, args = p.parse_args()

    print(f'Hello, {options.admin}!')


if __name__ == '__main__':
    main()
