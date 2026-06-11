package repository

import (
	"api/model"
	"errors"
	"newapi/model"
	"sync"
)

type EmployeeRepository interface {
	Create(emp model.Employee) error
	GetAll() []model.Employee
	GetByID(id string) (model.Employee, bool)
	Update(id string, emp model.Employee) error
	Delete(id string) error
}

type employeeRepository struct {
	data  map[string]model.Employee
	mutex sync.RWMutex
}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{
		data: make(map[string]model.Employee),
	}
}