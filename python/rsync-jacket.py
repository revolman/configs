from subprocess import call
import sys


source = "/home/revolman/python/dirs/dir1"
target = "/home/revolman/python/dirs/dir2"
rsync = "rsync"
arguments = "-a"
cmd = "%s %s %s %s" % (rsync, arguments, source, target)


def sync():
    ret = call(cmd, shell=True)
    if ret != 0:
        print("rsync failed")
        sys.exit(1)


sync()
