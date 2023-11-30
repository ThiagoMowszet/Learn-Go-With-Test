package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

// Methods are very similar to functions but they are called by invoking them on an instance of a particular type.
// Where you can just call functions wherever you like, such as Area(rectangle) you can only call methods on "things".

// Syntax to declare methods
// NOTE: func (receiverName ReceiverType) MethodName(args).
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// NOTE: It is a convention in Go to have the receiver variable be the first letter of the type -> c:Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Permiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
