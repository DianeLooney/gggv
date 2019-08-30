package errors

import (
	"fmt"
)

func newErr(kind string, s string) func(args ...interface{}) error {
	return func(args ...interface{}) error {
		return err{
			kind:      kind,
			fmtString: fmt.Sprintf("[%s] %s", kind, s),
			args:      args,
		}
	}
}

type err struct {
	kind      string
	fmtString string
	args      []interface{}
}

func (e err) Error() string {
	return fmt.Sprintf(e.fmtString, e.args...)
}

// Well-known errors
var (
	GLFShaderCompile = newErr("GLFShaderCompile", "fragment shader compilation failed: %v")
	GLVShaderCompile = newErr("GLVShaderCompile", "vertex shader compilation failed: %v")
	GLLinkProgram    = newErr("GLLinkProgram", "linking program %s failed: %v")

	FFDecoderOpenInput        = newErr("FFDecodeOpenInput", "unable to open input for file %s: %v")
	FFDecoderStreamInfo       = newErr("FFDecoderStreamInfo", "unable to find stream info for file %s: %v")
	FFDecoderUnsupportedCodec = newErr("FFDecoderUnsupportedCodec", "unsupported codec in file %s: %v")
	FFDecoderCopyCodecCtx     = newErr("FFDecoderCopyCodecCtx", "unable to copy codec context for file %s: %v")
	FFDecoderOpenCodec        = newErr("FFDecoderOpenCodec", "unable to open codec for file %s: %v")
	FFDecoderFrameAlloc       = newErr("FFDecoderFrameAlloc", "call to AvFrameAlloc returned nil for file %s")
	FFDecoderMissingStream    = newErr("FFDecoderMissingStream", "failed to find a video stream for file %s")
)
