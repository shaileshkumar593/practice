from django.shortcuts import render
from django.http import HttpResponse


# Create your views here.

def setcookie(request):
    response = HttpResponse("<h1 style='color:green'>Cookie is Set </h1>")
    response.set_cookie('user', 'Anurag')
    return response


def getcookie(request):
    user = request.COOKIES['user']
    return HttpResponse(user + '  cookie is successfully return ')


def delcookie(request):
    response = HttpResponse("<h1 style='color:Yellow'> Cookie Deleted </h1>")
    response.delete_cookie('user')
    return  response

def Referesh_Count(request):
    refresh_val = request.COOKIES.get('NEWCOOKIE', 0)
    refresh_cnt_updtd_val = int(refresh_val) + 1
    response = render(request, 'webapp/sess.html', {'refrsh_cnt': refresh_cnt_updtd_val})
    response.set_cookie('NEWCOOKIE', refresh_cnt_updtd_val)
    return response

    #hit = request.COOKIES.get('hit', 0)
    #newhit = int(hit) + 1
    #response = render(request, 'webapp/sess.html', {'hit': newhit})
    #response.set_cookie('hit', newhit)
    #return response



