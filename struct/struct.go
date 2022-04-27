package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	var acc = Account{
		Id:   1,
		Name: "Pavel",
		Person: Person{
			Name:    "Павел",
			Address: "Москва",
		},
	}
	fmt.Printf("%#v\n", acc)

	acc.Owner = Person{2, "Pavel", "Moscow"}
	fmt.Printf("%#v\n", acc)

	fmt.Println(acc.Name)
	fmt.Println(acc.Person.Name)
}
