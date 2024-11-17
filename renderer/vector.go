package renderer

type Vector struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

func NewVectorFromScalar(scalar float64) Vector {
	return Vector{X: scalar, Y: scalar, Z: scalar}
}
