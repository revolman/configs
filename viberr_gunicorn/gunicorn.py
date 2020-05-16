from multiprocessing import cpu_count
from os import environ

def max_workers():
    return cpu_count()


bind = '10.0.0.15:' + environ.get('PORT', '8000')
x_forwarded_for_header = 'X-Real-IP'
max_requests = 1000
worker_class = 'gevent'
workers = max_workers()
access_log_format = '%(t)s %({X-Real-IP}i)s %(l)s %(u)s "%(r)s" %(s)s %(b)s "%(f)s" "%(a)s" host:%(h)s'

env = {
    'DJANGO_SETTINGS_MODULE': 'Project1.settings'
}

reload = True
name = 'Project1'
