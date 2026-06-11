from django.urls import path
from . import views

urlpatterns = [
    path('jinja/', views.DateView)
]