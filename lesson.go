package main

import (
	"fmt"
	"strings"
)

const Pi = 3.14
const (
	Username = "test_user"
	Password = "test_pass"
)

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	xi := 1
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)

}
func main() {
	/*var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type=%T value=%v", u8, u8)
	fmt.Println(Pi, Username, Password)
	fmt.Println(i, f64, s, t, f)
	foo()
	fmt.Println("1+1=", 1+1)
	x := 0
	fmt.Println(x)
	*/
	//fmt.Println(string("Hello world"[0]))
	var s string = "Hello world"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "world"))
	fmt.Println("Test\n" +
		"Test")

}
