package main

import "fmt"

type programming struct {
	title     string
	salary    float64
	dephead   string
	employees []string
}

type design struct {
	title     string
	salary    float64
	dephead   string
	employees []string
}

type Department interface {
	EmployeesList() []string
	Title() string
	TotalEmployeesCount() int
	HeadOfDepartment() string
	NetSalary() float64
}

func (p programming) EmployeesList() string {

	lists := ""

	for _, value := range p.employees {
		lists += value + "\n"
	}

	return lists
}

func (d design) EmployeesList() string {

	lists := ""

	for _, value := range d.employees {
		lists += value + "\n"
	}

	return lists
}

func (p programming) Title() string {
	return p.title
}

func (d design) Title() string {
	return d.title
}

func (p programming) TotalEmployeesCount() int {
	return len(p.employees)
}

func (d design) TotalEmployeesCount() int {
	return len(d.employees)
}

func (p programming) HeadOfDepartment() string {
	return p.dephead
}

func (d design) HeadOfDepartment() string {
	return d.dephead
}

func (p programming) NetSalary() float64 {
	return p.salary * 0.75
}

func (d design) NetSalary() float64 {
	return d.salary * 0.5
}

func ShowDepartmentDetails(dep Department) string {

	return fmt.Sprint(" Department Employee Title -%v  \nHeadOfDepartment -  %v\nNetSalary of Employees in the Department - %v\nTotal Employees - %v\nList of employees in the Department: %v", dep.Title(), dep.HeadOfDepartment(), dep.NetSalary(), dep.TotalEmployeesCount(), dep.EmployeesList())

}

func main() {

	employe := []string{"Abebe", "belay", "Ayele"}
	designs := []string{"design 1", "design 2", "design 3"}
	programmings := programming{title: "programming title", salary: 2000.0, dephead: "CS", employees: employe}

	designss := design{title: "programming title", salary: 2000.0, dephead: "CS", employees: designs}

	//pro Department
	//des Department
	pro := []interface{}{programmings, designss}
	//pro = (programming)programmings
	// inter = []Department{}{programmings,design }

	// programming(pro) = programmings

	for _, i := range pro {
		ShowDepartmentDetails(i)

	}

}
