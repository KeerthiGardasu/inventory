package main

import (
	"fmt"
)

func main() {
	myinv := newinventory("india bazar")
	myinv.additems("Rice bags", 10)
	myinv.additems("suagar packets", 15)
	myinv.additems("salt", 11)
	myinv.additems("Toor dhal", 50)
	myinv.additems("Onion Bags", 40)
	myinv.removeitems("Onion Bags")

	fmt.Println(myinv.format())
	myinv.updateitems("tamirind", 15)
	fmt.Println(myinv.format())

	fmt.Println(myinv.format())
	myinv.save()

}
