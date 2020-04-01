import paramiko


hostname = '192.168.0.156'
port = 22
username = 'revolman'
secret = "PassHer3"


if __name__ == "__main__":
    s = paramiko.SSHClient()
    s.load_system_host_keys()
    s.connect(hostname, port, username, secret)
    stdin, stdout, stderr = s.exec_command('sudo ifconfig')
    stdout_str = stdout.read().decode('utf-8')
    print(stdout_str)

    s.close()
