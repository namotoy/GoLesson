package main

import (
	"fmt"
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
	//var s string = "Hello world"
	//s = strings.Replace(s, "H", "X", 1)
	//fmt.Println(s)
	//fmt.Println(strings.Contains(s, "world"))
	//fmt.Println("Test\n" +
	//	"Test")
	//t, f := true, false
	//fmt.Printf("%T %v %t\n ", t, t, 1)
	//fmt.Printf("%T %v %t\n", f, f, 2)
	//var x int = 1
	//xx := float64(x)
	//fmt.Printf("%T %v %f\n", xx, xx, xx)
	//
	//var y float64 = 1.2
	//yy := int(y)
	//fmt.Printf("%T %v %d\n", yy, yy, yy)
	//
	//var s string = "15"
	//i, _ := strconv.Atoi(s)
	//fmt.Printf("%T %v", i, i)
	//var a [2]int
	//a[0] = 100
	//a[1] = 200
	//fmt.Println(a)
	//
	//var b [2]int = [2]int{100, 200}
	//fmt.Println(b)
	//n := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(n)
	//fmt.Println(n[2])
	//fmt.Println(n[2:5])
	//fmt.Println(n[:2])
	//fmt.Println(n[2:])
	//fmt.Println(n[:])
	//n[2] = 100
	//fmt.Println(n)
	//var board = [][]int{
	//	[]int{0, 1, 2},
	//	[]int{3, 4, 5},
	//	[]int{6, 7, 8},
	//}
	//fmt.Println(board)
	//n = append(n, 100, 200, 300, 400)
	//fmt.Println(n)

	//n := make([]int, 3, 5)
	//fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	//n = append(n, 0, 0)
	//fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	//n = append(n, 1, 2, 3, 4, 5, 6)
	//var c []int
	//fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)
	//c = make([]int, 5)
	//c = make([]int, 0, 5)
	//for i := 0; i < 5; i++ {
	//	c = append(c, i)
	//	fmt.Println(c)
	//}
	//fmt.Println(c)

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 10000
	fmt.Println(m2)
}
