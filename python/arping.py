# -*- coding: utf-8 -*-

from subprocess import check_output
import sys


def arping(ipaddress="192.168.0.1"):
    out = check_output(["/usr/bin/arping", "-c 1", "-I", "enp1s0", ipaddress]).decode("utf-8")  # str
    result = out.split()
    for item in result:
        if ':' in item:
            print(item)


if __name__ == '__main__':
    if len(sys.argv) > 1:
        arping(sys.argv[1])
    else:
        arping()
