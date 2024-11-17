package renderer

import "math"

type Mesh struct {
	Triangles []Triangle
}

func NewMesh(triangles []Triangle) Mesh {
	return Mesh{triangles}
}

// Translates the mesh by adding vector to all triangles
func (m Mesh) Add(v Vector) Mesh {
	newTriangles := make([]Triangle, len(m.Triangles))
	for i, t := range m.Triangles {
		newTriangles[i] = t.Add(v)
	}
	return Mesh{Triangles: newTriangles}
}

// Scales the mesh by a scalar
func (m Mesh) ScalarMult(scalar float64) Mesh {
	newTriangles := make([]Triangle, len(m.Triangles))
	for i, t := range m.Triangles {
		newTriangles[i] = t.ScalarMult(scalar)
	}
	return Mesh{newTriangles}
}

// Bounding box of the mesh.
func (m Mesh) Bounds() (Vector, Vector) {
	if len(m.Triangles) == 0 {
		return Vector{}, Vector{}
	}

	min := m.Triangles[0].A
	max := m.Triangles[0].A

	for _, t := range m.Triangles {
		vertices := []Vector{t.A, t.B, t.C}
		for _, v := range vertices {
			min = Vector{
				X: math.Min(min.X, v.X),
				Y: math.Min(min.Y, v.Y),
				Z: math.Min(min.Z, v.Z),
			}
			max = Vector{
				X: math.Max(max.X, v.X),
				Y: math.Max(max.Y, v.Y),
				Z: math.Max(max.Z, v.Z),
			}
		}
	}

	return min, max
}
