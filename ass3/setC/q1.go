package main
import "fmt"

func main(){
	var m1, m2, res [3][3]int
	fmt.Println("Enter matrix 1:- ")
	for i:=0; i < 3; i++ {
		for j:=0; j<3; j++ {
			fmt.Print("[", i, "]", "[", j, "]: ")
			fmt.Scan(&m1[i][j])
		}
	}
	fmt.Println("Enter matrix 2:- ")
	for i:=0; i < 3; i++ {
		for j:=0; j<3; j++ {
			fmt.Print("[", i, "]", "[", j, "]: ")
			fmt.Scan(&m2[i][j])
		}
	}
	for i:=0; i < 3; i++ {
		for j:=0; j<3; j++ {
			for k:=0; k<3; k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	fmt.Println("Multiplication: ", res)
}
