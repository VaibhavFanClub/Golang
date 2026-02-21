package main
import "fmt"

func main(){
	var i interface{} = 42
	v := i.(int)
	fmt.Println("i =", v)
}
