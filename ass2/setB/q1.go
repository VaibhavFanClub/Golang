package main
import "fmt"

func swap(n1, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func main(){
	n1 := 4
	n2 := 5
	swap(&n1, &n2)
	fmt.Println("After swapping: ", n1, n2)
}
