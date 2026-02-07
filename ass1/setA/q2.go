package main
import("fmt")

func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)
	if n%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
}
