package main
import "fmt"

type Shape interface{
	area() float32
	perimeter() float32
}

type Circle struct{
	radius float32
}

type Rectangle struct{
	length, breadth float32
}

func (c Circle) area() float32 {
	return 3.14 * c.radius * c.radius;
}

func (r Rectangle) area() float32 {
	return r.length * r.breadth;
}

func (c Circle) perimeter() float32 {
	return 2 * 3.14 * c.radius * c.radius;
}

func (r Rectangle) perimeter() float32 {
	return 2 * (r.length + r.breadth);
}

func main(){
	var s Shape;
	s = Circle{radius: 5.0}
	fmt.Println("Area of circle: ", s.area())
	fmt.Println("Perimeter of circle: ", s.perimeter())
	
	s = Rectangle{length: 4.0, breadth: 5.0}
	fmt.Println("Area of rectangle: ", s.area())
	fmt.Println("Perimeter of rectangle: ", s.perimeter())
}
