package repository

import (
	"errors"
	"sync"

	"newapi/model"
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

func (r *employeeRepository) Create(emp model.Employee) error {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.data[emp.ID]; exists {
		return errors.New("employee already exists")
	}

	r.data[emp.ID] = emp

	return nil
}

func (r *employeeRepository) GetAll() []model.Employee {

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var employees []model.Employee

	for _, emp := range r.data {
		employees = append(employees, emp)
	}

	return employees
}

func (r *employeeRepository) GetByID(id string) (model.Employee, bool) {

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	emp, exists := r.data[id]

	return emp, exists
}

func (r *employeeRepository) Update(id string, emp model.Employee) error {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.data[id]; !exists {
		return errors.New("employee not found")
	}

	r.data[id] = emp

	return nil
}

func (r *employeeRepository) Delete(id string) error {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.data[id]; !exists {
		return errors.New("employee not found")
	}

	delete(r.data, id)

	return nil
}
