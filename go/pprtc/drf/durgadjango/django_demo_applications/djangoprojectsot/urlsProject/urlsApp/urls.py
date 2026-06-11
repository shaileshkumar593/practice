from django.conf.urls import url
from urlsApp import views

urlpatterns = [
    url(r'^hydjobs/', views.hydjobsInfo),
    url(r'^punejobs/', views.punejobsInfo),
    url(r'^mumbaijobs/', views.mumbaijobsInfo),
    url(r'^noidajobs/', views.noidajobsInfo),
]
