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
	FFDecoderCopyCodecCtx     = newErr("FFDecoderCopyCodecCtx", "unable to copy codec context for file %s: %v")
	FFDecoderFrameAlloc       = newErr("FFDecoderFrameAlloc", "call to AvFrameAlloc returned nil for file %s")
	FFDecoderMissingStream    = newErr("FFDecoderMissingStream", "failed to find a video stream for file %s")
	FFDecoderOpenCodec        = newErr("FFDecoderOpenCodec", "unable to open codec for file %s: %v")
	FFDecoderOpenInput        = newErr("FFDecodeOpenInput", "unable to open input for file %s: %v")
	FFDecoderStreamInfo       = newErr("FFDecoderStreamInfo", "unable to find stream info for file %s: %v")
	FFDecoderUnsupportedCodec = newErr("FFDecoderUnsupportedCodec", "unsupported codec in file %s: %v")
	GLShaderCompile           = newErr("GLFShaderCompile", "shader compilation failed: %v\n===== SOURCE =====\n%v")
	GLLinkProgram             = newErr("GLLinkProgram", "linking program %s failed: %v")
	NetTooManyArgs            = newErr("NetTooManyArgs", "too many arguments, expected %d but %d were provided")
	NetOutOfArgs              = newErr("NetOutOfArgs", "out of arguments, expected argument %d to be a %s")
	NetWrongArgType           = newErr("NetWrongArgType", "wrong type for argument %d, expected %s, was %T")
	SceneMissingWindowSource  = newErr("SceneMissingWindowSource", "unable to render scene, window source is missing")
	SceneMissingWindowProgram = newErr("SceneMissingWindowProgram", "unable to render scene, window program is missing")
	SceneRenderOrder          = newErr("SceneRenderOrder", "unable to determine source render order: %v")
	SourceDependencyLoop      = newErr("SourceDependencyLoop", "dependency loop found in sources: %s")
	SourceMissing             = newErr("SourceMissing", "source '%s' requires '%s', but no definition for '%s' was found. defined sources are: %v")
)
