from django.shortcuts import render

# Create your views here.
def index(request):
    return render(request,'newsapp/index.html')

def moviesinfo(request):
    head_msg='Latest Movie Information'
    msg1='Sonali slowly getting cured'
    msg2='Salman going to marriage soon '
    msg3='Narendra Modi is going to act in some movie'
    my_dict={'head_msg':head_msg,'msg1':msg1,'msg2':msg2,'msg3':msg3}
    return render(request,'newsapp/news.html',context=my_dict)

def sportsinfo(request):
    head_msg='Latest Sports Information'
    msg1='Kohli gave Left and Right to Brodman Record'
    msg2='India Performance not upto the mark in asian Games '
    msg3='First Gold acquired by China'
    my_dict={'head_msg':head_msg,'msg1':msg1,'msg2':msg2,'msg3':msg3}
    return render(request,'newsapp/news.html',context=my_dict)

def politicsinfo(request):
    head_msg='Latest Politics Information'
    msg1='In Telangana Elections will come just in 2 months'
    msg2='Achchedin AAgayaaa!!! '
    msg3='Kerala Money Gaya...Center wont accept and wont give!!!'
    my_dict={'head_msg':head_msg,'msg1':msg1,'msg2':msg2,'msg3':msg3}
    return render(request,'newsapp/news.html',context=my_dict)
