package main

import "fmt"

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}

func FindNextPrime(nb int) int {
	a := nb + 1

	for i := 2; i < a+1; i++ {
		if nb%i == 0 {
			break
		}
	}

	if i == a {
		fmt.Printf(a, "is the next prime number from "+nb)
		break
	} else {
		a += 1
	}
}
