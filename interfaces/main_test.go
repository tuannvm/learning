package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		10.0,
		10.0,
	}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func checkArea(t *testing.T, s Shape, want float64) {
	// to return error from test case itself, not this function
	t.Helper()
	got := s.Area()

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			7.0,
			2.0,
		}
		want := 14.0
		checkArea(t, rectangle, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})
}

func TestAreaTable(t *testing.T) {
	// Table driven test https://github.com/golang/go/wiki/TableDrivenTests
	testCases := []struct {
		s    Shape
		want float64
	}{
		{
			Rectangle{
				7.0,
				2.0,
			},
			14.0,
		},
		{
			Circle{10.0},
			314.1592653589793,
		},
	}

	for _, testCase := range testCases {
		checkArea(t, testCase.s, testCase.want)
	}
}
