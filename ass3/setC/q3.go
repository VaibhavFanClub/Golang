package main
import "fmt"

func main() {
	s := []int{10, 20, 30}
	fmt.Println("Original slice: ", s)
	
	s = append(s, 40, 50)
	fmt.Println("After append: ", s)
	
	index := 2
	s = append(s[:index], s[index+1:]...)
	fmt.Println("After removing element at index 2: ", s)
	
	cs := make([]int, len(s))
	copy(cs, s)
	fmt.Println("Copied slice: ", cs)
}
