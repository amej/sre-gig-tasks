package main

import "fmt"

func main() {
	var item = []string{"item1", "item2", "item3", "item4", "item5"}
	for i, v := range item {
		fmt.Printf("This is %s and its index in the array is %d. \n", v, i)
	}
}
