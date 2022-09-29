package main

import "fmt"

type MyInterface interface {
	Draw()
	Sketch()
}

type Demo struct {
}

func (d Demo) Draw() {
	fmt.Println("I am drawing")
}

func (d Demo) Sketch() {
	fmt.Println("I am sketchin")
}

type NewDemo struct{}

func (d NewDemo) Draw() {
	fmt.Println("I am drawing new")
}

func (d NewDemo) Sketch() {
	fmt.Println("I am sketching new")
}

func main() {
	d := Demo{}        // demo has the methods Draw and Sketch Implemented
	usingInterfaces(d) // beause Demo implements Draw and Sketch it can be used as a MyInterface

	nD := NewDemo{}
	usingInterfaces(nD)
}

func usingInterfaces(mI MyInterface) {
	mI.Draw()
	mI.Sketch()
}
