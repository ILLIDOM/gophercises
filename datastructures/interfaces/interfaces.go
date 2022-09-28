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

func main() {
	d := Demo{}        // demo has the methods Draw and Sketch Implemented
	usingInterfaces(d) // beause Demo implements Draw and Sketch it can be used as a MyInterface
}

func usingInterfaces(mI MyInterface) {
	mI.Draw()
	mI.Sketch()
}
