package main
import "fmt"

type Number struct{
	n1, n2 int
}

func (n Number) multiplication() int {
	return n.n1 * n.n2
}

func main(){
	n := Number{4, 5}
	fmt.Println("Multiplication: ", n.multiplication())
}
