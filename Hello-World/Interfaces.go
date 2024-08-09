package main

import "fmt"

type Animal interface {
	Says() string
	NumberofLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberofTeeth int
}

func main() {
	dog := Dog{
		Name:  "Karabaş",
		Breed: "German Shepherd",
	}
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "Black",
		NumberofTeeth: 38,
	}
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberofLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberofLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Ughughugh"
}

func (g *Gorilla) NumberofLegs() int {
	return 2
}
