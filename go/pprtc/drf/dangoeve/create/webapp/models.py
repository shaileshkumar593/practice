from django.db import models

# Create your models here.
class Company(models.Model):
    CName = models.CharField(max_length=40)
    Clogo = models.FileField(null = 'True', blank='True')
    Ccity = models.CharField(max_length=40)


