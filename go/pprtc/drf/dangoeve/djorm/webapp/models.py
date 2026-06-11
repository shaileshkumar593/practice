from django.db import models

# Create your models here.
class Employee_Model(models.Model):
    Eid = models.IntegerField()
    Name = models.CharField(max_length=100)
    Email = models.EmailField(max_length=250)
    Company = models.CharField(max_length=100)
    Salary = models.FloatField()
    Eaddress = models.CharField(max_length=200)