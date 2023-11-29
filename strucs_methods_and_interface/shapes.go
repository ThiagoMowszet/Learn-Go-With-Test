package main

type Rectangle struct {
	width  float64
	height float64
}

// Syntax to declare methods
// A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.
// NOTE: func (receiverName ReceiverType) MethodName(args).
func (r Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 0
}

func Permiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.height + rectangle.width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.height * rectangle.width
}
