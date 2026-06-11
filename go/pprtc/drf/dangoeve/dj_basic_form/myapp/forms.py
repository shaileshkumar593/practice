from django import forms

class EmpForm(forms.Form):
    Name = forms.CharField()
    job  = forms.CharField()
    loc  = forms.CharField()
    salary = forms.IntegerField()
    newform = forms.Form()