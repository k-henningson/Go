package main

import (
	"html/template"
	"os"
)

// type User struct {
// 	Name string
// }

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := struct {
		Name string
		Age  int
	}{
		Name: "John Smith",
		Age:  99,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
