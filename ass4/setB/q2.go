package main
import "fmt"

func main() {
    var value interface{} = "Raj"
    switch t := value.(type){
        case int64:
            fmt.Println("Type is an integer:", t)

        case float64:
            fmt.Println("Type is a float:", t)

        case string:
            fmt.Println("Type is a string:", t)

        case nil:
            fmt.Println("Type is nil.")

        case bool:
            fmt.Println("Type is a bool:", t)

        default:
            fmt.Println("Type is unknown!")
    }
}
