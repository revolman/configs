# -*- coding: utf-8 -*-

import http.client
import sys
import subprocess


def check_webserver(address, port, resource):
    # Создать соединение
    global conn
    if not resource.startswith('/'):
        resource = '/' + resource
    try:
        conn = http.client.HTTPConnection(address, port)
        print('Соединение установлено.\n')
        # Выполнить запрос
        conn.request('HEAD', resource)
        print('Запрос к %s выполнен.\n' % address)
        response = conn.getresponse()
        print('Статус: %s' % response.status)
    except OSError as e:
        print('Нет подключения: %s' % e)
        return False
    finally:
        conn.close()
        print("Http соединение успешно закрыто\n")

    if response.status in [200, 301]:
        return True
    else:
        return False


if __name__ == '__main__':
    from optparse import OptionParser
    parser = OptionParser(u'введи ключи.')
    parser.add_option("-a", "--address", dest="address", default="mywp.local",
                      help="ADDRESS of server", metavar="10.0.0.1")
    parser.add_option("-p", "--port", dest="port", default=84, type=int,
                      help="PORT of server", metavar="82")
    parser.add_option("-r", "--resource", dest="resource", default="/",
                      help="HTTP RESOURCE on server", metavar="/health.txt")
    (options, args) = parser.parse_args()
    check = check_webserver(options.address, options.port, options.resource)
    print("Функция вернула значение: %s\n" % check)

    if not check and options.address == "mywp.local" and options.port == 82:
        print("Так быть не должно! Ansible установит роль.")
        subprocess.call("ansible-playbook --inventory=/home/revolman/git/Ansible/hosts\
                        /home/revolman/git/Ansible/use-roles.yml", shell=True)
        print("\nANSIBLE DONE")
    else:
        sys.exit(not check)
