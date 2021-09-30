package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(door *Door) bool {
	PrintStr("Door Opening...")
	Door.state = "CLOSE"
	return true
}

func CloseDoor(door *Door) bool {
	PrintStr("Door Closing...")
	Door.state = "CLOSE"
	return false
}

func IsDoorOpen(door *Door) bool {
	PrintStr("is the Door opened ?")
	Door.state = "OPEN"
	return true
}

func IsDoorClose(door *Door) bool {
	PrintStr("is the Door closed ?")
	Door.state = "CLOSE"
	return false
}

func main() {
	door := &Door{}

	OpenDoor(door)

	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}
