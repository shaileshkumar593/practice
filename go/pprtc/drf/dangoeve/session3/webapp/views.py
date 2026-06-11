from django.shortcuts import render


# Create your views here.
def Referesh_Cnt(request):
    refresh = request.session.get('hit', 0)
    refrsh_updt = refresh + 1
    request.session['hit'] = refrsh_updt
    return render(request, 'webapp/sess.html', {'refrsh_cnt':refrsh_updt})

