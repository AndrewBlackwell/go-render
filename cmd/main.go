package main

import (
	"fmt"

	"github.com/AndrewBlackwell/go-render/renderer"
)

func main() {
	triangles := []renderer.Triangle{
		renderer.NewTriangle(
			renderer.NewVector(1, 2, 3),
			renderer.NewVector(4, 5, 6),
			renderer.NewVector(7, 8, 9),
		),
	}
	mesh := renderer.NewMesh(triangles)
	fmt.Println("Original Mesh:", mesh)

	translated := mesh.Add(renderer.NewVector(1, 1, 1))
	fmt.Println("Translated Mesh:", translated)

	scaled := mesh.ScalarMult(2)
	fmt.Println("Scaled Mesh:", scaled)
}
