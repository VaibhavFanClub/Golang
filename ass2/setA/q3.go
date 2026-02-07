package main
import "fmt"

func isPalindrome(n int) bool {
	og := n
	res := 0
	for n != 0 {
		rem := n % 10
		res = res * 10 + rem
		n /= 10
	}
	if og == res {
		return true
	}
	return false
}

func main(){
	n := 121
	ans := isPalindrome(n)
	if ans {
		fmt.Println("Number is palindorme: ", n)
	} else {
		fmt.Println("Number is not palindorme: ", n)
	}
}
