package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

//
// Concurrency
//

// goroutines - a thread
func hello() {
	time.Sleep(time.Second)
	fmt.Println("hello")
}

func threads() {
	go hello()
	fmt.Println("hi")
	hello()
}

// send and receive values through channels w/ channel operator <-
func multiply(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func count(c chan int) {
	for i := 1; i < 11; i++ {
		c <- i
	}
}

func channels() {
	ch := make(chan int)

	go multiply([]int{1, 2, 3}, ch)
	go multiply([]int{4, 5, 6}, ch)
	x, y := <-ch, <-ch

	fmt.Println(x, y)

	go count(ch)
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch, " ")
	}
}

func lottery(c chan int) {
	for i := 0; i < 10; i++ {
		c <- rand.Int()
	}
	close(c)
}

func closeChannels() {
	ch := make(chan int)
	go lottery(ch)
	for i := range ch {
		fmt.Println(i)
	}
}

func oneUntilQuit(c chan int) {
	for {
		fmt.Println(2)
		select {
		case <-c:
			return
			// default:
			// 	time.Sleep(time.Millisecond * 500)
			// 	fmt.Println(1)
		}
		fmt.Println(3)
	}
}

func requestOne() {
	quit := make(chan int)
	go oneUntilQuit(quit)
	time.Sleep(time.Second * 2)
	quit <- 0
	fmt.Println("ones done")
	time.Sleep(time.Second)
}

func main() {
	requestOne()
}

//
// Interfaces
//

// an interface type is a set of method signatures
type Phone interface {
	Call(string) string
	Throw()
}

type iPhone struct {
	broken bool
}

func (p *iPhone) Call(s string) string {
	if s == "George Costanza" && !p.broken {
		return "Believe it or not, George isn't at home."
	}
	return "Error"
}

func (p *iPhone) Throw() {
	p.broken = true
}

// type assertion (access interface type)
func typeAssertion() {
	var i interface{} = 1
	fmt.Println(i)
	j := i.(int)
	fmt.Println(j)
	j = 2
	fmt.Println(i)
}

// type switch
func typeSwitch(i interface{}) string {
	switch i.(type) {
	case int:
		return "is int"
	case complex128:
		return "is big complex"
	case complex64:
		return "is small complex"
	default:
		return "idk"
	}
}

//
// Methods
//

// no classes in go, functions instead defined as methods with a receiver argument
// it works while the receiver type is in the same package
func (t Triangle) Perimeter() int {
	return t.l1 + t.l2 + t.l3
}

// this is equivalent to
func Perimeter(t Triangle) int {
	return t.l1 + t.l2 + t.l3
}

// and can be any type
type X int

func (x X) Square() X {
	return x * x
}

// the following two functions receive a copy of t, not t itself. Changing the sides doesn't change og t
func badScale1(t Triangle) {
	t.l1 = t.l1 * 10
	t.l2 = t.l2 * 10
	t.l3 = t.l3 * 10
}

func (t Triangle) badScale2() {
	t.l1 = t.l1 * 10
	t.l2 = t.l2 * 10
	t.l3 = t.l3 * 10
}

// to change original t, have to give the function a pointer
func goodScale1(t *Triangle) {
	t.l1 = t.l1 * 10
	t.l2 = t.l2 * 10
	t.l3 = t.l3 * 10
}

// t.goodScale2() equivalent to (&t).goodScale2()
func (t *Triangle) goodScale2() {
	t.l1 = t.l1 * 10
	t.l2 = t.l2 * 10
	t.l3 = t.l3 * 10
}

//
// Basics
//

func functions(fn func(float64, float64) float64) func() int {
	fmt.Println(fn(5, 2))

	// create a closure function (references variables from outside body)
	var a int = 17
	return func() int { return a }
}

func maps() {
	var m map[string]int
	m = make(map[string]int)
	m["hallo"] = 1
	fmt.Println(m["hallo"])

	// map literal
	var m2 = map[int]string{1: "goodbye"}
	fmt.Println(m2)

	// using types
	var triangle = map[string]Triangle{
		"equilateral": {1, 1, 1},
		"isosceles":   {1, 1, 2},
		"scalene":     {1, 2, 3},
	}
	fmt.Println(triangle)
	delete(triangle, "scalene")
	fmt.Println(triangle)
	elem, ok := triangle["isosceles"]
	if ok {
		fmt.Println(elem)
	}
}

