package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) add() int {
	return v.X + v.Y
}

func (v Vertex) multiply() {
	v.X = v.X + 1
	v.Y = v.Y + 1

}

func add1(v Vertex) int {
	return v.X + v.Y
}

func multiply1(v *Vertex) {
	v.X = v.X + 1
	v.Y = v.Y + 1

}

func main() {
	v := Vertex{2, 3}
	fmt.Println("Hello World!! ")
	multiply1(v)
	fmt.Println(add1(v))

}
