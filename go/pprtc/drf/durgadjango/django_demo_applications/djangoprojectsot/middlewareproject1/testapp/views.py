from django.shortcuts import render
from django.http import HttpResponse
# Create your views here.
def welcome_view(request):
    print('This line printed by view function while processing request')
    return HttpResponse('<h1>Custom Middleare Demo</h1>')
