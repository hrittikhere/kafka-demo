package cmd

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"log"
)

type Faker struct {
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Country       string `json:"country"`
	UserAgent     string `json:"user_agent"`
	BirthdayMonth string `json:"birthday_month"`
}

func GetFakeUser() []byte {

	UserName := gofakeit.Name()
	Country := gofakeit.Country()
	Gender := gofakeit.Gender()
	Email := gofakeit.Email()
	Phone := gofakeit.Phone()
	UserAgent := gofakeit.UserAgent()
	BirthdayMonth := gofakeit.MonthString()

	FakeUser := Faker{UserName, Gender, Email, Phone, Country, UserAgent, BirthdayMonth}

	FakeUserEncoded, err := json.Marshal(FakeUser)

	if err != nil {
		log.Fatal(err)
	}

	return FakeUserEncoded
}
