package main

import "fmt"

func main() {
	var Menu = []string{"hamburger", "salad"}
	for _, v := range Menu {
		fmt.Println("Food: ", v)
	}
}
