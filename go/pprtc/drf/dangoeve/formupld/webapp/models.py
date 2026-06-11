from django.db import models

# Create your models here.

class Employee_Model(models.Model):
    Eid = models.IntegerField()
    Ename = models.CharField(max_length=100)
    Eadd = models.CharField(max_length= 200)
    Epict =models.FileField(blank=True,null=True)
