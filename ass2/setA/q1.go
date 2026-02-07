package main
import "fmt"

func add(n1, n2 int) {
	fmt.Printf("Addition (%d, %d): %d\n", n1, n2, n1 + n2)
}

func main(){
	n1 := 4
	n2 := 54
	add(n1, n2)
}
