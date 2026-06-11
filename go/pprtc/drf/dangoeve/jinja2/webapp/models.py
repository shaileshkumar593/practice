from django.db import models

# Create your models here.

class Emp_Model (models.Model):
    EmpId = models.IntegerField()
    EmpName = models.CharField(max_length=30)
    EmpSal = models.IntegerField()
    EmpAdd = models.CharField(max_length=30)




