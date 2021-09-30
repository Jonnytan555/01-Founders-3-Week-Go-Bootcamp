package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}

	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	arguement := os.Args[1:]

	file, err := os.Open(arguement[0])
	if err != nil {
		fmt.Printf("The mistake is: %v\n", err.Error())
	}

	stats, _ := file.Stat()

	theSize := stats.Size()

	arr := make([]byte, theSize)

	file.Read(arr)

	fmt.Print(string(arr))

	file.Close()
}
