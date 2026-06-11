from django import forms

class EmpForm(forms.Form):
    Name = forms.CharField()
    job = forms.CharField()
    Loc = forms.CharField()
    Salary = forms.FloatField()


