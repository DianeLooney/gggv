package opengl

import "github.com/dianelooney/gggv/internal/carbon"

// OpenGL is only supported up to 4.1 on macOS, and 4.2 adds glMemoryBarrier,
// so we unfortunately have to use this instead.
func (s *ShaderSource) memoryBarrier() {
	carbon.Finish()
}
