package main
import "fmt"

type Student struct{
	roll_no int;
	stud_name string;
	mark1 int;
	mark2 int;
	mark3 int;
}

func main(){
	var n int
	fmt.Print("Enter no. of students:- ")
	fmt.Scan(&n)
	s := make([]Student, n)
	
	for i := 0; i < n; i++ {
		fmt.Println("\nStudent ", i+1, ":- ")
		fmt.Print("Roll no: ")
		fmt.Scan(&s[i].roll_no)
		fmt.Print("Name: ")
		fmt.Scan(&s[i].stud_name)
		fmt.Print("Mark1: ")
		fmt.Scan(&s[i].mark1)
		fmt.Print("Mark2: ")
		fmt.Scan(&s[i].mark2)
		fmt.Print("Mark3: ")
		fmt.Scan(&s[i].mark3)
	}
	
	fmt.Println("\nDetails of ", n, " Student: ")
	for i := 0; i < n; i++ {
		fmt.Println("\nStudent ", i+1, ":- ")
		fmt.Println("Roll no: ", s[i].roll_no)
		fmt.Println("Name: ", s[i].stud_name)
		fmt.Println("Mark1: ", s[i].mark1)
		fmt.Println("Mark2: ", s[i].mark2)
		fmt.Println("Mark3: ", s[i].mark3)
		total := s[i].mark1 + s[i].mark2 + s[i].mark3
		avg := total/3
		fmt.Println("Total: ", total)
		fmt.Println("Average: ", avg)
	}
}
