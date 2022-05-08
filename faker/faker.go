package cmd

import (
	"github.com/brianvoe/gofakeit/v6"
	"encoding/json"
)

func UserDetails(){
	
	UserName := gofakeit.Name()
	Country := gofakeit.Country()
	Gender := gofakeit.Gender()
	Email := gofakeit.Email()
	Phone := gofakeit.Phone()
	UserAgent := gofakeit.UserAgent()
	BirthdayMonth := gofakeit.MonthString()

}