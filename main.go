package main

import "fmt"

func main() {
	b := newBill()
	fmt.Println(b)
	b.setName("bardia")
	fmt.Println(b)
	b.setName("another name ")
	fmt.Println(b.getName())
	b.setItems(map[string]int{
		"student id ":     99,
		"student subject": 2,
	})
	fmt.Println(b.items)

}
