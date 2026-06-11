from django.shortcuts import render
from django.core.files.storage import FileSystemStorage

# Create your views here.

def simple_upload(request):
    if request.method == 'POST' and request.FILES['myfile']:
        myfile = request.FILES['myfile']
        print('Name of the file is :', myfile.name)
        print('Size of the file is : ',myfile.size)

        fs = FileSystemStorage()
        filename = fs.save(myfile.name, myfile)
        uploaded_file_url = fs.url(filename)
        return render(request,'/webapp/fileupld.html', {'uploaded_file_url': uploaded_file_url})
    return (request, '/webapp/fileupld.html')





