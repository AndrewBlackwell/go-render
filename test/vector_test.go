package test

import (
	"testing"

	"github.com/AndrewBlackwell/go-render/renderer"
)

func TestVectorOperations(t *testing.T) {
	// Test addition
	v1 := renderer.NewVector(1, 2, 3)
	v2 := renderer.NewVector(4, 5, 6)
	expectedSum := renderer.NewVector(5, 7, 9)
	sum := v1.Add(v2)
	if sum != expectedSum {
		t.Errorf("Addition failed: expected %v, got %v", expectedSum, sum)
	}

	// Test subtraction
	expectedDiff := renderer.NewVector(-3, -3, -3)
	diff := v1.Sub(v2)
	if diff != expectedDiff {
		t.Errorf("Subtraction failed: expected %v, got %v", expectedDiff, diff)
	}

	// Test scalar multiplication
	scalar := 2.0
	expectedScaled := renderer.NewVector(2, 4, 6)
	scaled := v1.ScalarMult(scalar)
	if scaled != expectedScaled {
		t.Errorf("Scalar multiplication failed: expected %v, got %v", expectedScaled, scaled)
	}

	// Test dot product
	expectedDot := float64(32)
	dot := v1.Dot(v2)
	if dot != expectedDot {
		t.Errorf("Dot product failed: expected %v, got %v", expectedDot, dot)
	}
}
