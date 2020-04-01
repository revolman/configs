import os
import optparse


def main():
    p = optparse.OptionParser(description="Python 'ls' command clone",
                              prog="pyls",
                              version="0.1a",
                              usage="%prog [directory]")
    p.add_option("--verbose", "-v", default=False, action="store_true",
                 help="enables verbose output")

    opts, args = p.parse_args()
    if len(args) == 1:
        if opts.verbose:
            print("Verbose mode enabled")
        path = args[0]
        for filename in os.listdir(path):
            if opts.verbose:
                print(f"Filename: {filename}")
            else:
                print(filename)
    else:
        p.print_help()


if __name__ == '__main__':
    main()
