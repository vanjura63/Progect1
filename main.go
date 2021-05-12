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
