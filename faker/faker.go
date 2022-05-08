package cmd

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
)

type Faker struct {
	Name          string `json:"name"`
	Gender        string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Country       string `json:"country"`
	Company       string `json:"company"`
	UserAgent     string `json:"user_agent"`
	BirthdayMonth string `json:"birthday_month"`
}

// type connection struct {
// 	Clients []*client `json:"Clients"`
// }

func UserDetails() {

	UserName := gofakeit.Name()
	Country := gofakeit.Country()
	Gender := gofakeit.Gender()
	Email := gofakeit.Email()
	Phone := gofakeit.Phone()
	UserAgent := gofakeit.UserAgent()
	BirthdayMonth := gofakeit.MonthString()

	// create json Response

}
