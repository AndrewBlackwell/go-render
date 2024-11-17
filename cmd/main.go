package main

import (
	"fmt"

	"github.com/AndrewBlackwell/go-render/renderer"
)

func main() {
	v1 := renderer.NewVector(1, 2, 3)
	v2 := renderer.NewVector(4, 5, 6)

	sum := v1.Add(v2)
	dot := v1.Dot(v2)

	fmt.Println("Sum:", sum)
	fmt.Println("Dot Product:", dot)
}
