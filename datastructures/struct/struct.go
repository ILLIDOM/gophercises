package main

import "fmt"

type Person struct {
	age  int
	name string
}

func (p Person) SayHello() {
	fmt.Printf("My name is %s and I am %d years old.\n", p.name, p.age)
}

func main() {
	p1 := Person{age: 20, name: "domi"}
	fmt.Println(p1)
	p1.SayHello()
}
