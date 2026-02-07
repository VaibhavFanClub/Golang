package main
import ("fmt")

func main(){
	var n1 float32
	var n2 float32
	var symbol string
	fmt.Print("Enter number 1: ")
	fmt.Scanf("%f", &n1)
	fmt.Print("Enter number 2: ")
	fmt.Scanf("%f", &n2)
	fmt.Print("Enter symbol (+, -, *, /): ")
	fmt.Scanf("%s", &symbol)
	switch symbol {
		case "+":
			fmt.Println("Sum: ", n1 + n2)
		case "-":
			fmt.Println("Substraction: ", n1 - n2)
		case "*":
			fmt.Println("Multiplication: ", n1 * n2)
		case "/":
			fmt.Println("Division: ", n1 / n2)
		default :
			fmt.Println("Incorrect symbol")
	}
}
