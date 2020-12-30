package main

import "github.com/go-gl/gl/v4.1-core/gl"

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(positions []float32, normals []float32, indices []uint32) (uint32, uint32) {
	posVbo, normalVbo := makeFloatVbo(positions), makeFloatVbo(normals)
	indexVbo := makeIntVbo(indices)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, posVbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, normalVbo)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexVbo)

	gl.BindVertexArray(0)

	return vao, indexVbo
}

func makeFloatVbo(data []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(data), gl.Ptr(data), gl.STATIC_DRAW)
	return vbo
}

func makeIntVbo(data []uint32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(data), gl.Ptr(data), gl.STATIC_DRAW)
	return vbo
}
