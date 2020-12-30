package main

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type drawable interface {
	draw(program uint32)
}

type Cube struct {
	vao   uint32
	ibo   uint32
	angle float32
}

func createCube() Cube {
	log.Println("Creating cube")
	vao, ibo := makeVao(positions, normals, indices)
	return Cube{vao, ibo, 0.0}
}

func (self *Cube) draw(program uint32) {
	self.angle = self.angle + 0.01
	self.setUniforms(program)
	gl.BindVertexArray(self.vao)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, self.ibo)
	len := int32(len(indices)) // todo, assert len == 36
	gl.DrawElements(gl.TRIANGLES, len, gl.UNSIGNED_INT, nil)
}

func (self *Cube) setUniforms(program uint32) {
	// lightposition
	// diffuse reflectivity
	// diffuse light intensity

	// modelview ModelViewMatrix
	var model mgl32.Mat4 = mgl32.HomogRotate3D(
		self.angle,          // angle
		mgl32.Vec3{0, 0, 1}, // axis
	)
	var view mgl32.Mat4 = mgl32.LookAtV(
		mgl32.Vec3{5, 3, 3}, // eye
		mgl32.Vec3{0, 0, 0}, // center
		mgl32.Vec3{0, 0, 1}, // up
	)
	modelview := view.Mul4(model)
	modelviewLoc := gl.GetUniformLocation(program, gl.Str("ModelViewMatrix\x00"))
	gl.UniformMatrix4fv(modelviewLoc, 1, false, &modelview[0])

	// normal matrix
	normalLoc := gl.GetUniformLocation(program, gl.Str("NormalMatrix\x00"))
	normalMatrix := modelview.Mat3()
	gl.UniformMatrix4fv(normalLoc, 1, false, &normalMatrix[0])

	// projection matrix
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/windowHeight, 0.1, 10.0)
	projectionLoc := gl.GetUniformLocation(program, gl.Str("ProjectionMatrix\x00"))
	gl.UniformMatrix4fv(projectionLoc, 1, false, &projection[0])

	// MVP
	mvpLoc := gl.GetUniformLocation(program, gl.Str("MVP\x00"))
	mvp := projection.Mul4(modelview)
	gl.UniformMatrix4fv(mvpLoc, 1, false, &mvp[0])
}
