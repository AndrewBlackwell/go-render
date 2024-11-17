package renderer

type Triangle struct {
	A, B, C Vector
}

func NewTriangle(a, b, c Vector) Triangle {
	return Triangle{a, b, c}
}
