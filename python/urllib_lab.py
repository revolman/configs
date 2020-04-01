# import urllib.request
import sys
import requests
import ssl
import urllib3


ctx = ssl.create_default_context()
ctx.check_hostname = False
ctx.verify_mode = ssl.CERT_NONE

urllib3.disable_warnings()
url = sys.argv[1]
try:
    response = requests.get(url, verify=False)
    print(f"Код {response.status_code}")
    if response.status_code in [200, 301]:
        response.encoding = 'utf-8'
        print(response.text)
    else:
        print(response.reason)

except requests.exceptions.ConnectionError as http_err:
    print(f"HttpError: {http_err}")



# urllib.request.Request(url)
# f = urllib.request.urlopen(url)
# print(f.status())
