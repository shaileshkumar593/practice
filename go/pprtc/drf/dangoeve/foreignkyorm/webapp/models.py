from django.db import models

# Create your models here.
class Team_Model(models.Model):
    tname = models.CharField(max_length=255)
    tcountry  = models.CharField(max_length=200)
    def __str__(self):
        return self.tname + "  "+ self.tcountry


class Player_Model(models.Model):
    Pname = models.CharField(max_length=255)
    team = models.ForeignKey(Team_Model, on_delete= models.CASCADE)

    def __str__(self):
        return self.Pname