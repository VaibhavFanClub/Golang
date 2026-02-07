package main
import ("fmt")

func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)
	if n >= 0 && n <= 10 {
		fmt.Println("Number is single digit")
	} else {
		fmt.Println("Number is not single digit")
	}
}
