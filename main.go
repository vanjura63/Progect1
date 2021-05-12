package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	fmt.Println(math.Pi * r * r)
	return math.Pi * r * r

}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5
	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))
	fmt.Println("Yes")

	/*x := make(map[string]int)
	x["key"] = 10
	x["old"] = 15
	x["jonn"] = 30
	delete(x, "old")
	fmt.Println(x) */

}

type List struct {
}

func Size(list []int) int {
	return len(list)
}
func Add(list []int, el int) []int {
	return append(list, el)
}

func (linkedList List) Size() int {
	var size int = 0
	if linkedList.head != nil {
		var node Node = linkedList.head.next
		for size = 1; node.next != nil; size++ {
			node = node.next
		}
	}
	return size
}

/*func main() {
	someVariable := "some text"
	ourCallback := func() {
		someVariable = "Other"
		fmt.Println(someVariable)
	}
	runCallback(ourCallback)
}

func runCallback(callback func()) {
	callback()
}*/

/*printInterfaceVal(10)
	printInterfaceVal("string")
	printInterfaceVal(true)
}

func printInterfaceVal(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("целое число: %d\n", v)
	case string:
		fmt.Printf("строка: %s\n", v)
	case bool:
		fmt.Printf("булев тип: %t\n", v)
	}
}

/* fmt.Print("Enter a number: ")
var input float64
fmt.Scanf("%f", &input)
output := input * 2
fmt.Println(output)

number := 0

for i := 20; i > 0; i-- {
	number = i
}

println(number)

//	for i := 1; i < 100; i *= 2 {
//		i += 4

//		if i%2 == 0 {
//			println("чётное")
//		}
//	}

var theMatrixMoviePartsCount = 3

fmt.Printf("У фильма 'Матрица' %d части", theMatrixMoviePartsCount)

const maxCounter = 100

var counter int

for working := true; working != true; {
	counter++

	if counter >= 100 {
		println(counter)

		working = false
	}
}*/
