package main
import "fmt"

func nameReturn()(variable int) {
	variable = 2
	return variable
}

func main(){
	fmt.Println("Returned value: ", nameReturn())
}
