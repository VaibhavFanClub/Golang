package main
import "fmt"

type Student struct{
	roll int
	name string
}

func (s *Student) show() {
	fmt.Println("Roll no: ", s.roll)
	fmt.Println("Name: ", s.name)
}

func main(){
	s := Student{1, "Ashish"}
	s.show()
	
	ptr := &s
	ptr.show()
}
