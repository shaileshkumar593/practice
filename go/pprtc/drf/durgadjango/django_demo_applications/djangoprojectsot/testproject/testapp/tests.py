from django.test import TestCase
from testapp.models import Employee

class EmployeeTestCase(TestCase):
    def setUp(self):
        print('setUp actitity')
        Employee.objects.create(eno=100,ename='durga',esal=1000,eaddr='hyd')
        Employee.objects.create(eno=200,ename='sunny',esal=2000,eaddr='mumbai')

    def test_employee_info(self):
        print('testing Employee Information')
        qs=Employee.objects.all()
        self.assertEqual(qs.count(),3)
        e1=Employee.objects.get(ename='durga')
        e2=Employee.objects.get(ename='sunny')
        self.assertEqual(e1.get_salary(),1000)
        self.assertEqual(e2.get_salary(),2000)
