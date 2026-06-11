from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def setcookie(request):
    response = HttpResponse("<h1> Say Hey Cookie Set...!!<h1>")
    response.set_cookie('user', 'Raju')
    return response

def getcookie(request):
    if request.session.test_cookie_worked():
        print("Hello Cookies are Working Fine ...!!!")
        user = request.COOKIES['user']
    return HttpResponse(user + 'Hey Get Cookie Success')


def delcookie(request):
    response = HttpResponse('<h1> Say hey Cookie Deleted ...!!!</h1>')
    response.delete_cookie('user')
    return response

