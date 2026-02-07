package main
import ("fmt")

func main(){
	n1 := 4
	n2 := 5
	n1 += n2
	n2 = n1 - n2
	n1 -= n2
	fmt.Printf("Swapped n1: %d n2: %d\n", n1, n2)
}
