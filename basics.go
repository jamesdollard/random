package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
)

func main() {
	var f func() int = functions(math.Atan2)
	fmt.Println(f())
}

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
