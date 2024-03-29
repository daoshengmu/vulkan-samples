package triangle

import (
	"github.com/xlab/linmath"
)

// Y-up is -1 in Vulkan
var gVertexData = linmath.ArrayFloat32([]float32{
	-0.8, 0.8, 0,	1, 0, 0,
	0.8, 0.8, 0, 0, 1, 0,
	-0.8, -0.8, 0, 0, 0, 1,
	0.8, -0.8, 0, 0.5, 0.5, 0.5,
})

var gIndexData = linmath.ArrayUint16([]uint16{
	1, 0, 2, 1, 2, 3,
})