package main

import "fmt"

func main() {
	fmt.Println("Hellooo!!! :))")

	var whatToSay string
	var i int
	i = 61

	whatToSay = "Goodbye,cruel world"

	fmt.Println(whatToSay)

	fmt.Println("i setted by me", i)
	whatWasSaid := saySomething()
	fmt.Println("I said", whatWasSaid)

}

func saySomething() string {
	return "I am Seyfullah"
}
