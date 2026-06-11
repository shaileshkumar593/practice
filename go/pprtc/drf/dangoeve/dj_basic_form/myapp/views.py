from django.shortcuts import render
from .forms import EmpForm

# Create your views here.

def Basic_Form(request):
    print("We are learning Basics oifference between html forms and django forms")

    return render(request,'myapp/html_form.html')

def djforms(request):
    newform = EmpForm()
    return render(request, 'myapp/djform.html', {'sample_form': newform})