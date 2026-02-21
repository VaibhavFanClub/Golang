package main
import "fmt"

func sumOfSquares(n int, c chan int){
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	c <- sum
}

func sumOfCubes(number int, ch chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	ch <- sum
}

func main(){
	num := 123
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go sumOfSquares(num, ch1)
	go sumOfCubes(num, ch2)
	
	res1 := <- ch1
	res2 := <- ch2
	
	fmt.Println("sum of squares: ", res1)
	fmt.Println("sum of cubes: ", res2)
}
