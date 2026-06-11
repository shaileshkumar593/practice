from .models import Company
from django import forms

class Company_Form(forms.ModelForm):
    class Meta:
        model = Company
        fields = [
            'CName',
            'Clogo',
            'Ccity',
        ]
