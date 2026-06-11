package service

import (
	"newapi/model"
	"newapi/repository"
)

type EmployeeService interface {
	CreateEmployee(emp model.Employee) error
	GetEmployees() []model.Employee
	GetEmployee(id string) (model.Employee, bool)
	UpdateEmployee(id string, emp model.Employee) error
	DeleteEmployee(id string) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		repo: repo,
	}
}

func (s *employeeService) CreateEmployee(emp model.Employee) error {
	return s.repo.Create(emp)
}

func (s *employeeService) GetEmployees() []model.Employee {
	return s.repo.GetAll()
}

func (s *employeeService) GetEmployee(id string) (model.Employee, bool) {
	return s.repo.GetByID(id)
}

func (s *employeeService) UpdateEmployee(id string, emp model.Employee) error {
	return s.repo.Update(id, emp)
}

func (s *employeeService) DeleteEmployee(id string) error {
	return s.repo.Delete(id)
}
