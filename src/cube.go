package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type drawable interface {
	draw(program uint32)
}

type updatable interface {
	update(program uint32)
}

type GameObject struct {
	vao         uint32
	ibo         uint32
	angle       float32
	indexLength int32
}

func createCube() GameObject {
	vao, ibo := makeVao(cubePositions, cubeNormals, cubeIndices)
	return GameObject{
		vao,
		ibo,
		0.0,
		int32(len(cubeIndices)),
	}
}

func (self *GameObject) update() {
	self.angle = self.angle + 0.01
}

func (self *GameObject) draw(renderState *GameRenderState) {
	var model mgl32.Mat4 = mgl32.HomogRotate3D(
		self.angle,          // angle
		mgl32.Vec3{0, 0, 1}, // axis
	)
	renderState.SetModelMatrix(model)
	gl.BindVertexArray(self.vao)
	gl.DrawElements(gl.TRIANGLES, self.indexLength, gl.UNSIGNED_INT, nil)
}
