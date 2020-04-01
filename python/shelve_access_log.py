import shelve
from TheBook import apache_log_parser_regex

logfile = open('/home/revolman/python/access.log')
shelve_file = shelve.open('/home/revolman/python/access.s')

for line in logfile:
    d_line = apache_log_parser_regex.dictify_logline(line)
    shelve_file[d_line['remote_host']] = shelve_file.setdefault(d_line['remote_host'], 0) + int(d_line['bytes_sent'])
    # значение 127.0.0.1 = создать 127.0.0.1 со значением 0 + значение bytes_sent
    # каждая итерация будет приплюсовывать значение bytes_sent

for key in shelve_file:
    print(key, shelve_file[key])

shelve_file.close()
logfile.close()
