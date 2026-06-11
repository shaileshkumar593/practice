from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def hydjobsInfo(request):
    s='<h1>Hyderabad Jobs Information</h1>'
    return HttpResponse(s)

def punejobsInfo(request):
    s='<h1>Pune Jobs Information</h1>'
    return HttpResponse(s)

def mumbaijobsInfo(request):
    s='<h1>HyMumbaiderabad Jobs Information</h1>'
    return HttpResponse(s)

def noidajobsInfo(request):
    s='<h1>Noida Jobs Information</h1>'
    return HttpResponse(s)
