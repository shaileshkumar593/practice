from django.shortcuts import render
from django.http import HttpResponseRedirect, HttpResponse
from django.core.mail import EmailMessage
from django.conf import settings
from .forms import Attchment_Form

# Create your views here.
def MailAttachmnt(request):
    if request.method == 'POST':
        form = Attchment_Form(request.POST, request.FILES)
        if form.is_valid():
            Name = form.cleaned_data['Name']
            Sububject = form.cleaned_data['Subject']
            Email = form.cleaned_data['Email']
            Message = form.cleaned_data['Message']
            Attachmnt = request.FILES['Attchmnt']

            try:
                mail  = EmailMessage('got Email from '+ str(Email),
                "Name : " + str(Name)+"\n"
                "Subject : " +str(Sububject)+"\n"
                "Email : " + str(Email)+"\n"
                "Message : " + str(Message),
                settings.EMAIL_HOST_USER,
                ['shaileshkumar593@gmail.com', 'shaileshk485@gmail.com'])
                mail.attach(Attachmnt.name, Attachmnt.read(), Attachmnt.content_type)
                mail.send(fail_silently=False)
                return  HttpResponseRedirect('GoodBye_View')
            except:
                return HttpResponse("Fail to Attach.....")
        else:
            form = Attchment_Form()
        return render(request, 'mail.html',{'form' : form })


def GoodBye_View(request):
    return render(request,'thanks.html')

