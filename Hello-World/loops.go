package main

import (
	"log"
)

func main() {
	var firstLine = "Once upon a midnight dreary"
	firstLine = "x"
	// animals := make(map[string]string)
	// animals["dog"]="Fido"
	// animals["dog"]="Fluffy"
	for i, l := range firstLine {
		log.Println(i, ":", l)
	}
}
