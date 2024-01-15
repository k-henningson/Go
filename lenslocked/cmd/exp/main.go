package main

import (
	"errors"
	"fmt"
)

// type User struct {
// 	Name string
// 	Bio  string
// 	Age  int
// }

// func main() {
// 	t, err := template.ParseFiles("hello.gohtml")
// 	if err != nil {
// 		panic(err)
// 	}

// 	user := User{
// 		Name: "John Smith",
// 		Bio:  `<script>alert("Cross Site Scripting lesson");</script>`,
// 		Age:  123,
// 	}

// 	err = t.Execute(os.Stdout, user)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	err := CreateOrg()
	fmt.Println((err))
}

func Connect() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("create org: %w", err)
	}
	return nil
}
