package main
import "fmt"

type Employee struct {
	eno int
	ename string
	salary float32
}

func main(){
	var n int
	fmt.Print("Enter no. of employees: ")
	fmt.Scan(&n)
	e := make([]Employee, n)
	var max float32 = 0.0
	
	fmt.Println("\nEnter details of employees: ")
	for i:=0; i<n; i++ {
		fmt.Println("\nEmployee ", i+1)
		fmt.Print("Eno: ")
		fmt.Scan(&e[i].eno)
		fmt.Print("Ename: ")
		fmt.Scan(&e[i].ename)
		fmt.Print("Salary: ")
		fmt.Scan(&e[i].salary)
		if e[i].salary >= max {
			max = e[i].salary
		}
	}
	
	fmt.Println("\nDetails of employees with max salary: ")
	for i:=0; i<n; i++ {
		if e[i].salary == max {
			fmt.Println("\nEmployee ", i+1)
			fmt.Println("Eno: ", e[i].eno)
			fmt.Println("Ename: ", e[i].ename)
			fmt.Println("Salary: ", e[i].salary)
		}
	}
}
