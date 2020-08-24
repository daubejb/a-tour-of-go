package main

import "fmt"

// Vertex - x and y coordinates for a point
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
