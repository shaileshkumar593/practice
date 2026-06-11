from django.db import models

# Create your models here.
# class ContactInfo1(models.Model):
# 	name=models.CharField(max_length=64)
# 	email=models.EmailField()
# 	address=models.CharField(max_length=264)
#
# class Student1(ContactInfo1):
# 	rollno=models.IntegerField()
# 	marks=models.IntegerField()
#
# class Teacher1(ContactInfo1):
# 	subject=models.CharField(max_length=264)
# 	salary=models.FloatField()
class BasicModel(models.Model):
    f1=models.CharField(max_length=64)
    f2=models.CharField(max_length=64)
    f3=models.CharField(max_length=64)

class StandardModel(BasicModel):
    f4=models.CharField(max_length=64)
    f5=models.CharField(max_length=64)
