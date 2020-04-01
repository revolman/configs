import sys

num_arg = len(sys.argv) - 1

if num_arg == 0:
    sys.stderr.write('No args!\n')
else:
    print(sys.argv, "You typed in", num_arg, "arguments")
