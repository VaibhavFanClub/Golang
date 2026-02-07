package main
import ("fmt")

func main(){
	var n int = 123
	var add *int = &n
	fmt.Println("Address of n: ", add)
}
