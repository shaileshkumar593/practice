from django.shortcuts import render
from django.http import HttpResponse, HttpResponseRedirect

# Create your views here.
def WelcomeView(request):
    if request.method == 'POST':
        request.session['username'] = request.POST['username']
        return HttpResponseRedirect('/home')
    return render(request, 'webapp/session.html')


def home(request):
    if request.session.has_key('username'):
        user = request.session['username']
        return render(request, 'webapp/home.html', {'user' : user})
    else:
        return  render(request, 'warning.html')


def bye(request):
    try:
        del request.session['username']
    except KeyError:
        pass
    return HttpResponse('You are loggeded out')


