from django.shortcuts import render

# Create your views here.

def Home_view(request):
    return render(request, 'webapp/home.html')

def Sports_view(request):
    return render(request, 'webapp/sports.html')

def Shopping_view(request):
    return render(request, 'webapp/shopping.html')

def Academy_view(request):
    return render(request, 'webapp/Academic.html')


