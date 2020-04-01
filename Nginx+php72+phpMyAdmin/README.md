# nginx+php72+phpMyAdmin

### Установка

Содержит конфигурацию Nginx для работы с phpMyAdmin.
Создана папка snippets, в которой есть конфигурация локейшена с phpMyAdmin.

Так же настроен fastcgi_pass для обработки php с помощью php-fpm.
```
sudo yum install http://rpms.remirepo.net/enterprise/remi-release-7.rpm
sudo yum install yum-utils
sudo yum-config-manager --enable remi-php72
sudo yum update
```
Найти нужные модули можно так:
```
sudo yum search php72 | egrep 'fpm|gd|mysql|memcache'
```
Далее установить все нужные пакеты:
```
sudo yum install php72 php72-php-fpm php72-php-gd php72-php-json php72-php-mbstring php72-php-mysqlnd php72-php-xml php72-php-xmlrpc php72-php-opcache
```
Конфиг FPM находится по адресу:
```
/etc/opt/remi/php72/php-fpm.d/www.conf
```

#### Установка phpMyAdmin
```
sudo yum install -y phpmyadmin
```
phpMyAdmin установиться в /usr/share/phpMyAdmin, нужно будет указать путь к этой директории в локейшене для доступа к phpMyAdmin.

### Генерация самоподписанных сертификатов

```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/nginx/ssl/mynginx.key -out /etc/nginx/ssl/mynginx.crt
```
Т.к. сертификат для HTTP, то необходимо указать доменное имя хоста в формате example.com

### Сжатие

gzip, в конфиге есть пример. Можно включать для всего сервера и отключать для конкретных локейшенов.

### Кэширование

Можно включить настройки кэша:
```
location ~* \.(jpeg|jpg|png|gif|ico|js|css)$ {
	expires 30d;
	add_header Pragma public;
	add_header Cache-Control public;
	add_header Vary Accept-Encoding;
}
```
Кэш будет храниться не более 30 дней. 
Заголовки Pragma и Cache-Controll (public | private |...) это способ сказать как хранить кэш, локально в браузере, или на публичном шлюзе.

### FastCGI cache

#### Контекст http
Здесь определяются параметры кэша: имя:время актуальности, структура папок, размер, место хранения, ключ. Чем уникальнее ключ, тем лучше, но занимает больше места.
```
fastcgi_cache_path /tmp/nginx_cache levels=1:2 keys_zone=microcache:10m max_size=500m;
fastcgi_cache_key "$scheme$request_method$host$request_uri";
add_header microcache $upstream_cache_status;
```

#### Контекст server:
Так выглядит включение описаного выше кэша в location ~ \.php$ {...}
```
fastcgi_cache	microcache;
fastcgi_cache_valid 200 60m;
fastcgi_index	index.php;

fastcgi_pass	127.0.0.1:9111;
```

#### Приоритеты location:
```
1. =			-	полное совпадение
2. ^~			-	префикс преимущества
3. ~ и *~		-	регулярные выражения (обычные и регистронезависимые)
4. no modifier		-	совпадение префикса
```

