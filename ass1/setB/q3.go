package main
import ("fmt")

func main(){
	n1 := 0
	n2 := 1
	var n int
	fmt.Print("Enter number of terms: ")
	fmt.Scanf("%d", &n)
	fmt.Print("Fibonacci series: \n", n1, " ", n2, " ")
	for i:=2; i<n; i++ {
		n3 := n1 + n2
		fmt.Print(n3, " ")
		n1 = n2
		n2 = n3
	}
}
