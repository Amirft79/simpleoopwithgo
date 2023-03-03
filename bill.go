package main

type bill struct {
	name   string
	family string
	items  map[string]int
}

func newBill() bill {
	b := bill{
		name:   "amirhossein",
		family: "reza",
		items: map[string]int{
			"student id": 98234,
			"subject":    3,
		},
	}
	return b
}

func (b *bill) setName(name string) {
	b.name = name
}
func (b *bill) setFamily(family string) {
	b.family = family
}

func (b *bill) getName() string {
	return b.name
}
func (b *bill) geyFamily() string {
	return b.family
}
func (b *bill) setItems(items map[string]int) {
	b.items = items
}

func (b *bill) getItems() map[string]int {
	return b.items
}
