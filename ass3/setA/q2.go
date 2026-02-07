package main
import "fmt"

type Books struct {
	BookID int
	Title string
	Author string
	Price int
}

func main(){
	var n int
	fmt.Print("Enter no. of books: ")
	fmt.Scan(&n)
	books := make([]Books, n)
	
	for i:=0; i < n; i++ {
		fmt.Println("\nBook ", i+1, ": ")
		fmt.Print("Id:- ")
		fmt.Scan(&books[i].BookID)
		fmt.Print("Title:- ")
		fmt.Scan(&books[i].Title)
		fmt.Print("Author:- ")
		fmt.Scan(&books[i].Author)
		fmt.Print("Price:- ")
		fmt.Scan(&books[i].Price)
	}
	
	fmt.Println("\nDetails of ", n, " books:- ")
	for i:=0; i < n; i++ {
		fmt.Println("\nBook ", i+1, ": ")
		fmt.Println("Id:- ", books[i].BookID)
		fmt.Println("Title:- ", books[i].Title)
		fmt.Println("Author:- ", books[i].Author)
		fmt.Println("Price:- ", books[i].Price)
	}	
}
