from .models import  Employee_Model
from django import forms


class Emp_form(forms.ModelForm):
    class Meta:
        model = Employee_Model
        fields = [
            'Eid',
            'Ename',
            'Eadd',
            'Epict',
        ]