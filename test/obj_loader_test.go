package test

import (
	"strings"
	"testing"

	"github.com/AndrewBlackwell/go-render/internal/util"
	"github.com/AndrewBlackwell/go-render/renderer"
)

func TestLoadOBJ(t *testing.T) {
	objData := `
v 0 0 0
v 1 0 0
v 0 1 0
f 1 2 3
`
	reader := strings.NewReader(objData)
	mesh, err := util.LoadOBJ(reader)
	if err != nil {
		t.Fatalf("Failed to load OBJ: %v", err)
	}

	expectedMesh := renderer.NewMesh([]renderer.Triangle{
		renderer.NewTriangle(
			renderer.NewVector(0, 0, 0),
			renderer.NewVector(1, 0, 0),
			renderer.NewVector(0, 1, 0),
		),
	})

	if len(mesh.Triangles) != len(expectedMesh.Triangles) {
		t.Errorf("Expected %d triangles, got %d", len(expectedMesh.Triangles), len(mesh.Triangles))
	}

	for i, triangle := range mesh.Triangles {
		expected := expectedMesh.Triangles[i]
		if triangle.A != expected.A || triangle.B != expected.B || triangle.C != expected.C {
			t.Errorf("Triangle mismatch: expected %v, got %v", expected, triangle)
		}
	}
}
