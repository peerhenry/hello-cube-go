package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

func main() {
	log.Println("Application starting...")
	runtime.LockOSThread()
	window := initGlfw(windowWidth, windowHeight)
	defer glfw.Terminate()

	initOpenGL()
	glslProgram := initGlslProgram()

	cube := createCube()
	objects := []*GameObject{&cube}

	renderState := buildGameRenderState(glslProgram.handle)
	world := GameWorld{
		renderState: &renderState,
		glslProgram: &glslProgram,
		objects:     objects,
	}

	log.Println("Now running...")
	for !window.ShouldClose() {
		draw(window, &world)
	}
	log.Println("Application end")
}

func initGlslProgram() GLSLProgram {
	glslProgram := NewGLSLProgram()
	vertex := readFile("shaders/vertex.glsl")
	fragment := readFile("shaders/fragment.glsl")
	glslProgram.CompileAndAttachShader(vertex, gl.VERTEX_SHADER)
	glslProgram.CompileAndAttachShader(fragment, gl.FRAGMENT_SHADER)
	glslProgram.Link()
	return glslProgram
}

func draw(
	window *glfw.Window,
	world *GameWorld,
) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	// todo: update on separate thread
	world.update()
	world.render()
	glfw.PollEvents()
	window.SwapBuffers()
}
