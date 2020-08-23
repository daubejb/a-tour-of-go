package main

import "fmt"

func main() {
	sum := 1
	// for can also be C's while spelled for
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
