package exercise

import (
	"fmt"
)

type student struct {
	studentID string
	address   *string
}

func (student *student) changeAddress(newAddress string) {
	student.address = &newAddress
}

func (student *student) changeStudentID(newStudentID string) {
	student.studentID = newStudentID
}

func Exercise5() {
	student := student{
		studentID: "41635",
	}

	student.changeAddress("Krung Thep Maha Nakorn, Thailand")
	student.changeStudentID("99999")

	fmt.Println(student.studentID)
	fmt.Println(*student.address)
}
