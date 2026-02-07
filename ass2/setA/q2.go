package main
import "fmt"

func sumOfDigit(n int) int {
	if n == 0 {
		return 0
	}
	return n % 10 + sumOfDigit(n / 10)
}

func main(){
	n := 123
	ans := sumOfDigit(n)
	fmt.Println("Sum of digits of ", n, " : ", ans)
}
