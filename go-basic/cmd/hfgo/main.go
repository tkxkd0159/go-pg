package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello")
}

type Student1 struct {
	p      Person // Has-a embedding
	school string
}

type Student2 struct {
	Person // Is-a embedding
	school string
}

func (p *Student2) greeting() { // Overriding
	fmt.Println("New Hello")
}

func main() {
	var s Student1
	s.p.greeting()

	var s2 Student2
	s2.Person.greeting() // Hello
	s2.greeting()        // New Hello
}
