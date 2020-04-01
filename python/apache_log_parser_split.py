"""
ПОРЯДОК ИСПОЛЬЗОВАНИЯ:
apache_log_parser_split.py some_log_file
Этот сценарий принимает единственный аргумент командной строки: имя файла
журнала, который требуется проанализировать. Он анализирует содержимое файла
и генерирует отчет, содержащий перечень удаленных хостов и число байтов,
переданных каждому из них.
"""

import sys


def dictify_logline(line):
    '''возвращает словарь, содержащий информацию, извлеченную из
    комбинированного файла журнала
    В настоящее время нас интересуют только адреса удаленных хостов
    и количество переданных байтов, но для полноты картины мы
    добавили выборку кода состояния.
    '''
    split_line = line.split()
    return {'remote_host': split_line[0],
            'status': split_line[8],
            'bytes_sent': split_line[9],
            }


def generate_log_report(logfile):
    '''возвращает словарь в формате:
    remote_host=>[список числа переданных байтов]
    Эта функция принимает объект типа file, выполняет обход всех строк
    в файле и создает отчет о количестве байтов, переданных при каждом
    обращении удаленного хоста к веб серверу.
    '''
    report_dict = {}
    for line in logfile:
        line_dict = dictify_logline(line)
        print(line_dict)
        try:
            bytes_sent = int(line_dict['bytes_sent'])
        except ValueError:                  # полностью игнорировать непонятные нам ошибки
            continue
        report_dict.setdefault(line_dict['remote_host'],    # метод setdefault возвращает значение ключа или дефолтное
                 # значение. Ключ remote_host из словаря line_dict,
                 # дефолтное значение [](пустое). Соответственно будет получено значение split_line[0] или пустое.
                 # так выглядит вывод IP-адреса из обработанной строки в словарь report_dict
                                []).append(bytes_sent)       # метод append добавит в словарь report_dict количество байтов от этого адреса
    return report_dict


if __name__ == "__main__":
    if not len(sys.argv) > 1:
        print(__doc__)
        sys.exit(1)
    infile_name = sys.argv[1]
    try:
        infile = open(infile_name, 'r')
    except IOError:
        print("You must specify a valid file to parse")
        print(__doc__)
        sys.exit(1)
    log_report = generate_log_report(infile)
    print(log_report)
    infile.close()
