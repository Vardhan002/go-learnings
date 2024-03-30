package main

import "fmt"

type Square struct {
	length float32
}

func (square Square) Area() float32 {
	return square.length * square.length
}

type Area interface {
	Area() float32
}

func main() {

	squareobject := Square{length: 25}

	var shapeObject Area

	shapeObject = squareobject
	fmt.Println("Square Area = ", shapeObject.Area())

}
