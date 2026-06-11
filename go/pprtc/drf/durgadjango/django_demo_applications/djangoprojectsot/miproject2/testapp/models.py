from django.db import models

# Create your models here.
class BasicModel(models.Model):
    f1=models.CharField(max_length=64)
    f2=models.CharField(max_length=64)
    f3=models.CharField(max_length=64)

class StandardModel(BasicModel):
    f4=models.CharField(max_length=64)
    f5=models.CharField(max_length=64)
