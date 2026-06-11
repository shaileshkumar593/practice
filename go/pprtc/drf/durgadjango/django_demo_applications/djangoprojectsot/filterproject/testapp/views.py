from django.shortcuts import render
from testapp.models import FilterModel
# Create your views here.
def upper_view(request):
    data_list=FilterModel.objects.all()
    return render(request,'testapp/upper.html',{'data_list':data_list})

def lower_view(request):
    data_list=FilterModel.objects.all()
    return render(request,'testapp/lower.html',{'data_list':data_list})
