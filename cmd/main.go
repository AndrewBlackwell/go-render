package main

import (
	"fmt"
	"os"

	"github.com/AndrewBlackwell/go-render/internal/util"
)

func main() {
	file, err := os.Open("test/obj/test.obj")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	mesh, err := util.LoadOBJ(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Loaded Mesh:", mesh)
}
