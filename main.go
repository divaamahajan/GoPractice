package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12};
	for _,n := range numbers {
		if n%2 == 0 {
			fmt.Println(n," is an even number")
		} else {
            fmt.Println(n," is an odd number")
        }
	}
}
