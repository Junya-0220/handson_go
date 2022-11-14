package main

import "fmt"

type Person struct {
	name string
	gender string
	age int
}

func main(){
	var maria *Person = newPerson("maria","female",29)
	fmt.Println(*maria)
	maria.transformGender()
	fmt.Println(*maria)
}

func newPerson(n,g string, a int) *Person {
	person := new(Person)
	person.name = n
	person.gender = g
	person.age = a
	return person
}

func (p *Person) transformGender() {
	if p.gender == "male"{
		p.gender = "female"
	}else if p.gender == "female"{
		p.gender = "male"
	}
}
