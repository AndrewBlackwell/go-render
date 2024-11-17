package test

import (
	"reflect"
	"testing"

	"github.com/AndrewBlackwell/go-render/renderer"
)

// compareMeshes checks if two meshes are equal.
func compareMeshes(m1, m2 renderer.Mesh) bool {
	if len(m1.Triangles) != len(m2.Triangles) {
		return false
	}
	for i := range m1.Triangles {
		if !reflect.DeepEqual(m1.Triangles[i], m2.Triangles[i]) {
			return false
		}
	}
	return true
}

func TestMeshOperations(t *testing.T) {
	triangles := []renderer.Triangle{
		renderer.NewTriangle(
			renderer.NewVector(1, 2, 3),
			renderer.NewVector(4, 5, 6),
			renderer.NewVector(7, 8, 9),
		),
	}
	mesh := renderer.NewMesh(triangles)

	// Test translation
	expectedTranslated := renderer.NewMesh([]renderer.Triangle{
		renderer.NewTriangle(
			renderer.NewVector(2, 3, 4),
			renderer.NewVector(5, 6, 7),
			renderer.NewVector(8, 9, 10),
		),
	})
	translated := mesh.Add(renderer.NewVector(1, 1, 1))
	if !compareMeshes(translated, expectedTranslated) {
		t.Errorf("Translation failed: expected %v, got %v", expectedTranslated, translated)
	}

	// Test scaling
	expectedScaled := renderer.NewMesh([]renderer.Triangle{
		renderer.NewTriangle(
			renderer.NewVector(2, 4, 6),
			renderer.NewVector(8, 10, 12),
			renderer.NewVector(14, 16, 18),
		),
	})
	scaled := mesh.ScalarMult(2)
	if !compareMeshes(scaled, expectedScaled) {
		t.Errorf("Scaling failed: expected %v, got %v", expectedScaled, scaled)
	}
}
