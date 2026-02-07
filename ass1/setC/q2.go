package main
import ("fmt")

func main(){
	var s1, s2 string
	fmt.Print("Enter string 1: ")
	fmt.Scanf("%s", &s1)
	fmt.Print("Enter string 2: ")
	fmt.Scanf("%s", &s2)
	if s1 == s2 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}
}
