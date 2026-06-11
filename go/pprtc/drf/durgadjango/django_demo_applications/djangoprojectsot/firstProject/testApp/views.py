from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def welcome(request):
    s='<h1>Hello Welcome to Django Classes..Django is Nursery Level Concept</h1>'
    return HttpResponse(s)
