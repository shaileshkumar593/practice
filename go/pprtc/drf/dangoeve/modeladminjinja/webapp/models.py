from django.db import models

# Create your models here.
class Emp_Model(models.Model):
    Empid = models.IntegerField()
    EmpName = models.CharField(max_length=50)
    Empsal = models.IntegerField()
    Empadd = models.CharField(max_length=50)


