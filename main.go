package main
import "fmt"

func main() {
	var arr  []int
	 m :=make(map[string]string)
	fmt.Println(arr, m)
	arr = append(arr, 0)

	m["a"] = "a"

	var value string = "b"
	if value, ok := m["a"];ok {
		fmt.Println(value,ok)
	}
	fmt.Println(value)
}
