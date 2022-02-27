package exercise

import (
	"fmt"
	"main/exercise/model"
)

type User interface {
	UserType() string
}

func printUserType(user User) {
	fmt.Println(user.UserType())
}

func Exercise4() {
	normalUser := model.NormalUser{
		Name: "Abby",
		Account: model.Account{
			Username: "Abby",
			Password: "Abby12345",
		},
	}

	vipUser := model.VIPUser{
		Name:      "misterfocusth",
		Privilege: "KOD KOD VIP",
		Account: model.Account{
			Username: "misterfocusth",
			Password: "123456789",
		},
	}

	printUserType(normalUser)
	printUserType(vipUser)
	fmt.Println(normalUser.CheckPassword("12345"))
	fmt.Println(vipUser.CheckPassword("123456789"))
}
