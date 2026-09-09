package main

import "fmt"

type Student struct {
	Name string
	ID   int
}

type Teacher struct {
	Name string
	ID   int
}

type College struct {
	Student
	Teacher
}

func main() {
	c := College{
		Student: Student{
			Name: "Shailesh",
			ID:   101,
		},
		Teacher: Teacher{
			Name: "Anil",
			ID:   201,
		},
	}

	fmt.Println(c.Student.Name)
	fmt.Println(c.Student.ID)

	fmt.Println(c.Teacher.Name)
	fmt.Println(c.Teacher.ID)
}
