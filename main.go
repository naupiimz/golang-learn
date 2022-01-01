package main

import (
	"fmt"
)

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (student *Student) graduate() {
	student.Name = student.Name + " ST"
}

func main() {
	student := Student{12, "yupoo", 3.14}
	fmt.Println(student)

	student.graduate()

	fmt.Println(student)
}
