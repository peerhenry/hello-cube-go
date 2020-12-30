package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	windowWidth  = 500
	windowHeight = 500
)

func main() {
	log.Println("Application starting...")
	runtime.LockOSThread()
	window := initGlfw(windowWidth, windowHeight)
	defer glfw.Terminate()

	initOpenGL()
	glslProgram := NewGLSLProgram()
	vertex := readFile("shaders/vertex.glsl")
	fragment := readFile("shaders/fragment.glsl")
	log.Println("Compiling vertex shader...")
	glslProgram.CompileAndAttachShader(vertex, gl.VERTEX_SHADER)
	log.Println("Compiling fragment shader...")
	glslProgram.CompileAndAttachShader(fragment, gl.FRAGMENT_SHADER)
	log.Println("Now linking shaders...")
	glslProgram.Link()

	cube := createCube()
	// objects := []Cube{cube}

	log.Println("Now running...")
	for !window.ShouldClose() {
		draw(window, glslProgram, &cube)
	}
	log.Println("Application end")
}

func initOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL Version", version)
}

func draw(window *glfw.Window, program GLSLProgram, object *Cube) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	program.Use()

	// for _, object := range objects {
	// 	log.Println("starting with", object.angle)
	// 	object.draw(program.GetHandle())
	// }
	object.draw(program.GetHandle())

	glfw.PollEvents()
	window.SwapBuffers()
}
