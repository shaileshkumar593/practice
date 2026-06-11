from django.shortcuts import render
from webapp import forms
#from .forms import EmpForm
# Create your views here.
def EmpView(request):
    form_obj  = forms.EmpForm()
    return render(request,'webapp/index.html', {'form':form_obj})
