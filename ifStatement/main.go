package main

import (
    "errors"
    "fmt"
    "log"
)

func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    a := 10
    b := 0

	// calls divide function and populate result and err
	// if the err variable from above result is not nil
	
    if result, err := divide(a, b); err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("Result of division:", result)
    }
}
