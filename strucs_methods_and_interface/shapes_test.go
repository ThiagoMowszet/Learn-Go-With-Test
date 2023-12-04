package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Permiter(rectangle)
	want := 40.0

	if got != want {
		// NOTE: The f is for our float64 and the .2 means print 2 decimal places.
		t.Errorf("\nGOT = %2.f\nWANT = %.2f\n", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		// NOTE: The f is for our float64 and the .2 means print 2 decimal places.
		t.Errorf("\nGOT = %2.f\nWANT = %.2f\n", got, want)
	}
}

func TestArea2(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			// NOTE: Use of g will print a more precise decimal number in the error message
			t.Errorf("\nGOT = %g\nWANT = %g\n", got, want)
		}
	})

	t.Run("rectangles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("\nGOT = %g\nWANT = %g\n", got, want)
		}
	})

}

func TestArea3(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("\nGOT = %g\nWANT = %g\n", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

// Table Driven Tests
func TestArea4(t *testing.T) {
	// anonymous struct
	areaTest := []struct {
		shape Shape
		want  float64
	}{

		/*{Rectangle{12, 6}, 72.0},
		  {Circle{10}, 314.1592653589793},
		  {Triangle{12, 6}, 36.0},*/

		// NOTE: So far you've only been shown syntax for creating instances of structs MyStruct{val1, val2} but you can optionally name the fields.
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0}}

	for _, tt := range areaTest {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}

	}
}

func TestArea5(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		// And you can run specific tests within your table with `go test -run TestArea5/Rectangle`.
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				// NOTE: The %#v format string will print out our struct with the values in its field, so the developer can see at a glance the properties that are being tested.
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}
