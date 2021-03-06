package main
import “testing”
func TestMain(t *testing.T){
	r:= Rectangle{length:10, width:10}
	var s Shape
	s = r
	if s.perimeter() != 40{
		t.Error(“Expected 16, got “, s.perimeter())
	}
	c := Circle{pi:3.14, r:10}
	s = c
	if s.perimeter() != 31.400002{
		t.Error(“Expect 62.800004, got “, s.perimeter())
	}
}
