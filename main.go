package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/chonlatee/simplerules/validation"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)

}

func main() {

	users := make([]validation.User, 100)

	for i := 0; i < 100; i++ {
		users[i] = validation.User{
			Name:     randStr(rand.Intn(30-5+1) + 5),
			Password: randStr(rand.Intn(30-5+1) + 5),
			Age:      rand.Intn(30-5+1) + 5,
		}
	}

	v := validation.NewUserValidator()
	v.Add(validation.UserNameMinLength(10))
	v.Add(validation.UserNameMaxLength(25))
	v.Add(validation.UserPasswordMinLength(8))
	v.Add(validation.UserPasswordMaxLength(20))
	v.Add(validation.UserAge(18))

	for id, val := range users {
		errs := v.ValidateUser(val)
		if len(errs) != 0 {
			log.Printf("id:%v err: %v", id, errs)
		}
	}

	u := v.SelectUser(users)

	for _, v := range u {
		log.Printf("%+v", v)
	}
}
