// package main

// import (
// 	"log"
// 	"time"
// )

// type User struct {
// 	firstName   string
// 	lastName    string
// 	phoneNumber string
// 	age         int
// 	birthDate   time.Time
// }

// func main() {
// 	user := User{
// 		firstName:   "Seyfullah",
// 		lastName:    "Demirci",
// 		phoneNumber: "5389717460"}
// 	log.Println(user.firstName, user.lastName, user.phoneNumber)
// }

package main

import "log"


type myStruct struct{
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main(){
	var myvar myStruct
	myvar.FirstName ="Seyfullah"

	myvar2:=myStruct{
		FirstName="Elizabeth",
	}

	log.Println(myvar.printFirstName()," get marrieed with")
	log.Println(myvar2.printFirstName())
}