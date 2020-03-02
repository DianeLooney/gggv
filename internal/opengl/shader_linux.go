package opengl

import "github.com/dianelooney/gggv/internal/carbon"

func (s *ShaderSource) memoryBarrier() {
	carbon.MemoryBarrier(carbon.SHADER_IMAGE_ACCESS_BARRIER_BIT)
}
