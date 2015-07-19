package main

import (
	"fmt"
)

func main() {
	doors := [100]bool{}

	for pass := 1; pass <= 100; pass++ {
		for i := 0; i < 100; i += pass {
			doors[i] = !doors[i]	
		}
	}

	fmt.Println("Open doors:")
	for num, isOpen := range doors {
		if isOpen  {
			fmt.Println("Door ", num)
		}
	}
}
