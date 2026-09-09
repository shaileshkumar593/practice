package main

import "fmt"

type PrintMethod interface {
	PrintFunc()
}

type TeacherInfo struct {
	name string
	id   string
}

type StudentInfo struct {
	name string
	id   string
}

type College struct {
	StudentInfo
	TeacherInfo
}

func (s StudentInfo) PrintFunc() {
	fmt.Println("Student Name:", s.name)
	fmt.Println("Student ID:", s.id)
}

func main() {
	fmt.Println("Hello, World!")

	s := StudentInfo{
		name: "Shailesh",
		id:   "12345",
	}

	t := TeacherInfo{
		name: "Anil",
		id:   "5678",
	}

	c := College{
		StudentInfo: s,
		TeacherInfo: t,
	}

	// Direct method call
	s.PrintFunc()

	// Access embedded structs
	fmt.Println("Teacher:", c.TeacherInfo.name, c.TeacherInfo.id)
	fmt.Println("Student:", c.StudentInfo.name, c.StudentInfo.id)

	// Interface
	var p PrintMethod = s

	p.PrintFunc()
}
