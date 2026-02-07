package main
import ("fmt")

func main(){
	s1 := "Hello "
	s2 := "World"
	p1 := &s1
	p2 := &s2
	fmt.Println(*p1, *p2)
}
