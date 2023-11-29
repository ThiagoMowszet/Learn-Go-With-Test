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
	t.Run("rectangles", func(t *testing.T){
			rectangle := Rectangle{12.0, 6.0}
			got := rectangle.Area()
			want := 72.0

			if got != want {
			// NOTE: Use of g will print a more precise decimal number in the error message
			t.Errorf("\nGOT = %g\nWANT = %g\n", got, want)
		}
	})

	t.Run("rectangles", func(t *testing.T){
			circle := Circle{10}
			got := circle.Area()
			want := 314.1592653589793 

			if got != want {
			t.Errorf("\nGOT = %g\nWANT = %g\n", got, want)
		}
	})

}