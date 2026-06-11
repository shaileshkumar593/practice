from django.shortcuts import render
from .models import Team_Model

# Create your views here.
def teamslist(request):
    items = Team_Model.objects.all()
    return render(request, 'webapp/index.html', {'items':items})
