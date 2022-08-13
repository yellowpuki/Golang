package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (p *Account) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	//pers := Person{1, "Pavel"}
	pers := new(Person)
	pers.SetName("Pavel Timokhovich")

	//fmt.Printf("update person: %#v\n", pers)

	var acc = Account{
		Id:   1,
		Name: "tpavel",
		Person: Person{
			Id:   2,
			Name: "Pavel Timokhovich",
		},
	}

	acc.SetName("timokhovich.pavel")
	acc.Person.SetName("Test")

	//fmt.Printf("%#v \n", acc)

	sl := MySlice([]int{1, 2})
	sl.Add(5)
	fmt.Println(sl.Count(), sl)

	return
}
