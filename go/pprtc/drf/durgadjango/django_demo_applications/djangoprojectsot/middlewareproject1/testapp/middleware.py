class ExecutionFlowMiddleware(object):

    def __init__(self,get_response):
        self.get_response = get_response

    def __call__(self,request):
        print('This line printed at pre processing of request')
        response = self.get_response(request)
        print('This line printed at post processing of request')
        return response

from django.http import HttpResponse
class AppMaintenanceMiddleware(object):
    def __init__(self,get_response):
        self.get_response=get_response

    def __call__(self,request):
        return HttpResponse('<h1>Currently Application under maintenance...Please try after 2 days !!!</h1>')

class ErrorMessageMiddleware(object):
    def __init__(self,get_response):
        self.get_response=get_response

    def __call__(self,request):
        response=self.get_response(request)
        return response

    def process_exception(self,request,exception):
        s1='<h1>Currently we are facing some technical problems, Plz try after some time </h1><hr>'
        s2='<h2>Raised Exception:{}</h2>'.format(exception.__class__.__name__)
        s3='<h2>Exception Description/Message:{}</h2>'.format(exception)
        return HttpResponse(s1+s2+s3)

class FirstMiddleware(object):

    def __init__(self,get_response):
        self.get_response = get_response

    def __call__(self,request):
        print('This line printed by First Middleware at pre processing of request')
        response = self.get_response(request)
        print('This line printed by First Middleware at post processing of request')
        return response

class SecondMiddleware(object):

    def __init__(self,get_response):
        self.get_response = get_response

    def __call__(self,request):
        print('This line printed by Second Middleware at pre processing of request')
        response = self.get_response(request)
        print('This line printed by Second Middleware at post processing of request')
        return response

class ThirdMiddleware(object):

    def __init__(self,get_response):
        self.get_response = get_response

    def __call__(self,request):
        print('This line printed by Third Middleware at pre processing of request')
        response = self.get_response(request)
        print('This line printed by Third Middleware at post processing of request')
        return response
