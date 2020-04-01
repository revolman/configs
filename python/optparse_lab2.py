import optparse
import os


def main():
    p = optparse.OptionParser(description="Python 'ls' command clone",
                              prog="pyls",
                              version="0.1a",
                              usage="%prog [directory]")
    options, args = p.parse_args()
    if len(args) == 1:
        path = args[0]
        for filename in os.listdir(path):
            print(filename)
    else:
        p.print_help()


if __name__ == '__main__':
    main()
