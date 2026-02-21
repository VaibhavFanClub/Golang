package main
import "fmt"

type Student struct{
	rollno int
	name string
	percentage float32
}

type Students []Student
func (s Students) sort() {
	for i:=1; i<len(s); i++{
		for j:=0; j<len(s)-i; j++{
			if s[j].percentage < s[j+1].percentage{
				temp := s[j]
				s[j] = s[j+1]
				s[j+1] = temp
			}
		}
	}
}

func main(){
	var n int
	fmt.Print("Enter no of students: ")
	fmt.Scan(&n)
	fmt.Println("\nEnter details of students: ")
	s := make(Students, n)
	for i:=0; i<n; i++{
		fmt.Println("\nStudent ", i+1)
		fmt.Print("Roll no: ")
		fmt.Scan(&s[i].rollno)
		fmt.Print("Name: ")
		fmt.Scan(&s[i].name)
		fmt.Print("Percentage: ")
		fmt.Scan(&s[i].percentage)
	}
	
	s.sort()
	fmt.Println("\nDetails of students: ")
	for i:=0; i<n; i++{
		fmt.Println("\nStudent ", i+1)
		fmt.Println("Roll no: ", s[i].rollno)
		fmt.Println("Name: ", s[i].name)
		fmt.Println("Percentage: ", s[i].percentage)
	}
}
