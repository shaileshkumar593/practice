import os
os.environ.setdefault('DJANGO_SETTINGS_MODULE','djorm.settings')

import django
django.setup()


from webapp.models import *
from faker import Faker
from random import *


fake = Faker()
def populate(n):
    for i in range(n):
        fid = randint(1001,9999)
        fname = fake.name()
        femail = fake.company_email()
        fcompany = fake.address()
        fsalary = randint(200000,6000000)
        fAdd = fake.address()
        emp_record = Employee_Model.objects.get_or_create(Eid = fid,Name=fname,Email=femail,Company=fcompany,Salary = fsalary, Eaddress = fAdd)



populate(30)



