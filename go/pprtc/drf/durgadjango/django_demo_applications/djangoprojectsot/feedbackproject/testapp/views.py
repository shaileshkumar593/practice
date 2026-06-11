from django.shortcuts import render
from . import forms

# Create your views here.
def thankyou_view(request):
    return render(request,'testapp/thankyou.html')

def feedback_view(request):
    if request.method=='GET':
        form=forms.FeedBackForm()

    if request.method=='POST':
        form=forms.FeedBackForm(request.POST)
        if form.is_valid():
            print('Validation completed...printing feedback info')
            print('Student Name:',form.cleaned_data['name'])
            print('Student Rollno:',form.cleaned_data['rollno'])
            print('Student Email:',form.cleaned_data['email'])
            print('Student Feedback:',form.cleaned_data['feedback'])
    return render(request,'testapp/feedback.html',{'form':form})
