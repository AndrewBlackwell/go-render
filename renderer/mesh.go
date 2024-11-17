package renderer

type Mesh struct {
	Triangles []Triangle
}

func NewMesh(triangles []Triangle) Mesh {
	return Mesh{triangles}
}
