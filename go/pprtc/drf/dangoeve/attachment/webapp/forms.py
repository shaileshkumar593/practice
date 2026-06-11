from django import forms


class Attchment_Form(forms.Form):
    Name = forms.CharField(max_length=70)
    Subject = forms.CharField(max_length=200)
    Email = forms.CharField(max_length=80)
    Message = forms.CharField(max_length=300)
    Attchmnt = forms.FileField(required= False)
    # FileField can accept any things audio, video, attachment pdf, img.