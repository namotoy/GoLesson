package main

import (
	"fmt"
)

//const Pi = 3.14
//const (
//	Username = "test_user"
//	Password = "test_pass"
//)
//
//var (
//	i    int     = 1
//	f64  float64 = 1.2
//	s    string  = "test"
//	t, f bool    = true, false
//)
//
//func add(x, y int) (int, int) {
//	return x + y, x - y
//}
//
//func cal(price, item int) (result int) {
//	result = price * item
//	return result
//}

//func foo() {
//	xi := 1
//	var xf32 float32 = 1.2
//	xs := "test"
//	xt, xf := true, false
//	fmt.Println(xi, xf32, xs, xt, xf)
//	fmt.Printf("%T\n", xf32)
//	fmt.Printf("%T\n", xi)
//}

//func incrementGenerator() func() int {
//	x := 0
//	return func() int {
//		x++
//		return x
//	}
//}
//func circleArea(pi float64) func(radius float64) float64 {
//	return func(radius float64) float64 {
//		return pi * radius * radius
//	}
//}

//func by2(num int) string {
//	if num%2 == 0 {
//		return "ok"
//	} else {
//		return "no"
//	}
//}
//
//func getOsName() string {
//	return "macmac"
//}
//
//func boo() {
//	defer fmt.Println("world boo")
//	fmt.Println("goodbye boo")
//}
//
//func LoggingSettings(logFile string) {
//	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	multiLogFile := io.MultiWriter(os.Stdout, logfile)
//	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
//	log.SetOutput(multiLogFile)
//}

//func thirdPartyConnectDB() {
//	panic("Unable to connect database!")
//}
//
//func save() {
//	defer func() {
//		s := recover()
//		fmt.Println(s)
//	}()
//	thirdPartyConnectDB()
//}

func one(x *int) {
	*x = 1
}
func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)
	fmt.Println(&*&n)
	//fmt.Println(n)
	//fmt.Println(&n)
	//var p *int = &n
	//fmt.Println(p)
	//fmt.Println(*p)
	//save()
	//fmt.Println("OK!")

	//file, err := os.Open("./lesson.go")
	//if err != nil {
	//	log.Fatalln("Error!")
	//}
	//defer file.Close()
	//data := make([]byte, 100)
	//count, err := file.Read(data)
	//if err != nil {
	//	log.Fatalln("Error")
	//}
	//fmt.Println(count, string(data))

	//
	//LoggingSettings("test.log")
	//_, err := os.Open("affaafafa")
	//if err != nil {
	//	log.Fatalln("Exit", err)
	//}
	//log.Println("logging")
	//log.Printf("%T %v", "test", "test")
	//log.Fatalf("%T %v", "test", "test")
	//log.Fatalln("error!!")
	//fmt.Println("ok!")
	/*
		fmt.Println("run")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("success")
	*/
	/*
		/*boo()
		defer fmt.Println("world")
		fmt.Println("hello")
	*/

	//switch os := getOsName(); os {
	//case "mac":
	//	fmt.Println("Mac!!")
	//case "windows":
	//	fmt.Println("Windows!!")
	//default:
	//	fmt.Println("Default!!", os)
	//}
	//
	//t := time.Now()
	//fmt.Println(t.Hour())
	//switch {
	//case t.Hour() < 12:
	//	fmt.Println("Morning")
	//case t.Hour() < 17:
	//	fmt.Println("Afternoon")
	//}
	//l := []string{"python", "go", "java"}
	//for i := 0; i < len(l); i++ {
	//	fmt.Println(i, l[i])
	//}
	//for i, v := range l {
	//	fmt.Println(i, v)
	//}
	//for _, v := range l {
	//	fmt.Println(v)
	//}
	//
	//m := map[string]int{"apple": 100, "banana": 200}
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
	//for k := range m {
	//	fmt.Println(k)
	//}
	//for _, v := range m {
	//	fmt.Println(v)
	//}
	//for i := 0; i < 10; i++ {
	//	if i == 3 {
	//		fmt.Println("continue")
	//		continue
	//	}
	//
	//	if i > 5 {
	//		fmt.Println("break")
	//		break
	//	}
	//	fmt.Println(i)
	//}
	//sum := 1
	//for sum < 10 {
	//	sum += sum
	//	fmt.Println(sum)
	//}
	//fmt.Println(sum)

	//result := by2(10)
	//
	//if result == "ok" {
	//	fmt.Println("great")
	//}
	//if result2 := by2(10); result2 == "ok" {
	//	fmt.Println("great2")
	//}
	/*
		num := 9
		if num%2 == 0 {
			fmt.Println("by 2")
		} else if num%3 == 0 {
			fmt.Println("by 3")
		} else {
			fmt.Println("else")
		}
	*/

	x, y := 11, 20
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}
	if x == 10 || y == 10 {
		fmt.Println("||")
	}
	//counter := incrementGenerator()
	//fmt.Println(counter())
	//fmt.Println(counter())
	//fmt.Println(counter())
	//fmt.Println(counter())
	//
	//c1 := circleArea(3.14)
	//fmt.Println(c1(2))
	//
	//c2 := circleArea(3)
	//fmt.Println(c2(2))

	//r1, r2 := add(10, 30)
	//fmt.Println(r1, r2)
	//
	//r3 := cal(100, 3)
	//fmt.Println(r3)
	//
	//f := func(x int) {
	//	fmt.Println("inner func", x)
	//}
	//f(1)
	//
	//func(x int) {
	//	fmt.Println("inner func", x)
	//}(1)
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

	//m := map[string]int{"apple": 100, "banana": 200}
	//fmt.Println(m)
	//fmt.Println(m["apple"])
	//m["banana"] = 300
	//fmt.Println(m)
	//m["new"] = 500
	//fmt.Println(m)
	//
	//fmt.Println(m["nothing"])
	//
	//v, ok := m["apple"]
	//fmt.Println(v, ok)
	//
	//v2, ok2 := m["nothing"]
	//fmt.Println(v2, ok2)
	//
	//m2 := make(map[string]int)
	//m2["pc"] = 10000
	//fmt.Println(m2)

	//b := []byte{72, 73}
	//fmt.Println(b)
	//fmt.Println(string(b))
}
