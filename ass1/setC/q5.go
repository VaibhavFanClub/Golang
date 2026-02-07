package main
import (
	"fmt"
	"strings"
)

func main(){
	s1 := "This is main string"
	s2 := "string"
	isSubstring := strings.Contains(s1, s2)
	if isSubstring {
		fmt.Printf("%s is substring of %s\n", s2, s1)
	} else {
		fmt.Printf("%s is not substring of %s\n", s2, s1)
	}
}
