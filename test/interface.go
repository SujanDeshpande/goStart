package main

import "fmt"

type I interface {
	test()
}

type S struct {
	X int
	Y int
}

func (s *S) newStruct() *S {
	return &S{5, 6}
}

func (s *S) add() int {
	return s.X + s.Y
}

func (s S) test() {
	fmt.Println("In Interface method")
}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	fmt.Println("Hello World!")
	s := S{2, 3}
	fmt.Println(s)
	u := s.newStruct()
	fmt.Println(u)
	j := u.add()
	fmt.Println(j)
	i = u
	i.test()
	describe(i)
}
