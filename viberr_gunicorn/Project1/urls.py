from django.contrib import admin
from django.conf.urls import include, url

app_name = 'dbdata'

urlpatterns = [
    url(r'^admin/', admin.site.urls),
    url(r'^music/', include(('dbdata.urls', 'dbdata'), namespace='music')),
]
