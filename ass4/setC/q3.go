package main
import "fmt"

type StudentInterface interface{
	display()
}

type MarksInterface interface{
	calc()
}

type FinalDetails interface{
	StudentInterface
	MarksInterface
}

type Student struct{
	rollno int
	name string
	mark1, mark2 float32
}

func (s Student) display() {
	fmt.Println("Rollno: ", s.rollno)
	fmt.Println("Name: ", s.name)
}

func (s Student) calc() {
	fmt.Println("Percentage: ", (s.mark1+s.mark2)/200*100)
}

func main(){
	s := Student{1, "Raj", 78, 87}
	var ss FinalDetails = s
	ss.display()
	ss.calc()
}
