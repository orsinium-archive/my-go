package main

import (
	"fmt"
	"math"
)

// Rectangle is struct for triangle shape
type Rectangle struct {
	a, b float64
}

func (shape *Rectangle) area() float64 {
	return shape.a * shape.b
}

// Circle is struct for circle shape
type Circle struct {
	r float64
}

func (shape *Circle) area() float64 {
	return math.Pi * math.Pow(shape.r, 2)
}

// Shape is interface for shapes with area method
type Shape interface {
	area() float64
}

// Polygon is collection of shapes
type Polygon struct {
	shapes []Shape
}

func (polygon *Polygon) area() (result float64) {
	for _, shape := range polygon.shapes {
		result += shape.area()
	}
	return
}

func main() {
	sail := Circle{4.5}
	basket := Rectangle{2, 3.5}
	shapes := []Shape{&sail, &basket}
	balloon := Polygon{shapes}
	fmt.Printf("Sail area: %f\n", sail.area())
	fmt.Printf("Basket area: %f\n", basket.area())
	fmt.Printf("Balloon area: %f\n", balloon.area())
}
