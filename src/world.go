package main

type GameWorld struct {
	renderState *GameRenderState
	glslProgram *GLSLProgram
	objects     []*GameObject
}

func (self *GameWorld) update() {
	for _, object := range self.objects {
		object.update()
	}
}

func (self *GameWorld) render() {
	// todo:
	// lightposition
	// diffuse reflectivity
	// diffuse light intensity
	self.glslProgram.Use()
	for _, object := range self.objects {
		object.draw(self.renderState)
	}
}
