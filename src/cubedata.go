package main

var (
	cubePositions = []float32{
		-1, -1, -1, // south
		-1, -1, 1,
		1, -1, -1,
		1, -1, 1,

		1, -1, -1, // east
		1, -1, 1,
		1, 1, -1,
		1, 1, 1,

		1, 1, -1, // north
		1, 1, 1,
		-1, 1, -1,
		-1, 1, 1,

		-1, 1, -1, // west
		-1, 1, 1,
		-1, -1, -1,
		-1, -1, 1,

		-1, -1, 1, // top
		-1, 1, 1,
		1, -1, 1,
		1, 1, 1,

		-1, 1, -1, // bottom
		-1, -1, -1,
		1, 1, -1,
		1, -1, -1,
	}
)

var (
	cubeNormals = []float32{
		0, -1, 0, // south normal -y
		0, -1, 0,
		0, -1, 0,
		0, -1, 0,
		1, 0, 0, // east normal +x
		1, 0, 0,
		1, 0, 0,
		1, 0, 0,
		0, 1, 0, // north normal +y
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
		-1, 0, 0, // west normal -x
		-1, 0, 0,
		-1, 0, 0,
		-1, 0, 0,
		0, 0, 1, // top normal +z
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
		0, 0, -1, // bottom normal -z
		0, 0, -1,
		0, 0, -1,
		0, 0, -1,
	}
)

// 36 ints
var (
	cubeIndices = []uint32{
		0, 1, 2, 2, 1, 3, // south
		4, 5, 6, 6, 5, 7, // east
		8, 9, 10, 10, 9, 11, // north
		12, 13, 14, 14, 13, 15, // west
		16, 17, 18, 18, 17, 19, // top
		20, 21, 22, 22, 21, 23, // bottom
	}
)
