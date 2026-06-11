package main

import "fmt"

// *** Nested  structure

type Student struct {
	name   string
	branch string
	year   int
}

type Teacher struct {
	name     string
	subject  string
	expernc  int
	studetls Student
}

func main() {

	res := Teacher{
		name:     "Srinivasan",
		subject:  "Discrete",
		expernc:  13,
		studetls: Student{"kannan", "CSE", 2018},
	}

	fmt.Println("Details of the teacher")
	fmt.Println("Name of Teacher", res.name)
	fmt.Println("Subject ", res.subject)
	fmt.Println("Experience is ", res.expernc)
	fmt.Println("Student Detail")
	fmt.Println("Name of Student ", res.studetls.name)
	fmt.Println("Branch of Student ", res.studetls.branch)
	fmt.Println("Year of Student ", res.studetls.year)

}
