from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def hydjobsinfo(request):
    s='<h1>Hyderabad Jobs Information</h1>'
    return HttpResponse(s)

def punejobsinfo(request):
    s='<h1>Pune Jobs Information</h1>'
    return HttpResponse(s)

def mumbaijobsinfo(request):
    s='<h1>Mumbai Jobs Information</h1>'
    return HttpResponse(s)

def chennaijobsinfo(request):
    s='<h1>Chennai Jobs Information</h1>'
    return HttpResponse(s)
