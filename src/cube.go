package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type drawable interface {
	draw(program uint32)
}

type Cube struct {
	vao uint32
	ibo uint32
}

func createCube() Cube {
	vao, ibo := makeVao(positions, normals, indices)
	return Cube{vao, ibo}
}

func (self Cube) draw(program uint32) {
	self.setUniforms(program)
	gl.BindVertexArray(self.vao)

	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, self.ibo)
	// len := int32(len(indices) / 3)
	gl.DrawElements(gl.TRIANGLES, 3, gl.UNSIGNED_INT, nil)
}

func (self Cube) setUniforms(program uint32) {
	// lightposition
	// diffuse reflectivity
	// diffuse light intensity
	// modelview ModelViewMatrix
	angle := 0
	var model mgl32.Mat4 = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})
	var view mgl32.Mat4 = mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	modelview := view.Mul4(model)
	modelviewLoc := gl.GetUniformLocation(program, gl.Str("ModelViewMatrix\x00"))
	gl.UniformMatrix4fv(modelviewLoc, 1, false, &modelview[0])
	// normal matrix
	// projection matrix
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/windowHeight, 0.1, 10.0)
	projectionLoc := gl.GetUniformLocation(program, gl.Str("ProjectionMatrix\x00"))
	gl.UniformMatrix4fv(projectionLoc, 1, false, &projection[0])
	// MVP
	mvpLoc := gl.GetUniformLocation(program, gl.Str("MVP\x00"))
	mvp := projection.Mul4(modelview)
	gl.UniformMatrix4fv(mvpLoc, 1, false, &mvp[0])
}
