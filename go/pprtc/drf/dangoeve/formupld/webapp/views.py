from django.shortcuts import render
from django.http import HttpResponseRedirect, HttpResponse
from webapp.forms import Emp_form

# Create your views here.
def Empview(request):
    return  render(request,'home.html')


def AddNewEmp(request):
    form = Emp_form(request.POST or None, request.FILES or None)
    if form.is_valid():
        x = form.save()
        #x.save()
        #return  HttpResponseRedirect('/') works when url contains '', views.XXXXX in urlpatterns
        # add code to check whether you want ot go to home or same page
        return HttpResponseRedirect('/home')
    return render(request,'modl.html', {'form':form})

