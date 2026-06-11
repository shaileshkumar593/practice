from django.shortcuts import render
from webapp.forms import Emp_Form
from webapp.models import Emp_Model
from django.http import HttpResponseRedirect


# Create your views here.

def ThankView(request):
    return render(request, 'webapp/thanks.html')


def Emp_View(request):
    if request.method == 'POST':
        form_obj = Emp_Form(request.POST)

        if form_obj.is_valid():
            form_obj.save(commit=True)
            return HttpResponseRedirect('/thanks')
    else:
        form_obj = Emp_Form()

    return render(request, 'webapp/index.html', {'myform':form_obj})


def Emp_Model_data(request):
    Emp_data = Emp_Model.objects.all()
    return  render(request,'webapp/data_disp.html', {'elist':Emp_data})

