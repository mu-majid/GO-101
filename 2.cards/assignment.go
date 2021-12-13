package main

import "fmt"

func mainForAssignment() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range nums {
		if v%2 == 0 {
			fmt.Println(string(v) + " is even")
		} else {
			fmt.Println(string(v) + " is odd")
		}
	}
}
