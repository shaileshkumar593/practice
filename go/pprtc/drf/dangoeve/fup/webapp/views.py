from django.shortcuts import render
from django.core.files.storage import FileSystemStorage

# Create your views here.

def simple_upload(request):
    if request.method == 'POST' and request.FILES['myfile']:
        myfnme = request.FILES['myfile']
        print("File uploaded is  : ",myfnme)
        print(myfnme.name)
        print(myfnme.size)
        fs = FileSystemStorage()
        filename = fs.save(myfnme.name, myfnme)
        uploaded_file_url = fs.url(filename)
        return render(request, 'index.html',{'uploaded_file_url': uploaded_file_url})
    return render(request, 'index.html')