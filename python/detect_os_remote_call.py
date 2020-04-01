import subprocess


machines = ["node1", "node3", "node4"]
cmd = "python3.7 ~/detecting_os.py"

for machine in machines:
    subprocess.Popen("ssh revolman@%s %s" % (machine, cmd), shell=True)
