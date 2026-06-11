from django.db import models

# Create your models here.

class Person_Model(models.Model):
    Person_name = models.CharField(max_length=220)
    def __str__(self):
        return self.Person_name

class Uniqueid_Model(models.Model):
    Adhar_No = models.OneToOneField(Person_Model, on_delete=models.CASCADE)
    SSN  = models.OneToOneField(Person_Model, on_delete=models.CASCADE)
    state_name = models.CharField(max_length=150)
    def __str__(self):
        return (self.Adhar_No + " " + self.SSN)