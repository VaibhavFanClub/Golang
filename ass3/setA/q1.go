package main
import ("fmt")

func main(){
	arr := []int{8,4,5,1,9,2,3,6}
	max := arr[0]
	min := arr[0]
	for i:=0; i < len(arr); i++ {
		if arr[i] >= max {
			max = arr[i]
		}
		if arr[i] <= min {
			min = arr[i]
		}
	}
	fmt.Println("Max = ", max)
	fmt.Println("Min = ", min)
}
