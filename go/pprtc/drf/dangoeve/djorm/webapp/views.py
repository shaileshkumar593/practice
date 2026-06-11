from django.shortcuts import render
from webapp.models import Employee_Model
# Create your views here.

def Emp_view(request):
    Emp_qs = Employee_Model.objects.filter(Name__iendswith='huffman	')
    print(type(Emp_qs))
    return  render(request, 'webapp/base.html',{'Employee':Emp_qs})

