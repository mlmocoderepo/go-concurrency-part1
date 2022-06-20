package main

import (
	"fmt"
	"math"
)

// All 3 shapes have their respective method to calculate Area
// And, all 3 shapes have the same method signature: Area() float64
// All 3 shapes are concrete types that satisfies (implements) the Shape interface: Area() float64{}

type Shapes interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	a, b, c float64 // length of sides of the triangle
}

type Rectangle struct {
	length, breath float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Heron's formula for area of a traingle
func (t Triangle) Area() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s * t.a) * (s * t.b) * (s * t.c))
}

func (r Rectangle) Area() float64 {
	return r.length * r.breath
}

// calculate the angle of each side of a given triangle with 3 sides
func (t Triangle) Angles() []float64 {
	return []float64{
		Angle(t.b, t.c, t.a),
		Angle(t.a, t.c, t.b),
		Angle(t.a, t.b, t.c),
	}
}

func Angle(a, b, c float64) float64 {
	return math.Acos((a*a+b*b-c*c)/(2*a*b)) * 180.0 / math.Pi
}

func main() {

	// Circle, Triangle and Retangle are concrete types with methods that satisfies Shape's interface
	// Therefore, the slice s of []Shapes is able to take in all 3 given shapes
	s := []Shapes{
		Circle{1.0},
		Triangle{10, 4, 7},
		Rectangle{5, 10},
	}

	var presentedShape string

	for _, obj := range s {
		// type switch statement to check for the type of the shape used
		switch obj.(type) {
		case Circle:
			presentedShape = "Circle"
		case Triangle:
			presentedShape = "Triangle"
		case Rectangle:
			presentedShape = "Rectangle"
		}

		fmt.Printf("%s area is: %.2f\n", presentedShape, obj.Area())

		// TYPE ASSERTION can also be used to check if the instance is of type Triangle where ok = true
		if t, ok := obj.(Triangle); ok {
			fmt.Println("Angles of the triangle are:", t.Angles())
		}
	}
}
