package main

import (
	"fmt"
	"os"
)

type inventory struct {
	name  string
	items map[string]float64
}

func newinventory(name string) inventory {
	i := inventory{
		name:  name,
		items: map[string]float64{},
	}
	return i
}
func (i *inventory) format() string {
	fs := "total inventory\n"

	for k, v := range i.items {
		fs += fmt.Sprintf("%-25v.....%v\n", k+":", v)
	}
	return fs
}
func (i *inventory) additems(name string, quantity float64) {
	i.items[name] = quantity

}
func (i *inventory) removeitems(name string) {
	if _, exists := i.items[name]; exists {
		delete(i.items, name)
	} else {
		fmt.Printf("Item %s does not exist in inventory.\n", name)

	}
}
func (i *inventory) updateitems(name string, quantity float64) {
	if _, exists := i.items[name]; exists {
		i.items[name] = quantity
	} else {
		fmt.Printf("Item %s does not exist in inventory.\n", name)

	}
}

// save the items
func (i *inventory) save() {
	data := []byte(i.format())
	err := os.WriteFile("list/"+i.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("list of items are saved to file")
}
