from django.shortcuts import render
from django.http import HttpResponse, HttpResponseRedirect
from django.urls import reverse


# Create your views here.
def Home(request):
    return HttpResponse("<a href = '/hi'>Click Here</a>")


def MyView(request):

    return HttpResponseRedirect(reverse("Bye"))


def ByeView(request):
    return HttpResponse("Good Bye ......!")


