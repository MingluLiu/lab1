package main
import ("fmt"; "math")

type Shape interface{
	perimeter() float64
}

type Circle struct {
     r float64 }

type Rectangle struct {
     l,w float64
}

func (r *Rectangle) Perimeter() float64 {
     return 2*r.l + 2*r.w
}


func (c *Circle) Perimeter() float64 {
     return math.Pi * 2*c.r
}

func main(){
     c := Circle{5}
     fmt.Println(c.Perimeter())
     r := Rectangle{10, 10}
     fmt.Println(r.Perimeter())
}
