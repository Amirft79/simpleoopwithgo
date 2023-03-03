package main

import "fmt"

func main() {
	b := newBill("amirhossein", "fardi", map[string]int{
		"student id":       112,
		"steudent subject": 2,
	})
	fmt.Println(b)
	b.setName("leila")
	b.setFamily("ghafari")
	fmt.Println(b)
}
