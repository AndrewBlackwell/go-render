package util

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/AndrewBlackwell/go-render/renderer"
)

// LoadOBJ reads an OBJ file from the given reader and returns a Mesh.
func LoadOBJ(r io.Reader) (renderer.Mesh, error) {
	var vertices []renderer.Vector
	var triangles []renderer.Triangle

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "v ") {
			// Parse vertex
			parts := strings.Fields(line)
			if len(parts) < 4 {
				return renderer.Mesh{}, errors.New("invalid vertex line: " + line)
			}
			x, _ := strconv.ParseFloat(parts[1], 64)
			y, _ := strconv.ParseFloat(parts[2], 64)
			z, _ := strconv.ParseFloat(parts[3], 64)
			vertices = append(vertices, renderer.NewVector(x, y, z))
		} else if strings.HasPrefix(line, "f ") {
			// Parse face
			parts := strings.Fields(line)
			if len(parts) < 4 {
				return renderer.Mesh{}, errors.New("invalid face line: " + line)
			}
			a, _ := strconv.Atoi(strings.Split(parts[1], "/")[0])
			b, _ := strconv.Atoi(strings.Split(parts[2], "/")[0])
			c, _ := strconv.Atoi(strings.Split(parts[3], "/")[0])
			triangles = append(triangles, renderer.NewTriangle(
				vertices[a-1], vertices[b-1], vertices[c-1],
			))
		}
	}

	if err := scanner.Err(); err != nil {
		return renderer.Mesh{}, fmt.Errorf("error reading OBJ file: %w", err)
	}

	return renderer.NewMesh(triangles), nil
}
