#!/bin/bash

yum -y update
yum -y install httpd

rm /etc/httpd/conf.d/welcome.conf
touch /etc/httpd/conf.d/hello.conf

{
cat > /etc/httpd/conf.d/hello.conf << EOF
<Directory /var/www/html/index.html>
AllowOverride None
Require all granted
</Directory>
EOF
}

myip=`curl http://169.254.169.254/latest/meta-data/local-ipv4`
echo "<h2>WebServer with IP: $myip</h2><br>Build bu Terraform" > /var/www/html/index.html

systemctl start httpd
systemctl enable httpd