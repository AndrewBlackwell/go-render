package test

import (
	"testing"

	"github.com/AndrewBlackwell/go-render/renderer"
)

func TestTriangleOperations(t *testing.T) {
	v1 := renderer.NewVector(1, 2, 3)
	v2 := renderer.NewVector(4, 5, 6)
	v3 := renderer.NewVector(7, 8, 9)
	triangle := renderer.NewTriangle(v1, v2, v3)

	// Test translation
	expectedTranslated := renderer.NewTriangle(
		renderer.NewVector(2, 3, 4),
		renderer.NewVector(5, 6, 7),
		renderer.NewVector(8, 9, 10),
	)
	translated := triangle.Add(renderer.NewVector(1, 1, 1))
	if translated != expectedTranslated {
		t.Errorf("Expected %v, got %v", expectedTranslated, translated)
	}

	// Test scaling
	expectedScaled := renderer.NewTriangle(
		renderer.NewVector(2, 4, 6),
		renderer.NewVector(8, 10, 12),
		renderer.NewVector(14, 16, 18),
	)
	scaled := triangle.ScalarMult(2)
	if scaled != expectedScaled {
		t.Errorf("Expected %v, got %v", expectedScaled, scaled)
	}
}
