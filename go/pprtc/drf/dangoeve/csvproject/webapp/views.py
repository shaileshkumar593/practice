from django.shortcuts import render
from django.http import HttpResponse
import csv


# Create your views here.
def Downld_view(request):
     return render(request, 'webapp/downld.html')


def Getfile_view(request):
    response = HttpResponse(content_type='taext/csv')
    response['Content-Disposition'] = 'attachment; filename = "Bigdata.csv"'
    writer = csv.writer(response)
    writer.writerow(['1001', 'Shailesh', 'Software Engg'])
    writer.writerow(['1002', 'Anand', 'Thomson', 'Networking', 'Testing'])
    writer.writerow(['1003', 'Smrity', 'Poddar', 'Lecturer'])

    return response
