package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine:", i)
		}
	}()

	fmt.Println("___")
}
