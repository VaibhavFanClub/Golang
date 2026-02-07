package main
import ("fmt")

func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)
	for i:=1; i<=10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}
