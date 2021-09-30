package main

import "fmt"

type Pilot struct {
	Name     string
	Life     float32
	Age      int
	Aircraft int
}

func isdonnie(donnie *Pilot) {
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1
}

func main() {
	donnie := &Pilot{}

	isdonnie(donnie)

	fmt.Println(donnie)
}

const AIRCRAFT1 = 1
