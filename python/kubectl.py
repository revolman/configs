# Бесполезная обёртка для kubectl. Вспоминаю как пользоваться вызовом системных утилит и флагами.
import subprocess
import optparse


def kubeStatus(verb, resource, name, namespace):
    cmd = "/usr/local/bin/kubectl" + " " + verb + " " + resource + " " + name + "-n " + namespace
    process = subprocess.Popen(cmd, stdout=subprocess.PIPE, shell=True)
    counter = 0
    while process.poll() is None:
        line = process.stdout.readline()
        if line.decode('utf-8').__contains__("calico"):
            counter += 1

    print(counter)



if __name__ == '__main__':
    parser = optparse.OptionParser()
    parser.add_option("-v", "--verb", dest="verb", default="get",
                      help="Verb to execute get, edit, describe, log ...",
                      metavar="VERB")

    parser.add_option("-r", "--resource", dest="resource", default="pod",
                      help="Kubernetes resource: node, port, svc ...",
                      metavar="RESOURCE")

    parser.add_option("-n", "--name", dest="name", default="",
                      help="Name of Kubernetes resource",
                      metavar="NAME")

    parser.add_option("-N", "--namespace", dest="namespace", default="default",
                      help="Namespace",
                      metavar="NAMESPACE")

    (options, args) = parser.parse_args()
    # print(options.verb)

    kubeStatus(options.verb, options.resource, options.name, options.namespace)
