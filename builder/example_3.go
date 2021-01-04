package main

import "fmt"

// Functional Builder

//Employee struct
type Employee struct {
	name, company, address string
}

type handler func(employee *Employee)

//EmployeeBuilder struct
type EmployeeBuilder struct {
	actions []handler
}

//Called sets name of the employee
func (b *EmployeeBuilder) Called(value string) *EmployeeBuilder {
	b.actions = append(b.actions, func(e *Employee) {
		e.name = value
	})
	return b
}

//WorksFor sets company of the employee
func (b *EmployeeBuilder) WorksFor(value string) *EmployeeBuilder {
	b.actions = append(b.actions, func(e *Employee) {
		e.company = value
	})
	return b
}

//At sets address of the employee
func (b *EmployeeBuilder) At(value string) *EmployeeBuilder {
	b.actions = append(b.actions, func(e *Employee) {
		e.address = value
	})
	return b
}

//Build builds the employee object
func (b *EmployeeBuilder) Build() Employee {
	emp := Employee{}
	for _, a := range b.actions {
		a(&emp)
	}
	return emp
}

//NewEmployeeBuilder - constructor
func NewEmployeeBuilder() *EmployeeBuilder {
	return &EmployeeBuilder{}
}

//RunFunctionalBuilder example
func RunFunctionalBuilder() {
	e := NewEmployeeBuilder()
	employee := e.Called("Surya").WorksFor("IBM").At("Bangalore").Build()
	fmt.Println(employee)
}
