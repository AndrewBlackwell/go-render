package renderer

// Traingle is a struct that represents a triangle in 3D space
type Triangle struct {
	A, B, C Vector
}

// NewTriangle creates a new Triangle struct
func NewTriangle(a, b, c Vector) Triangle {
	return Triangle{a, b, c}
}

// Add translates the triangle by adding a vector to all its vertices.
func (t Triangle) Add(v Vector) Triangle {
	return Triangle{
		A: t.A.Add(v),
		B: t.B.Add(v),
		C: t.C.Add(v),
	}
}

// MulScalar scales the triangle by multiplying each vertex by a scalar.
func (t Triangle) ScalarMult(scalar float64) Triangle {
	return Triangle{
		t.A.ScalarMult(scalar),
		t.B.ScalarMult(scalar),
		t.C.ScalarMult(scalar),
	}
}

// Normal calculates the normal vector of the triangle.
func (t Triangle) Normal() Vector {
	ab := t.B.Sub(t.A)
	ac := t.C.Sub(t.A)

	// Cross product to calculate the normal vector.
	return Vector{
		ab.Y*ac.Z - ab.Z*ac.Y,
		ab.Z*ac.X - ab.X*ac.Z,
		ab.X*ac.Y - ab.Y*ac.X,
	}.ScalarMult(1.0 / ab.Dot(ac))
}
