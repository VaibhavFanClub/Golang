package main
import "fmt"

type array [5]int

func (a *array) copy(srcArray array) {
	for i:=0; i<len(srcArray); i++ {
		a[i] = srcArray[i]
	}	
}

func main(){
	a := array{1, 2, 3, 4, 5}
	var b array
	b.copy(a)
	fmt.Println("Copied array: ", b)
}
