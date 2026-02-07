package main
import ("fmt")

func main(){
	a := 1
	p := &a
	pp := &p
	
	fmt.Println("Value of a: ", a)
	fmt.Println("Value of p: ", p)
	fmt.Println("Value of a using pointer p: ", *p)
	fmt.Println("Value of pp: ", pp)
}
