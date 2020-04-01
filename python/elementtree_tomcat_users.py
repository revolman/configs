from xml.etree import ElementTree as ET

if __name__ == '__main__':
    xmlfile = '/home/revolman/python/tomcat-users.xml'
    tomcat_users = ET.parse(xmlfile)
    for user in [e for e in tomcat_users.findall('./user') if
                 e.get('name') == 'tomcat']:
        print(user.attrib)
