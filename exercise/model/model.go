package model

import (
	"fmt"
)

type NormalUser struct {
	Name string
	Account
}

type VIPUser struct {
	Name      string
	Privilege string
	Account
}

type Account struct {
	Username string
	Password string
}

func (account Account) CheckPassword(password string) bool {
	defer func() { fmt.Print("[CheckPassword Func]: Password Checked Result: ") }()
	return password == account.Password
}

func (user NormalUser) UserType() string {
	return "Normal"
}

func (user VIPUser) UserType() string {
	return "VIP"
}
