package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

// MAPS
func main() {
	myMap := make(map[string]User)
	me := User{
		FirstName: "Seyfullah",
		LastName:  "Demirci",
	}
	myMap["me"] = me
	log.Println(myMap["me"].FirstName)

	var myNewVar float32
	myNewVar = 11.1
}

//SLICES
// func main() {
// 	numbers:= []int {1,2,3,4,5,6,7,8,9,10}
// 	log.Println(numbers)
// 	log.Println(numbers [6:9])
// }
