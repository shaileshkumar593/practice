from django import forms
from .models import Emp_Model

class Emp_Form(forms.ModelForm):
    class Meta:
        model = Emp_Model
        fields = '__all__'

