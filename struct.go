package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{4, 5}
	v.Scale(10)
	//fmt.Println(Area(v))
	fmt.Println(v.Area())
}