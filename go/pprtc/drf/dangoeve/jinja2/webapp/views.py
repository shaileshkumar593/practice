from django.shortcuts import render
from .models import Emp_Model


# Create your views here.
def Emp_Model_View(request):
    emp_list = Emp_Model.objects.all()
    return render(request, 'webapp/index.html', {'elist': emp_list})
