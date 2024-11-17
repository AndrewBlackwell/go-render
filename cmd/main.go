package main

import (
	"fmt"

	"github.com/AndrewBlackwell/go-render/renderer"
)

func main() {
	v1 := renderer.NewVector(1, 2, 3)
	v2 := renderer.NewVector(4, 5, 6)
	v3 := renderer.NewVector(7, 8, 9)

	triangle := renderer.NewTriangle(v1, v2, v3)
	fmt.Println("Original Triangle:", triangle)

	translated := triangle.Add(renderer.NewVector(1, 1, 1))
	fmt.Println("Translated Triangle:", translated)

	scaled := triangle.A.ScalarMult(2)
	fmt.Println("Scaled Triangle:", scaled)
}
