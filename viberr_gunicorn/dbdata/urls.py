from django.urls import re_path
from . import views


urlpatterns = [
    # /music/
    re_path(r'^$', views.IndexView.as_view(), name='index'),
    # /music/$album_id/
    re_path(r'^(?P<pk>[0-9]+)/$', views.DetailView.as_view(), name='detail'),
    # /music/album/add
    re_path(r'album/add/$', views.AlbumCreate.as_view(), name='album-add'),
]
