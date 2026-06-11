from django.shortcuts import render
from django.shortcuts import render
import datetime


# Create your views here.

def DateView(request):
    dt=datetime.datetime.now()
    name="PYTHON DJANGO WEB FRAMEWORK..!!"
    dic={'dt':dt,'name':name}
    return render(request,'webapp/index.html',dic)

