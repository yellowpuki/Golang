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

type Address struct {
	host string
	port int
}

func main() {
	// var acc = Account{
	// 	Id:   1,
	// 	Name: "Pavel",
	// 	Person: Person{
	// 		Name:    "Павел",
	// 		Address: "Москва",
	// 	},
	// }
	// fmt.Printf("%#v\n", acc)
	//
	// acc.Owner = Person{2, "Pavel", "Moscow"}
	// fmt.Printf("%#v\n", acc)
	//
	// fmt.Println(acc.Name)
	// fmt.Println(acc.Person.Name)

	hits := make(map[Address]int)
	hits[Address{"go.dev", 443}]++

	seen := make(map[string]struct{})

	s := "hello"
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
	}

	fmt.Println(hits)
	fmt.Println(seen)
}