func arraysAndSlices() {
	// arrays have fixed length
	var a [10]complex128
	a[0] = 1

	// slices reference sections of an array. They don't store data and are apparently more common than arrays
	var s []complex128 = a[0:1]
	s[0] = 1i

	// array and slice literals
	al := [3]bool{true, true, false}
	al[0] = false
	sl := []bool{true, true, false}
	sl[0] = false

	// slice has length and capacity where capacity is determined by underlying array
	fmt.Println(cap(sl))
	fmt.Println(len(sl))
	sl = sl[:1]
	fmt.Println(len(sl))
	sl = sl[:3] // can extend the length of the slice
	fmt.Println(len(sl))
}

func moreSlices() {
	// use make to create dynamically sized arrays
	s := make([]int, 5)
	s[4] = 2
	s = append(s, 8)
	fmt.Println(s)

	// range loop
	for i, v := range s {
		fmt.Println(i, v)
	}

	// range loop ignore index
	for _, v := range s {
		fmt.Print(v)
	}
}

// struct - collection of fields
type Triangle struct {
	l1 int
	l2 int
	l3 int
}

func structure() {
	var t Triangle = Triangle{3, 4, 5}
	fmt.Println(t)
	fmt.Println(t.l1)

	// pointer to struct
	p := &t
	p.l1 = 4
	(*p).l2 = 4
	p.l3 = 4
	fmt.Println(t)
}

func moreStructureStuff() {
	// copying structs
	var t1 Triangle = Triangle{1, 1, 1}
	t2 := t1
	t2.l1 = 2
	if t1.l1 != t2.l1 {
		fmt.Println("t1 and t2 are separate instances")
	} else {
		fmt.Println("t1 and t2 are the same instance")
	}
	p := &t1
	t3 := *p
	t3.l1 = 3
	if t1.l1 != t3.l1 {
		fmt.Println("t1 and t3 are separate instances")
	} else {
		fmt.Println("t1 and t3 are the same instance")
	}

	// creating struct initializing certain values
	var t Triangle = Triangle{l1: 1}
	var t0 Triangle = Triangle{}
	fmt.Println(t)
	fmt.Println(t0)

}

// go has pointers but no pointer arithmetic
func pointers() {
	x, y := float32(1), complex128(1<<20+3<<16)
	p1, p2 := &x, &y

	fmt.Print("Location in memory of x is ")
	fmt.Println(p1)
	fmt.Print("Value of x is ")
	fmt.Println(*p1)
	fmt.Print("Location in memory of y is ")
	fmt.Println(p2)
	fmt.Print("Value of y is ")
	fmt.Println(*p2)

	fmt.Print("New value of x is ")
	*p1 = rand.Float32()
	fmt.Println(p1)

}

// executes deferred statements when func exits. Adds as stack
func refed() {
	defer fmt.Println(5)
	fmt.Println(1)
	defer fmt.Println(4)
	fmt.Println(2)
	defer fmt.Println(3)
}

func sweetch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Windows or smthing")
	}

	// use to write if-then-else chain, first true breaks
	switch {
	case 1 == 2:
		fmt.Println("notmaff")
	case true == false:
		fmt.Println("notmaff")
	case (2+8i)+(3+6i) == (5 + 14i):
		fmt.Println("maff")
	}
}

func ifs() {
	if 2+5i == 5i+2 {
		fmt.Println(true)
	}
	i := 2
	if 1+2i == 5 {
		fmt.Println(i)
	} else {
		fmt.Println(false)
	}
}
func loops() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// optional init and post statements
	var sum int = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop
	for {
	}
}

const (
	Big   = 1 << 62
	Small = 1 >> 1
)

func printBigSmall() {
	fmt.Println(Big)
	fmt.Println(Small)
}

func quickmaffs() int {
	var four int = 2 + 2
	three := four - 1
	return three
}
