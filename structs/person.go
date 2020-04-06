package main

import (
	"fmt"
)

type Contact struct {
	email   string
	zipcode int
}

type Person struct {
	firstname string
	lastname  string
	contact   Contact
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (p *Person) updateName(firstname string) {
	p.firstname = firstname
}
