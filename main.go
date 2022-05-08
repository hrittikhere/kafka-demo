package main

import (
	"fmt"
	"github.com/hrittikhere/kafka-demo/cmd"
)

func main() {

	Name := gofakeit.Name()
	Address := gofakeit.Country()
	Gender := gofakeit.Gender()
	Email := gofakeit.Email()
	Phone := gofakeit.Phone()
	UserAgent := gofakeit.UserAgent()
	BirthdayMonth := gofakeit.MonthString()

	fmt.Println(Name, Address, Gender, Email, Phone, UserAgent, BirthdayMonth)

}
