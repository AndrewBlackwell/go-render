package renderer

type Mesh struct {
	Triangles []Triangle
}

func NewMesh(triangles []Triangle) Mesh {
	return Mesh{triangles}
}

func (m Mesh) ScalarMult(scalar float64) Mesh {
	newTriangles := make([]Triangle, len(m.Triangles))
	for i, t := range m.Triangles {
		newTriangles[i] = t.ScalarMult(scalar)
	}
	return Mesh{newTriangles}
}
