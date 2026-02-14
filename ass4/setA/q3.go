package main
import "fmt"

type Author struct {
	name, bookName string
	noOfBooks int
}

func (a Author) show() {
	fmt.Println("Author name: ", a.name)
	fmt.Println("Book name: ", a.bookName)
	fmt.Println("No. of books: ", a.noOfBooks)
}

func main(){
	a := Author{"Raj", "Book 1", 5}
	a.show()
	fmt.Println()
	a = Author{"Rohan", "Rohans book", 4}
	a.show()
}
