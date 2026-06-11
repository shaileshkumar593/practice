from django.contrib import admin
from .models import Emp_Model

# Register your models here.

class EmpAdmin(admin.ModelAdmin):
    list_display = ['Eid', 'ESalary', 'EDesignation', 'EAddrs']

admin.site.register(Emp_Model, EmpAdmin)
