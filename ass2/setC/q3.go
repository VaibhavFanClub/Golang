package main
import "fmt"

func mul() (int, int) {
	a := 1
	b := 2
	return a, b
}

func main(){
	a, b := mul()
	fmt.Printf("a = %d, b = %d\n", a, b)
}
