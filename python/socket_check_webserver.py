# -*- coding: utf-8 -*-

import socket
import re
import sys


def check_webserver(address, port, resource):
    # Создание строки запроса
    global s
    if not resource.startswith('/'):
        resource = '/' + resource
    request_string = "GET %s HTTP/1.1\nHost: %s\n\n" % (resource, address)
    try:
        s = socket.socket()
        print("Подключение к %s на порт %s.\n" % (address, port))
        s.connect((address, port))
        s.send(request_string)
        response = s.recv(100)
        print("Получены первые 100 байт ответа:\n")
        print("%s\n" % response)
    except socket.error as e:
        print("Подключение к %s:%s неудачно. Причина: %s" % (address, port, e))
        return False
    finally:
        print("Закрытие соединения\n")
        s.close()
    lines = response.splitlines()
    try:
        version, status, message = re.split(r'\s+', lines[0], 2)
        print("Версия протокола: %s, код ответа: %s, сообщение: %s." % (version, status, message))
    except ValueError:
        print("Не получается разбить строку по пробелам")
        return False
    if status in ['200', '301']:
        print("Получен статус: %s. Сервер доступен!" % status)
        return True
    else:
        print("Получен статус: %s. Сервер НЕ доступен!" % status)


if __name__ == '__main__':
    from optparse import OptionParser
    parser = OptionParser(u'введи ключи.')
    parser.add_option("-a", "--address", dest="address", default="localhost",
                      help="ADDRESS of server", metavar="10.0.0.1")
    parser.add_option("-p", "--port", dest="port", default=80, type=int,
                      help="PORT of server", metavar="80")
    parser.add_option("-r", "--resource", dest="resource", default="/",
                      help="HTTP RESOURCE on server", metavar="/health.txt")
    (options, args) = parser.parse_args()
    check = check_webserver(options.address, options.port, options.resource)
    print("Функция вернула значение: %s" % check)
    sys.exit(not check)
