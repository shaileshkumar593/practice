from django.contrib import admin
from .models import Emp_Model


# Register your models here.


class Emp_Model_Admin(admin.ModelAdmin):
    list_display = ['id', 'EmpId', 'EmpName', 'EmpSal']


admin.site.register(Emp_Model, Emp_Model_Admin)

# admin.site.register(Emp_Model) this will register whole table with admin
