package main
import "fmt"

func main(){
	arr := []int{5, 3, 6, 4, 1, 2,  9, 7}
	
	for pass := 1; pass <= len(arr); pass++ {
		for i := 0; i < len(arr) - pass; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
	}
	
	fmt.Println("Array after sorting: ", arr)
}
