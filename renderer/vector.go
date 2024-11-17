package renderer

// Vector represents a 3D vector with X, Y, and Z components.
type Vector struct {
	X, Y, Z float64
}

// Creates a new vector with the given X, Y, and Z.
func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

// Creates a new vector with the given scalar value for all components.
func NewVectorFromScalar(scalar float64) Vector {
	return Vector{X: scalar, Y: scalar, Z: scalar}
}

// Vector Addiditon
func (a Vector) Add(b Vector) Vector {
	return NewVector(a.X+b.X, a.Y+b.Y, a.Z+b.Z)
}

// Vector Subtraction
func (a Vector) Sub(b Vector) Vector {
	return NewVector(a.X-b.X, a.Y-b.Y, a.Z-b.Z)
}

// Scalar Multiplication
func (a Vector) ScalarMult(scalar float64) Vector {
	return NewVector(a.X*scalar, a.Y*scalar, a.Z*scalar)
}

// Scalar Multiplication
func (a Vector) ScalarDiv(scalar float64) Vector {
	return NewVector(a.X/scalar, a.Y/scalar, a.Z/scalar)
}

// Dot Product of two vectors
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}
