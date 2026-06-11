from django.db import models

# Create your models here.
class Emp_Model(models.Model):
    Eid  = models.IntegerField()
    ESalary = models.IntegerField()
    EDesignation = models.CharField(max_length=80)
    EAddrs = models.CharField(max_length=120)
