package opengl

import (
	"github.com/dianelooney/gggv/internal/carbon"
	"github.com/go-gl/gl/all-core/gl"
)

func init() {
	carbon.Accum = gl.Accum
	carbon.ActiveProgramEXT = gl.ActiveProgramEXT

	carbon.ActiveShaderProgram = gl.ActiveShaderProgram
	carbon.ActiveShaderProgramEXT = gl.ActiveShaderProgramEXT

	carbon.ActiveTexture = gl.ActiveTexture

	carbon.AlphaFunc = gl.AlphaFunc
	carbon.ApplyFramebufferAttachmentCMAAINTEL = gl.ApplyFramebufferAttachmentCMAAINTEL

	carbon.AreTexturesResident = gl.AreTexturesResident

	carbon.ArrayElement = gl.ArrayElement

	carbon.AttachShader = gl.AttachShader

	carbon.Begin = gl.Begin

	carbon.BeginConditionalRender = gl.BeginConditionalRender
	carbon.BeginConditionalRenderNV = gl.BeginConditionalRenderNV
	carbon.BeginPerfMonitorAMD = gl.BeginPerfMonitorAMD
	carbon.BeginPerfQueryINTEL = gl.BeginPerfQueryINTEL

	carbon.BeginQuery = gl.BeginQuery
	carbon.BeginQueryIndexed = gl.BeginQueryIndexed

	carbon.BeginTransformFeedback = gl.BeginTransformFeedback

	carbon.BindAttribLocation = gl.BindAttribLocation

	carbon.BindBuffer = gl.BindBuffer

	carbon.BindBufferBase = gl.BindBufferBase

	carbon.BindBufferRange = gl.BindBufferRange

	carbon.BindBuffersBase = gl.BindBuffersBase

	carbon.BindBuffersRange = gl.BindBuffersRange

	carbon.BindFragDataLocation = gl.BindFragDataLocation

	carbon.BindFragDataLocationIndexed = gl.BindFragDataLocationIndexed

	carbon.BindFramebuffer = gl.BindFramebuffer

	carbon.BindImageTexture = gl.BindImageTexture

	carbon.BindImageTextures = gl.BindImageTextures
	carbon.BindMultiTextureEXT = gl.BindMultiTextureEXT

	carbon.BindProgramPipeline = gl.BindProgramPipeline
	carbon.BindProgramPipelineEXT = gl.BindProgramPipelineEXT

	carbon.BindRenderbuffer = gl.BindRenderbuffer

	carbon.BindSampler = gl.BindSampler

	carbon.BindSamplers = gl.BindSamplers

	carbon.BindTexture = gl.BindTexture

	carbon.BindTextureUnit = gl.BindTextureUnit

	carbon.BindTextures = gl.BindTextures

	carbon.BindTransformFeedback = gl.BindTransformFeedback

	carbon.BindVertexArray = gl.BindVertexArray

	carbon.BindVertexBuffer = gl.BindVertexBuffer

	carbon.BindVertexBuffers = gl.BindVertexBuffers

	carbon.Bitmap = gl.Bitmap
	carbon.BlendBarrierKHR = gl.BlendBarrierKHR
	carbon.BlendBarrierNV = gl.BlendBarrierNV

	carbon.BlendColor = gl.BlendColor

	carbon.BlendEquation = gl.BlendEquation

	carbon.BlendEquationSeparate = gl.BlendEquationSeparate
	carbon.BlendEquationSeparatei = gl.BlendEquationSeparatei
	carbon.BlendEquationSeparateiARB = gl.BlendEquationSeparateiARB
	carbon.BlendEquationi = gl.BlendEquationi
	carbon.BlendEquationiARB = gl.BlendEquationiARB

	carbon.BlendFunc = gl.BlendFunc

	carbon.BlendFuncSeparate = gl.BlendFuncSeparate
	carbon.BlendFuncSeparatei = gl.BlendFuncSeparatei
	carbon.BlendFuncSeparateiARB = gl.BlendFuncSeparateiARB
	carbon.BlendFunci = gl.BlendFunci
	carbon.BlendFunciARB = gl.BlendFunciARB
	carbon.BlendParameteriNV = gl.BlendParameteriNV

	carbon.BlitFramebuffer = gl.BlitFramebuffer

	carbon.BlitNamedFramebuffer = gl.BlitNamedFramebuffer
	carbon.BufferAddressRangeNV = gl.BufferAddressRangeNV

	carbon.BufferData = gl.BufferData
	carbon.BufferPageCommitmentARB = gl.BufferPageCommitmentARB

	carbon.BufferStorage = gl.BufferStorage

	carbon.BufferSubData = gl.BufferSubData
	carbon.CallCommandListNV = gl.CallCommandListNV

	carbon.CallList = gl.CallList

	carbon.CallLists = gl.CallLists

	carbon.CheckFramebufferStatus = gl.CheckFramebufferStatus

	carbon.CheckNamedFramebufferStatus = gl.CheckNamedFramebufferStatus
	carbon.CheckNamedFramebufferStatusEXT = gl.CheckNamedFramebufferStatusEXT

	carbon.ClampColor = gl.ClampColor

	carbon.Clear = gl.Clear

	carbon.ClearAccum = gl.ClearAccum

	carbon.ClearBufferData = gl.ClearBufferData

	carbon.ClearBufferSubData = gl.ClearBufferSubData
	carbon.ClearBufferfi = gl.ClearBufferfi
	carbon.ClearBufferfv = gl.ClearBufferfv
	carbon.ClearBufferiv = gl.ClearBufferiv
	carbon.ClearBufferuiv = gl.ClearBufferuiv

	carbon.ClearColor = gl.ClearColor

	carbon.ClearDepth = gl.ClearDepth

	carbon.ClearDepthf = gl.ClearDepthf

	carbon.ClearIndex = gl.ClearIndex

	carbon.ClearNamedBufferData = gl.ClearNamedBufferData
	carbon.ClearNamedBufferDataEXT = gl.ClearNamedBufferDataEXT

	carbon.ClearNamedBufferSubData = gl.ClearNamedBufferSubData
	carbon.ClearNamedBufferSubDataEXT = gl.ClearNamedBufferSubDataEXT
	carbon.ClearNamedFramebufferfi = gl.ClearNamedFramebufferfi
	carbon.ClearNamedFramebufferfv = gl.ClearNamedFramebufferfv
	carbon.ClearNamedFramebufferiv = gl.ClearNamedFramebufferiv
	carbon.ClearNamedFramebufferuiv = gl.ClearNamedFramebufferuiv

	carbon.ClearStencil = gl.ClearStencil

	carbon.ClearTexImage = gl.ClearTexImage

	carbon.ClearTexSubImage = gl.ClearTexSubImage

	carbon.ClientActiveTexture = gl.ClientActiveTexture
	carbon.ClientAttribDefaultEXT = gl.ClientAttribDefaultEXT

	carbon.ClientWaitSync = gl.ClientWaitSync

	carbon.ClipControl = gl.ClipControl

	carbon.ClipPlane = gl.ClipPlane
	carbon.Color3b = gl.Color3b
	carbon.Color3bv = gl.Color3bv
	carbon.Color3d = gl.Color3d
	carbon.Color3dv = gl.Color3dv
	carbon.Color3f = gl.Color3f
	carbon.Color3fv = gl.Color3fv
	carbon.Color3i = gl.Color3i
	carbon.Color3iv = gl.Color3iv
	carbon.Color3s = gl.Color3s
	carbon.Color3sv = gl.Color3sv
	carbon.Color3ub = gl.Color3ub
	carbon.Color3ubv = gl.Color3ubv
	carbon.Color3ui = gl.Color3ui
	carbon.Color3uiv = gl.Color3uiv
	carbon.Color3us = gl.Color3us
	carbon.Color3usv = gl.Color3usv
	carbon.Color4b = gl.Color4b
	carbon.Color4bv = gl.Color4bv
	carbon.Color4d = gl.Color4d
	carbon.Color4dv = gl.Color4dv
	carbon.Color4f = gl.Color4f
	carbon.Color4fv = gl.Color4fv
	carbon.Color4i = gl.Color4i
	carbon.Color4iv = gl.Color4iv
	carbon.Color4s = gl.Color4s
	carbon.Color4sv = gl.Color4sv
	carbon.Color4ub = gl.Color4ub
	carbon.Color4ubv = gl.Color4ubv
	carbon.Color4ui = gl.Color4ui
	carbon.Color4uiv = gl.Color4uiv
	carbon.Color4us = gl.Color4us
	carbon.Color4usv = gl.Color4usv
	carbon.ColorFormatNV = gl.ColorFormatNV
	carbon.ColorMask = gl.ColorMask
	carbon.ColorMaski = gl.ColorMaski

	carbon.ColorMaterial = gl.ColorMaterial

	carbon.ColorPointer = gl.ColorPointer
	carbon.CommandListSegmentsNV = gl.CommandListSegmentsNV
	carbon.CompileCommandListNV = gl.CompileCommandListNV

	carbon.CompileShader = gl.CompileShader
	carbon.CompileShaderIncludeARB = gl.CompileShaderIncludeARB
	carbon.CompressedMultiTexImage1DEXT = gl.CompressedMultiTexImage1DEXT
	carbon.CompressedMultiTexImage2DEXT = gl.CompressedMultiTexImage2DEXT
	carbon.CompressedMultiTexImage3DEXT = gl.CompressedMultiTexImage3DEXT
	carbon.CompressedMultiTexSubImage1DEXT = gl.CompressedMultiTexSubImage1DEXT
	carbon.CompressedMultiTexSubImage2DEXT = gl.CompressedMultiTexSubImage2DEXT
	carbon.CompressedMultiTexSubImage3DEXT = gl.CompressedMultiTexSubImage3DEXT

	carbon.CompressedTexImage1D = gl.CompressedTexImage1D

	carbon.CompressedTexImage2D = gl.CompressedTexImage2D

	carbon.CompressedTexImage3D = gl.CompressedTexImage3D

	carbon.CompressedTexSubImage1D = gl.CompressedTexSubImage1D

	carbon.CompressedTexSubImage2D = gl.CompressedTexSubImage2D

	carbon.CompressedTexSubImage3D = gl.CompressedTexSubImage3D
	carbon.CompressedTextureImage1DEXT = gl.CompressedTextureImage1DEXT
	carbon.CompressedTextureImage2DEXT = gl.CompressedTextureImage2DEXT
	carbon.CompressedTextureImage3DEXT = gl.CompressedTextureImage3DEXT

	carbon.CompressedTextureSubImage1D = gl.CompressedTextureSubImage1D
	carbon.CompressedTextureSubImage1DEXT = gl.CompressedTextureSubImage1DEXT

	carbon.CompressedTextureSubImage2D = gl.CompressedTextureSubImage2D
	carbon.CompressedTextureSubImage2DEXT = gl.CompressedTextureSubImage2DEXT

	carbon.CompressedTextureSubImage3D = gl.CompressedTextureSubImage3D
	carbon.CompressedTextureSubImage3DEXT = gl.CompressedTextureSubImage3DEXT
	carbon.ConservativeRasterParameterfNV = gl.ConservativeRasterParameterfNV
	carbon.ConservativeRasterParameteriNV = gl.ConservativeRasterParameteriNV

	carbon.CopyBufferSubData = gl.CopyBufferSubData

	carbon.CopyImageSubData = gl.CopyImageSubData
	carbon.CopyMultiTexImage1DEXT = gl.CopyMultiTexImage1DEXT
	carbon.CopyMultiTexImage2DEXT = gl.CopyMultiTexImage2DEXT
	carbon.CopyMultiTexSubImage1DEXT = gl.CopyMultiTexSubImage1DEXT
	carbon.CopyMultiTexSubImage2DEXT = gl.CopyMultiTexSubImage2DEXT
	carbon.CopyMultiTexSubImage3DEXT = gl.CopyMultiTexSubImage3DEXT

	carbon.CopyNamedBufferSubData = gl.CopyNamedBufferSubData
	carbon.CopyPathNV = gl.CopyPathNV

	carbon.CopyPixels = gl.CopyPixels

	carbon.CopyTexImage1D = gl.CopyTexImage1D

	carbon.CopyTexImage2D = gl.CopyTexImage2D

	carbon.CopyTexSubImage1D = gl.CopyTexSubImage1D

	carbon.CopyTexSubImage2D = gl.CopyTexSubImage2D

	carbon.CopyTexSubImage3D = gl.CopyTexSubImage3D
	carbon.CopyTextureImage1DEXT = gl.CopyTextureImage1DEXT
	carbon.CopyTextureImage2DEXT = gl.CopyTextureImage2DEXT

	carbon.CopyTextureSubImage1D = gl.CopyTextureSubImage1D
	carbon.CopyTextureSubImage1DEXT = gl.CopyTextureSubImage1DEXT

	carbon.CopyTextureSubImage2D = gl.CopyTextureSubImage2D
	carbon.CopyTextureSubImage2DEXT = gl.CopyTextureSubImage2DEXT

	carbon.CopyTextureSubImage3D = gl.CopyTextureSubImage3D
	carbon.CopyTextureSubImage3DEXT = gl.CopyTextureSubImage3DEXT
	carbon.CoverFillPathInstancedNV = gl.CoverFillPathInstancedNV
	carbon.CoverFillPathNV = gl.CoverFillPathNV
	carbon.CoverStrokePathInstancedNV = gl.CoverStrokePathInstancedNV
	carbon.CoverStrokePathNV = gl.CoverStrokePathNV
	carbon.CoverageModulationNV = gl.CoverageModulationNV
	carbon.CoverageModulationTableNV = gl.CoverageModulationTableNV

	carbon.CreateBuffers = gl.CreateBuffers
	carbon.CreateCommandListsNV = gl.CreateCommandListsNV

	carbon.CreateFramebuffers = gl.CreateFramebuffers
	carbon.CreatePerfQueryINTEL = gl.CreatePerfQueryINTEL

	carbon.CreateProgram = gl.CreateProgram

	carbon.CreateProgramPipelines = gl.CreateProgramPipelines

	carbon.CreateQueries = gl.CreateQueries

	carbon.CreateRenderbuffers = gl.CreateRenderbuffers

	carbon.CreateSamplers = gl.CreateSamplers

	carbon.CreateShader = gl.CreateShader
	carbon.CreateShaderProgramEXT = gl.CreateShaderProgramEXT

	carbon.CreateShaderProgramv = gl.CreateShaderProgramv
	carbon.CreateShaderProgramvEXT = gl.CreateShaderProgramvEXT
	carbon.CreateStatesNV = gl.CreateStatesNV

	carbon.CreateSyncFromCLeventARB = gl.CreateSyncFromCLeventARB

	carbon.CreateTextures = gl.CreateTextures

	carbon.CreateTransformFeedbacks = gl.CreateTransformFeedbacks

	carbon.CreateVertexArrays = gl.CreateVertexArrays

	carbon.CullFace = gl.CullFace

	// carbon.DebugMessageCallback = gl.DebugMessageCallback
	// carbon.DebugMessageCallbackARB = gl.DebugMessageCallbackARB
	// carbon.DebugMessageCallbackKHR = gl.DebugMessageCallbackKHR

	carbon.DebugMessageControl = gl.DebugMessageControl
	carbon.DebugMessageControlARB = gl.DebugMessageControlARB
	carbon.DebugMessageControlKHR = gl.DebugMessageControlKHR

	carbon.DebugMessageInsert = gl.DebugMessageInsert
	carbon.DebugMessageInsertARB = gl.DebugMessageInsertARB
	carbon.DebugMessageInsertKHR = gl.DebugMessageInsertKHR

	carbon.DeleteBuffers = gl.DeleteBuffers
	carbon.DeleteCommandListsNV = gl.DeleteCommandListsNV

	carbon.DeleteFramebuffers = gl.DeleteFramebuffers

	carbon.DeleteLists = gl.DeleteLists
	carbon.DeleteNamedStringARB = gl.DeleteNamedStringARB
	carbon.DeletePathsNV = gl.DeletePathsNV
	carbon.DeletePerfMonitorsAMD = gl.DeletePerfMonitorsAMD
	carbon.DeletePerfQueryINTEL = gl.DeletePerfQueryINTEL

	carbon.DeleteProgram = gl.DeleteProgram

	carbon.DeleteProgramPipelines = gl.DeleteProgramPipelines
	carbon.DeleteProgramPipelinesEXT = gl.DeleteProgramPipelinesEXT

	carbon.DeleteQueries = gl.DeleteQueries

	carbon.DeleteRenderbuffers = gl.DeleteRenderbuffers

	carbon.DeleteSamplers = gl.DeleteSamplers

	carbon.DeleteShader = gl.DeleteShader
	carbon.DeleteStatesNV = gl.DeleteStatesNV

	carbon.DeleteSync = gl.DeleteSync

	carbon.DeleteTextures = gl.DeleteTextures

	carbon.DeleteTransformFeedbacks = gl.DeleteTransformFeedbacks

	carbon.DeleteVertexArrays = gl.DeleteVertexArrays

	carbon.DepthFunc = gl.DepthFunc

	carbon.DepthMask = gl.DepthMask

	carbon.DepthRange = gl.DepthRange
	carbon.DepthRangeArrayv = gl.DepthRangeArrayv

	carbon.DepthRangeIndexed = gl.DepthRangeIndexed

	carbon.DepthRangef = gl.DepthRangef

	carbon.DetachShader = gl.DetachShader
	carbon.Disable = gl.Disable
	carbon.DisableClientState = gl.DisableClientState
	carbon.DisableClientStateIndexedEXT = gl.DisableClientStateIndexedEXT
	carbon.DisableClientStateiEXT = gl.DisableClientStateiEXT
	carbon.DisableIndexedEXT = gl.DisableIndexedEXT

	carbon.DisableVertexArrayAttrib = gl.DisableVertexArrayAttrib
	carbon.DisableVertexArrayAttribEXT = gl.DisableVertexArrayAttribEXT
	carbon.DisableVertexArrayEXT = gl.DisableVertexArrayEXT

	carbon.DisableVertexAttribArray = gl.DisableVertexAttribArray
	carbon.Disablei = gl.Disablei

	carbon.DispatchCompute = gl.DispatchCompute
	carbon.DispatchComputeGroupSizeARB = gl.DispatchComputeGroupSizeARB

	carbon.DispatchComputeIndirect = gl.DispatchComputeIndirect

	carbon.DrawArrays = gl.DrawArrays

	carbon.DrawArraysIndirect = gl.DrawArraysIndirect

	carbon.DrawArraysInstanced = gl.DrawArraysInstanced
	carbon.DrawArraysInstancedARB = gl.DrawArraysInstancedARB

	carbon.DrawArraysInstancedBaseInstance = gl.DrawArraysInstancedBaseInstance
	carbon.DrawArraysInstancedEXT = gl.DrawArraysInstancedEXT

	carbon.DrawBuffer = gl.DrawBuffer

	carbon.DrawBuffers = gl.DrawBuffers
	carbon.DrawCommandsAddressNV = gl.DrawCommandsAddressNV
	carbon.DrawCommandsNV = gl.DrawCommandsNV
	carbon.DrawCommandsStatesAddressNV = gl.DrawCommandsStatesAddressNV
	carbon.DrawCommandsStatesNV = gl.DrawCommandsStatesNV

	carbon.DrawElements = gl.DrawElements

	carbon.DrawElementsBaseVertex = gl.DrawElementsBaseVertex

	carbon.DrawElementsIndirect = gl.DrawElementsIndirect

	carbon.DrawElementsInstanced = gl.DrawElementsInstanced
	carbon.DrawElementsInstancedARB = gl.DrawElementsInstancedARB

	carbon.DrawElementsInstancedBaseInstance = gl.DrawElementsInstancedBaseInstance

	carbon.DrawElementsInstancedBaseVertex = gl.DrawElementsInstancedBaseVertex

	carbon.DrawElementsInstancedBaseVertexBaseInstance = gl.DrawElementsInstancedBaseVertexBaseInstance
	carbon.DrawElementsInstancedEXT = gl.DrawElementsInstancedEXT

	carbon.DrawPixels = gl.DrawPixels

	carbon.DrawRangeElements = gl.DrawRangeElements

	carbon.DrawRangeElementsBaseVertex = gl.DrawRangeElementsBaseVertex

	carbon.DrawTransformFeedback = gl.DrawTransformFeedback

	carbon.DrawTransformFeedbackInstanced = gl.DrawTransformFeedbackInstanced

	carbon.DrawTransformFeedbackStream = gl.DrawTransformFeedbackStream

	carbon.DrawTransformFeedbackStreamInstanced = gl.DrawTransformFeedbackStreamInstanced
	carbon.DrawVkImageNV = gl.DrawVkImageNV

	carbon.EGLImageTargetTexStorageEXT = gl.EGLImageTargetTexStorageEXT

	carbon.EGLImageTargetTextureStorageEXT = gl.EGLImageTargetTextureStorageEXT

	carbon.EdgeFlag = gl.EdgeFlag
	carbon.EdgeFlagFormatNV = gl.EdgeFlagFormatNV

	carbon.EdgeFlagPointer = gl.EdgeFlagPointer
	carbon.EdgeFlagv = gl.EdgeFlagv

	carbon.Enable = gl.Enable

	carbon.EnableClientState = gl.EnableClientState
	carbon.EnableClientStateIndexedEXT = gl.EnableClientStateIndexedEXT
	carbon.EnableClientStateiEXT = gl.EnableClientStateiEXT
	carbon.EnableIndexedEXT = gl.EnableIndexedEXT

	carbon.EnableVertexArrayAttrib = gl.EnableVertexArrayAttrib
	carbon.EnableVertexArrayAttribEXT = gl.EnableVertexArrayAttribEXT
	carbon.EnableVertexArrayEXT = gl.EnableVertexArrayEXT

	carbon.EnableVertexAttribArray = gl.EnableVertexAttribArray
	carbon.Enablei = gl.Enablei
	carbon.End = gl.End
	carbon.EndConditionalRender = gl.EndConditionalRender
	carbon.EndConditionalRenderNV = gl.EndConditionalRenderNV
	carbon.EndList = gl.EndList
	carbon.EndPerfMonitorAMD = gl.EndPerfMonitorAMD
	carbon.EndPerfQueryINTEL = gl.EndPerfQueryINTEL
	carbon.EndQuery = gl.EndQuery
	carbon.EndQueryIndexed = gl.EndQueryIndexed
	carbon.EndTransformFeedback = gl.EndTransformFeedback
	carbon.EvalCoord1d = gl.EvalCoord1d
	carbon.EvalCoord1dv = gl.EvalCoord1dv
	carbon.EvalCoord1f = gl.EvalCoord1f
	carbon.EvalCoord1fv = gl.EvalCoord1fv
	carbon.EvalCoord2d = gl.EvalCoord2d
	carbon.EvalCoord2dv = gl.EvalCoord2dv
	carbon.EvalCoord2f = gl.EvalCoord2f
	carbon.EvalCoord2fv = gl.EvalCoord2fv
	carbon.EvalMesh1 = gl.EvalMesh1
	carbon.EvalMesh2 = gl.EvalMesh2
	carbon.EvalPoint1 = gl.EvalPoint1
	carbon.EvalPoint2 = gl.EvalPoint2
	carbon.EvaluateDepthValuesARB = gl.EvaluateDepthValuesARB

	carbon.FeedbackBuffer = gl.FeedbackBuffer

	carbon.FenceSync = gl.FenceSync

	carbon.Finish = gl.Finish

	carbon.Flush = gl.Flush

	carbon.FlushMappedBufferRange = gl.FlushMappedBufferRange

	carbon.FlushMappedNamedBufferRange = gl.FlushMappedNamedBufferRange
	carbon.FlushMappedNamedBufferRangeEXT = gl.FlushMappedNamedBufferRangeEXT
	carbon.FogCoordFormatNV = gl.FogCoordFormatNV

	carbon.FogCoordPointer = gl.FogCoordPointer
	carbon.FogCoordd = gl.FogCoordd
	carbon.FogCoorddv = gl.FogCoorddv
	carbon.FogCoordf = gl.FogCoordf
	carbon.FogCoordfv = gl.FogCoordfv
	carbon.Fogf = gl.Fogf
	carbon.Fogfv = gl.Fogfv
	carbon.Fogi = gl.Fogi
	carbon.Fogiv = gl.Fogiv
	carbon.FragmentCoverageColorNV = gl.FragmentCoverageColorNV
	carbon.FramebufferDrawBufferEXT = gl.FramebufferDrawBufferEXT
	carbon.FramebufferDrawBuffersEXT = gl.FramebufferDrawBuffersEXT
	carbon.FramebufferFetchBarrierEXT = gl.FramebufferFetchBarrierEXT

	carbon.FramebufferParameteri = gl.FramebufferParameteri
	carbon.FramebufferReadBufferEXT = gl.FramebufferReadBufferEXT

	carbon.FramebufferRenderbuffer = gl.FramebufferRenderbuffer
	carbon.FramebufferSampleLocationsfvARB = gl.FramebufferSampleLocationsfvARB
	carbon.FramebufferSampleLocationsfvNV = gl.FramebufferSampleLocationsfvNV

	carbon.FramebufferTexture = gl.FramebufferTexture
	carbon.FramebufferTexture1D = gl.FramebufferTexture1D

	carbon.FramebufferTexture2D = gl.FramebufferTexture2D
	carbon.FramebufferTexture3D = gl.FramebufferTexture3D
	carbon.FramebufferTextureARB = gl.FramebufferTextureARB
	carbon.FramebufferTextureFaceARB = gl.FramebufferTextureFaceARB

	carbon.FramebufferTextureLayer = gl.FramebufferTextureLayer
	carbon.FramebufferTextureLayerARB = gl.FramebufferTextureLayerARB
	carbon.FramebufferTextureMultiviewOVR = gl.FramebufferTextureMultiviewOVR

	carbon.FrontFace = gl.FrontFace

	carbon.Frustum = gl.Frustum

	carbon.GenBuffers = gl.GenBuffers

	carbon.GenFramebuffers = gl.GenFramebuffers

	carbon.GenLists = gl.GenLists
	carbon.GenPathsNV = gl.GenPathsNV
	carbon.GenPerfMonitorsAMD = gl.GenPerfMonitorsAMD

	carbon.GenProgramPipelines = gl.GenProgramPipelines
	carbon.GenProgramPipelinesEXT = gl.GenProgramPipelinesEXT

	carbon.GenQueries = gl.GenQueries

	carbon.GenRenderbuffers = gl.GenRenderbuffers

	carbon.GenSamplers = gl.GenSamplers

	carbon.GenTextures = gl.GenTextures

	carbon.GenTransformFeedbacks = gl.GenTransformFeedbacks

	carbon.GenVertexArrays = gl.GenVertexArrays

	carbon.GenerateMipmap = gl.GenerateMipmap
	carbon.GenerateMultiTexMipmapEXT = gl.GenerateMultiTexMipmapEXT

	carbon.GenerateTextureMipmap = gl.GenerateTextureMipmap
	carbon.GenerateTextureMipmapEXT = gl.GenerateTextureMipmapEXT

	carbon.GetActiveAtomicCounterBufferiv = gl.GetActiveAtomicCounterBufferiv

	carbon.GetActiveAttrib = gl.GetActiveAttrib

	carbon.GetActiveSubroutineName = gl.GetActiveSubroutineName

	carbon.GetActiveSubroutineUniformName = gl.GetActiveSubroutineUniformName
	carbon.GetActiveSubroutineUniformiv = gl.GetActiveSubroutineUniformiv

	carbon.GetActiveUniform = gl.GetActiveUniform

	carbon.GetActiveUniformBlockName = gl.GetActiveUniformBlockName

	carbon.GetActiveUniformBlockiv = gl.GetActiveUniformBlockiv

	carbon.GetActiveUniformName = gl.GetActiveUniformName

	carbon.GetActiveUniformsiv = gl.GetActiveUniformsiv

	carbon.GetAttachedShaders = gl.GetAttachedShaders

	carbon.GetAttribLocation = gl.GetAttribLocation
	carbon.GetBooleanIndexedvEXT = gl.GetBooleanIndexedvEXT
	carbon.GetBooleani_v = gl.GetBooleani_v
	carbon.GetBooleanv = gl.GetBooleanv

	carbon.GetBufferParameteri64v = gl.GetBufferParameteri64v

	carbon.GetBufferParameteriv = gl.GetBufferParameteriv
	carbon.GetBufferParameterui64vNV = gl.GetBufferParameterui64vNV

	carbon.GetBufferPointerv = gl.GetBufferPointerv

	carbon.GetBufferSubData = gl.GetBufferSubData

	carbon.GetClipPlane = gl.GetClipPlane
	carbon.GetCommandHeaderNV = gl.GetCommandHeaderNV
	carbon.GetCompressedMultiTexImageEXT = gl.GetCompressedMultiTexImageEXT

	carbon.GetCompressedTexImage = gl.GetCompressedTexImage

	carbon.GetCompressedTextureImage = gl.GetCompressedTextureImage
	carbon.GetCompressedTextureImageEXT = gl.GetCompressedTextureImageEXT

	carbon.GetCompressedTextureSubImage = gl.GetCompressedTextureSubImage
	carbon.GetCoverageModulationTableNV = gl.GetCoverageModulationTableNV

	carbon.GetDebugMessageLog = gl.GetDebugMessageLog
	carbon.GetDebugMessageLogARB = gl.GetDebugMessageLogARB
	carbon.GetDebugMessageLogKHR = gl.GetDebugMessageLogKHR
	carbon.GetDoubleIndexedvEXT = gl.GetDoubleIndexedvEXT
	carbon.GetDoublei_v = gl.GetDoublei_v
	carbon.GetDoublei_vEXT = gl.GetDoublei_vEXT
	carbon.GetDoublev = gl.GetDoublev

	carbon.GetError = gl.GetError
	carbon.GetFirstPerfQueryIdINTEL = gl.GetFirstPerfQueryIdINTEL
	carbon.GetFloatIndexedvEXT = gl.GetFloatIndexedvEXT
	carbon.GetFloati_v = gl.GetFloati_v
	carbon.GetFloati_vEXT = gl.GetFloati_vEXT
	carbon.GetFloatv = gl.GetFloatv

	carbon.GetFragDataIndex = gl.GetFragDataIndex

	carbon.GetFragDataLocation = gl.GetFragDataLocation

	carbon.GetFramebufferAttachmentParameteriv = gl.GetFramebufferAttachmentParameteriv

	carbon.GetFramebufferParameteriv = gl.GetFramebufferParameteriv
	carbon.GetFramebufferParameterivEXT = gl.GetFramebufferParameterivEXT

	carbon.GetGraphicsResetStatus = gl.GetGraphicsResetStatus
	carbon.GetGraphicsResetStatusARB = gl.GetGraphicsResetStatusARB
	carbon.GetGraphicsResetStatusKHR = gl.GetGraphicsResetStatusKHR
	carbon.GetImageHandleARB = gl.GetImageHandleARB
	carbon.GetImageHandleNV = gl.GetImageHandleNV
	carbon.GetInteger64i_v = gl.GetInteger64i_v
	carbon.GetInteger64v = gl.GetInteger64v
	carbon.GetIntegerIndexedvEXT = gl.GetIntegerIndexedvEXT
	carbon.GetIntegeri_v = gl.GetIntegeri_v
	carbon.GetIntegerui64i_vNV = gl.GetIntegerui64i_vNV
	carbon.GetIntegerui64vNV = gl.GetIntegerui64vNV
	carbon.GetIntegerv = gl.GetIntegerv
	carbon.GetInternalformatSampleivNV = gl.GetInternalformatSampleivNV
	carbon.GetInternalformati64v = gl.GetInternalformati64v

	carbon.GetInternalformativ = gl.GetInternalformativ
	carbon.GetLightfv = gl.GetLightfv
	carbon.GetLightiv = gl.GetLightiv
	carbon.GetMapdv = gl.GetMapdv
	carbon.GetMapfv = gl.GetMapfv
	carbon.GetMapiv = gl.GetMapiv
	carbon.GetMaterialfv = gl.GetMaterialfv
	carbon.GetMaterialiv = gl.GetMaterialiv
	carbon.GetMultiTexEnvfvEXT = gl.GetMultiTexEnvfvEXT
	carbon.GetMultiTexEnvivEXT = gl.GetMultiTexEnvivEXT
	carbon.GetMultiTexGendvEXT = gl.GetMultiTexGendvEXT
	carbon.GetMultiTexGenfvEXT = gl.GetMultiTexGenfvEXT
	carbon.GetMultiTexGenivEXT = gl.GetMultiTexGenivEXT
	carbon.GetMultiTexImageEXT = gl.GetMultiTexImageEXT
	carbon.GetMultiTexLevelParameterfvEXT = gl.GetMultiTexLevelParameterfvEXT
	carbon.GetMultiTexLevelParameterivEXT = gl.GetMultiTexLevelParameterivEXT
	carbon.GetMultiTexParameterIivEXT = gl.GetMultiTexParameterIivEXT
	carbon.GetMultiTexParameterIuivEXT = gl.GetMultiTexParameterIuivEXT
	carbon.GetMultiTexParameterfvEXT = gl.GetMultiTexParameterfvEXT
	carbon.GetMultiTexParameterivEXT = gl.GetMultiTexParameterivEXT

	carbon.GetMultisamplefv = gl.GetMultisamplefv

	carbon.GetNamedBufferParameteri64v = gl.GetNamedBufferParameteri64v

	carbon.GetNamedBufferParameteriv = gl.GetNamedBufferParameteriv
	carbon.GetNamedBufferParameterivEXT = gl.GetNamedBufferParameterivEXT
	carbon.GetNamedBufferParameterui64vNV = gl.GetNamedBufferParameterui64vNV

	carbon.GetNamedBufferPointerv = gl.GetNamedBufferPointerv
	carbon.GetNamedBufferPointervEXT = gl.GetNamedBufferPointervEXT

	carbon.GetNamedBufferSubData = gl.GetNamedBufferSubData
	carbon.GetNamedBufferSubDataEXT = gl.GetNamedBufferSubDataEXT

	carbon.GetNamedFramebufferAttachmentParameteriv = gl.GetNamedFramebufferAttachmentParameteriv
	carbon.GetNamedFramebufferAttachmentParameterivEXT = gl.GetNamedFramebufferAttachmentParameterivEXT

	carbon.GetNamedFramebufferParameteriv = gl.GetNamedFramebufferParameteriv
	carbon.GetNamedFramebufferParameterivEXT = gl.GetNamedFramebufferParameterivEXT
	carbon.GetNamedProgramLocalParameterIivEXT = gl.GetNamedProgramLocalParameterIivEXT
	carbon.GetNamedProgramLocalParameterIuivEXT = gl.GetNamedProgramLocalParameterIuivEXT
	carbon.GetNamedProgramLocalParameterdvEXT = gl.GetNamedProgramLocalParameterdvEXT
	carbon.GetNamedProgramLocalParameterfvEXT = gl.GetNamedProgramLocalParameterfvEXT
	carbon.GetNamedProgramStringEXT = gl.GetNamedProgramStringEXT
	carbon.GetNamedProgramivEXT = gl.GetNamedProgramivEXT

	carbon.GetNamedRenderbufferParameteriv = gl.GetNamedRenderbufferParameteriv
	carbon.GetNamedRenderbufferParameterivEXT = gl.GetNamedRenderbufferParameterivEXT
	carbon.GetNamedStringARB = gl.GetNamedStringARB
	carbon.GetNamedStringivARB = gl.GetNamedStringivARB
	carbon.GetNextPerfQueryIdINTEL = gl.GetNextPerfQueryIdINTEL

	carbon.GetObjectLabel = gl.GetObjectLabel
	carbon.GetObjectLabelEXT = gl.GetObjectLabelEXT
	carbon.GetObjectLabelKHR = gl.GetObjectLabelKHR

	carbon.GetObjectPtrLabel = gl.GetObjectPtrLabel
	carbon.GetObjectPtrLabelKHR = gl.GetObjectPtrLabelKHR
	carbon.GetPathCommandsNV = gl.GetPathCommandsNV
	carbon.GetPathCoordsNV = gl.GetPathCoordsNV
	carbon.GetPathDashArrayNV = gl.GetPathDashArrayNV
	carbon.GetPathLengthNV = gl.GetPathLengthNV
	carbon.GetPathMetricRangeNV = gl.GetPathMetricRangeNV
	carbon.GetPathMetricsNV = gl.GetPathMetricsNV
	carbon.GetPathParameterfvNV = gl.GetPathParameterfvNV
	carbon.GetPathParameterivNV = gl.GetPathParameterivNV
	carbon.GetPathSpacingNV = gl.GetPathSpacingNV
	carbon.GetPerfCounterInfoINTEL = gl.GetPerfCounterInfoINTEL
	carbon.GetPerfMonitorCounterDataAMD = gl.GetPerfMonitorCounterDataAMD
	carbon.GetPerfMonitorCounterInfoAMD = gl.GetPerfMonitorCounterInfoAMD
	carbon.GetPerfMonitorCounterStringAMD = gl.GetPerfMonitorCounterStringAMD
	carbon.GetPerfMonitorCountersAMD = gl.GetPerfMonitorCountersAMD
	carbon.GetPerfMonitorGroupStringAMD = gl.GetPerfMonitorGroupStringAMD
	carbon.GetPerfMonitorGroupsAMD = gl.GetPerfMonitorGroupsAMD
	carbon.GetPerfQueryDataINTEL = gl.GetPerfQueryDataINTEL
	carbon.GetPerfQueryIdByNameINTEL = gl.GetPerfQueryIdByNameINTEL
	carbon.GetPerfQueryInfoINTEL = gl.GetPerfQueryInfoINTEL
	carbon.GetPixelMapfv = gl.GetPixelMapfv
	carbon.GetPixelMapuiv = gl.GetPixelMapuiv
	carbon.GetPixelMapusv = gl.GetPixelMapusv
	carbon.GetPointerIndexedvEXT = gl.GetPointerIndexedvEXT
	carbon.GetPointeri_vEXT = gl.GetPointeri_vEXT

	carbon.GetPointerv = gl.GetPointerv
	carbon.GetPointervKHR = gl.GetPointervKHR

	carbon.GetPolygonStipple = gl.GetPolygonStipple

	carbon.GetProgramBinary = gl.GetProgramBinary

	carbon.GetProgramInfoLog = gl.GetProgramInfoLog
	carbon.GetProgramInterfaceiv = gl.GetProgramInterfaceiv

	carbon.GetProgramPipelineInfoLog = gl.GetProgramPipelineInfoLog
	carbon.GetProgramPipelineInfoLogEXT = gl.GetProgramPipelineInfoLogEXT
	carbon.GetProgramPipelineiv = gl.GetProgramPipelineiv
	carbon.GetProgramPipelineivEXT = gl.GetProgramPipelineivEXT

	carbon.GetProgramResourceIndex = gl.GetProgramResourceIndex

	carbon.GetProgramResourceLocation = gl.GetProgramResourceLocation

	carbon.GetProgramResourceLocationIndex = gl.GetProgramResourceLocationIndex

	carbon.GetProgramResourceName = gl.GetProgramResourceName
	carbon.GetProgramResourcefvNV = gl.GetProgramResourcefvNV
	carbon.GetProgramResourceiv = gl.GetProgramResourceiv
	carbon.GetProgramStageiv = gl.GetProgramStageiv

	carbon.GetProgramiv = gl.GetProgramiv
	carbon.GetQueryBufferObjecti64v = gl.GetQueryBufferObjecti64v
	carbon.GetQueryBufferObjectiv = gl.GetQueryBufferObjectiv
	carbon.GetQueryBufferObjectui64v = gl.GetQueryBufferObjectui64v
	carbon.GetQueryBufferObjectuiv = gl.GetQueryBufferObjectuiv

	carbon.GetQueryIndexediv = gl.GetQueryIndexediv
	carbon.GetQueryObjecti64v = gl.GetQueryObjecti64v
	carbon.GetQueryObjectiv = gl.GetQueryObjectiv
	carbon.GetQueryObjectui64v = gl.GetQueryObjectui64v

	carbon.GetQueryObjectuiv = gl.GetQueryObjectuiv

	carbon.GetQueryiv = gl.GetQueryiv

	carbon.GetRenderbufferParameteriv = gl.GetRenderbufferParameteriv
	carbon.GetSamplerParameterIiv = gl.GetSamplerParameterIiv
	carbon.GetSamplerParameterIuiv = gl.GetSamplerParameterIuiv
	carbon.GetSamplerParameterfv = gl.GetSamplerParameterfv
	carbon.GetSamplerParameteriv = gl.GetSamplerParameteriv

	carbon.GetShaderInfoLog = gl.GetShaderInfoLog

	carbon.GetShaderPrecisionFormat = gl.GetShaderPrecisionFormat

	carbon.GetShaderSource = gl.GetShaderSource

	carbon.GetShaderiv = gl.GetShaderiv
	carbon.GetStageIndexNV = gl.GetStageIndexNV

	carbon.GetString = gl.GetString
	carbon.GetStringi = gl.GetStringi

	carbon.GetSubroutineIndex = gl.GetSubroutineIndex

	carbon.GetSubroutineUniformLocation = gl.GetSubroutineUniformLocation

	carbon.GetSynciv = gl.GetSynciv
	carbon.GetTexEnvfv = gl.GetTexEnvfv
	carbon.GetTexEnviv = gl.GetTexEnviv
	carbon.GetTexGendv = gl.GetTexGendv
	carbon.GetTexGenfv = gl.GetTexGenfv
	carbon.GetTexGeniv = gl.GetTexGeniv

	carbon.GetTexImage = gl.GetTexImage
	carbon.GetTexLevelParameterfv = gl.GetTexLevelParameterfv
	carbon.GetTexLevelParameteriv = gl.GetTexLevelParameteriv
	carbon.GetTexParameterIiv = gl.GetTexParameterIiv
	carbon.GetTexParameterIuiv = gl.GetTexParameterIuiv
	carbon.GetTexParameterfv = gl.GetTexParameterfv
	carbon.GetTexParameteriv = gl.GetTexParameteriv
	carbon.GetTextureHandleARB = gl.GetTextureHandleARB
	carbon.GetTextureHandleNV = gl.GetTextureHandleNV

	carbon.GetTextureImage = gl.GetTextureImage
	carbon.GetTextureImageEXT = gl.GetTextureImageEXT
	carbon.GetTextureLevelParameterfv = gl.GetTextureLevelParameterfv
	carbon.GetTextureLevelParameterfvEXT = gl.GetTextureLevelParameterfvEXT
	carbon.GetTextureLevelParameteriv = gl.GetTextureLevelParameteriv
	carbon.GetTextureLevelParameterivEXT = gl.GetTextureLevelParameterivEXT
	carbon.GetTextureParameterIiv = gl.GetTextureParameterIiv
	carbon.GetTextureParameterIivEXT = gl.GetTextureParameterIivEXT
	carbon.GetTextureParameterIuiv = gl.GetTextureParameterIuiv
	carbon.GetTextureParameterIuivEXT = gl.GetTextureParameterIuivEXT
	carbon.GetTextureParameterfv = gl.GetTextureParameterfv
	carbon.GetTextureParameterfvEXT = gl.GetTextureParameterfvEXT
	carbon.GetTextureParameteriv = gl.GetTextureParameteriv
	carbon.GetTextureParameterivEXT = gl.GetTextureParameterivEXT
	carbon.GetTextureSamplerHandleARB = gl.GetTextureSamplerHandleARB
	carbon.GetTextureSamplerHandleNV = gl.GetTextureSamplerHandleNV

	carbon.GetTextureSubImage = gl.GetTextureSubImage

	carbon.GetTransformFeedbackVarying = gl.GetTransformFeedbackVarying
	carbon.GetTransformFeedbacki64_v = gl.GetTransformFeedbacki64_v
	carbon.GetTransformFeedbacki_v = gl.GetTransformFeedbacki_v

	carbon.GetTransformFeedbackiv = gl.GetTransformFeedbackiv

	carbon.GetUniformBlockIndex = gl.GetUniformBlockIndex

	carbon.GetUniformIndices = gl.GetUniformIndices

	carbon.GetUniformLocation = gl.GetUniformLocation
	carbon.GetUniformSubroutineuiv = gl.GetUniformSubroutineuiv
	carbon.GetUniformdv = gl.GetUniformdv

	carbon.GetUniformfv = gl.GetUniformfv
	carbon.GetUniformi64vARB = gl.GetUniformi64vARB
	carbon.GetUniformi64vNV = gl.GetUniformi64vNV

	carbon.GetUniformiv = gl.GetUniformiv
	carbon.GetUniformui64vARB = gl.GetUniformui64vARB
	carbon.GetUniformui64vNV = gl.GetUniformui64vNV
	carbon.GetUniformuiv = gl.GetUniformuiv
	carbon.GetVertexArrayIndexed64iv = gl.GetVertexArrayIndexed64iv
	carbon.GetVertexArrayIndexediv = gl.GetVertexArrayIndexediv
	carbon.GetVertexArrayIntegeri_vEXT = gl.GetVertexArrayIntegeri_vEXT
	carbon.GetVertexArrayIntegervEXT = gl.GetVertexArrayIntegervEXT
	carbon.GetVertexArrayPointeri_vEXT = gl.GetVertexArrayPointeri_vEXT
	carbon.GetVertexArrayPointervEXT = gl.GetVertexArrayPointervEXT

	carbon.GetVertexArrayiv = gl.GetVertexArrayiv

	carbon.GetVertexAttribIiv = gl.GetVertexAttribIiv

	carbon.GetVertexAttribIuiv = gl.GetVertexAttribIuiv

	carbon.GetVertexAttribLdv = gl.GetVertexAttribLdv
	carbon.GetVertexAttribLi64vNV = gl.GetVertexAttribLi64vNV
	carbon.GetVertexAttribLui64vARB = gl.GetVertexAttribLui64vARB
	carbon.GetVertexAttribLui64vNV = gl.GetVertexAttribLui64vNV

	carbon.GetVertexAttribPointerv = gl.GetVertexAttribPointerv

	carbon.GetVertexAttribdv = gl.GetVertexAttribdv

	carbon.GetVertexAttribfv = gl.GetVertexAttribfv

	carbon.GetVertexAttribiv = gl.GetVertexAttribiv

	carbon.GetVkProcAddrNV = gl.GetVkProcAddrNV

	carbon.GetnCompressedTexImage = gl.GetnCompressedTexImage
	carbon.GetnCompressedTexImageARB = gl.GetnCompressedTexImageARB

	carbon.GetnTexImage = gl.GetnTexImage
	carbon.GetnTexImageARB = gl.GetnTexImageARB
	carbon.GetnUniformdv = gl.GetnUniformdv
	carbon.GetnUniformdvARB = gl.GetnUniformdvARB
	carbon.GetnUniformfv = gl.GetnUniformfv
	carbon.GetnUniformfvARB = gl.GetnUniformfvARB
	carbon.GetnUniformfvKHR = gl.GetnUniformfvKHR
	carbon.GetnUniformi64vARB = gl.GetnUniformi64vARB
	carbon.GetnUniformiv = gl.GetnUniformiv
	carbon.GetnUniformivARB = gl.GetnUniformivARB
	carbon.GetnUniformivKHR = gl.GetnUniformivKHR
	carbon.GetnUniformui64vARB = gl.GetnUniformui64vARB
	carbon.GetnUniformuiv = gl.GetnUniformuiv
	carbon.GetnUniformuivARB = gl.GetnUniformuivARB
	carbon.GetnUniformuivKHR = gl.GetnUniformuivKHR

	carbon.Hint = gl.Hint
	carbon.IndexFormatNV = gl.IndexFormatNV

	carbon.IndexMask = gl.IndexMask

	carbon.IndexPointer = gl.IndexPointer
	carbon.Indexd = gl.Indexd
	carbon.Indexdv = gl.Indexdv
	carbon.Indexf = gl.Indexf
	carbon.Indexfv = gl.Indexfv
	carbon.Indexi = gl.Indexi
	carbon.Indexiv = gl.Indexiv
	carbon.Indexs = gl.Indexs
	carbon.Indexsv = gl.Indexsv
	carbon.Indexub = gl.Indexub
	carbon.Indexubv = gl.Indexubv

	carbon.InitNames = gl.InitNames
	carbon.InsertEventMarkerEXT = gl.InsertEventMarkerEXT

	carbon.InterleavedArrays = gl.InterleavedArrays
	carbon.InterpolatePathsNV = gl.InterpolatePathsNV

	carbon.InvalidateBufferData = gl.InvalidateBufferData

	carbon.InvalidateBufferSubData = gl.InvalidateBufferSubData

	carbon.InvalidateFramebuffer = gl.InvalidateFramebuffer

	carbon.InvalidateNamedFramebufferData = gl.InvalidateNamedFramebufferData

	carbon.InvalidateNamedFramebufferSubData = gl.InvalidateNamedFramebufferSubData

	carbon.InvalidateSubFramebuffer = gl.InvalidateSubFramebuffer

	carbon.InvalidateTexImage = gl.InvalidateTexImage

	carbon.InvalidateTexSubImage = gl.InvalidateTexSubImage

	carbon.IsBuffer = gl.IsBuffer
	carbon.IsBufferResidentNV = gl.IsBufferResidentNV
	carbon.IsCommandListNV = gl.IsCommandListNV
	carbon.IsEnabled = gl.IsEnabled
	carbon.IsEnabledIndexedEXT = gl.IsEnabledIndexedEXT
	carbon.IsEnabledi = gl.IsEnabledi

	carbon.IsFramebuffer = gl.IsFramebuffer
	carbon.IsImageHandleResidentARB = gl.IsImageHandleResidentARB
	carbon.IsImageHandleResidentNV = gl.IsImageHandleResidentNV

	carbon.IsList = gl.IsList
	carbon.IsNamedBufferResidentNV = gl.IsNamedBufferResidentNV
	carbon.IsNamedStringARB = gl.IsNamedStringARB
	carbon.IsPathNV = gl.IsPathNV
	carbon.IsPointInFillPathNV = gl.IsPointInFillPathNV
	carbon.IsPointInStrokePathNV = gl.IsPointInStrokePathNV

	carbon.IsProgram = gl.IsProgram

	carbon.IsProgramPipeline = gl.IsProgramPipeline
	carbon.IsProgramPipelineEXT = gl.IsProgramPipelineEXT

	carbon.IsQuery = gl.IsQuery

	carbon.IsRenderbuffer = gl.IsRenderbuffer

	carbon.IsSampler = gl.IsSampler

	carbon.IsShader = gl.IsShader
	carbon.IsStateNV = gl.IsStateNV

	carbon.IsSync = gl.IsSync

	carbon.IsTexture = gl.IsTexture
	carbon.IsTextureHandleResidentARB = gl.IsTextureHandleResidentARB
	carbon.IsTextureHandleResidentNV = gl.IsTextureHandleResidentNV

	carbon.IsTransformFeedback = gl.IsTransformFeedback

	carbon.IsVertexArray = gl.IsVertexArray
	carbon.LabelObjectEXT = gl.LabelObjectEXT
	carbon.LightModelf = gl.LightModelf
	carbon.LightModelfv = gl.LightModelfv
	carbon.LightModeli = gl.LightModeli
	carbon.LightModeliv = gl.LightModeliv
	carbon.Lightf = gl.Lightf
	carbon.Lightfv = gl.Lightfv
	carbon.Lighti = gl.Lighti
	carbon.Lightiv = gl.Lightiv

	carbon.LineStipple = gl.LineStipple

	carbon.LineWidth = gl.LineWidth

	carbon.LinkProgram = gl.LinkProgram

	carbon.ListBase = gl.ListBase
	carbon.ListDrawCommandsStatesClientNV = gl.ListDrawCommandsStatesClientNV

	carbon.LoadIdentity = gl.LoadIdentity
	carbon.LoadMatrixd = gl.LoadMatrixd
	carbon.LoadMatrixf = gl.LoadMatrixf

	carbon.LoadName = gl.LoadName
	carbon.LoadTransposeMatrixd = gl.LoadTransposeMatrixd
	carbon.LoadTransposeMatrixf = gl.LoadTransposeMatrixf

	carbon.LogicOp = gl.LogicOp
	carbon.MakeBufferNonResidentNV = gl.MakeBufferNonResidentNV
	carbon.MakeBufferResidentNV = gl.MakeBufferResidentNV
	carbon.MakeImageHandleNonResidentARB = gl.MakeImageHandleNonResidentARB
	carbon.MakeImageHandleNonResidentNV = gl.MakeImageHandleNonResidentNV
	carbon.MakeImageHandleResidentARB = gl.MakeImageHandleResidentARB
	carbon.MakeImageHandleResidentNV = gl.MakeImageHandleResidentNV
	carbon.MakeNamedBufferNonResidentNV = gl.MakeNamedBufferNonResidentNV
	carbon.MakeNamedBufferResidentNV = gl.MakeNamedBufferResidentNV
	carbon.MakeTextureHandleNonResidentARB = gl.MakeTextureHandleNonResidentARB
	carbon.MakeTextureHandleNonResidentNV = gl.MakeTextureHandleNonResidentNV
	carbon.MakeTextureHandleResidentARB = gl.MakeTextureHandleResidentARB
	carbon.MakeTextureHandleResidentNV = gl.MakeTextureHandleResidentNV
	carbon.Map1d = gl.Map1d
	carbon.Map1f = gl.Map1f
	carbon.Map2d = gl.Map2d
	carbon.Map2f = gl.Map2f

	carbon.MapBuffer = gl.MapBuffer

	carbon.MapBufferRange = gl.MapBufferRange
	carbon.MapGrid1d = gl.MapGrid1d
	carbon.MapGrid1f = gl.MapGrid1f
	carbon.MapGrid2d = gl.MapGrid2d
	carbon.MapGrid2f = gl.MapGrid2f

	carbon.MapNamedBuffer = gl.MapNamedBuffer
	carbon.MapNamedBufferEXT = gl.MapNamedBufferEXT

	carbon.MapNamedBufferRange = gl.MapNamedBufferRange
	carbon.MapNamedBufferRangeEXT = gl.MapNamedBufferRangeEXT
	carbon.Materialf = gl.Materialf
	carbon.Materialfv = gl.Materialfv
	carbon.Materiali = gl.Materiali
	carbon.Materialiv = gl.Materialiv
	carbon.MatrixFrustumEXT = gl.MatrixFrustumEXT
	carbon.MatrixLoad3x2fNV = gl.MatrixLoad3x2fNV
	carbon.MatrixLoad3x3fNV = gl.MatrixLoad3x3fNV
	carbon.MatrixLoadIdentityEXT = gl.MatrixLoadIdentityEXT
	carbon.MatrixLoadTranspose3x3fNV = gl.MatrixLoadTranspose3x3fNV
	carbon.MatrixLoadTransposedEXT = gl.MatrixLoadTransposedEXT
	carbon.MatrixLoadTransposefEXT = gl.MatrixLoadTransposefEXT
	carbon.MatrixLoaddEXT = gl.MatrixLoaddEXT
	carbon.MatrixLoadfEXT = gl.MatrixLoadfEXT

	carbon.MatrixMode = gl.MatrixMode
	carbon.MatrixMult3x2fNV = gl.MatrixMult3x2fNV
	carbon.MatrixMult3x3fNV = gl.MatrixMult3x3fNV
	carbon.MatrixMultTranspose3x3fNV = gl.MatrixMultTranspose3x3fNV
	carbon.MatrixMultTransposedEXT = gl.MatrixMultTransposedEXT
	carbon.MatrixMultTransposefEXT = gl.MatrixMultTransposefEXT
	carbon.MatrixMultdEXT = gl.MatrixMultdEXT
	carbon.MatrixMultfEXT = gl.MatrixMultfEXT
	carbon.MatrixOrthoEXT = gl.MatrixOrthoEXT
	carbon.MatrixPopEXT = gl.MatrixPopEXT
	carbon.MatrixPushEXT = gl.MatrixPushEXT
	carbon.MatrixRotatedEXT = gl.MatrixRotatedEXT
	carbon.MatrixRotatefEXT = gl.MatrixRotatefEXT
	carbon.MatrixScaledEXT = gl.MatrixScaledEXT
	carbon.MatrixScalefEXT = gl.MatrixScalefEXT
	carbon.MatrixTranslatedEXT = gl.MatrixTranslatedEXT
	carbon.MatrixTranslatefEXT = gl.MatrixTranslatefEXT
	carbon.MaxShaderCompilerThreadsARB = gl.MaxShaderCompilerThreadsARB
	carbon.MaxShaderCompilerThreadsKHR = gl.MaxShaderCompilerThreadsKHR

	carbon.MemoryBarrier = gl.MemoryBarrier
	carbon.MemoryBarrierByRegion = gl.MemoryBarrierByRegion

	carbon.MinSampleShading = gl.MinSampleShading
	carbon.MinSampleShadingARB = gl.MinSampleShadingARB
	carbon.MultMatrixd = gl.MultMatrixd
	carbon.MultMatrixf = gl.MultMatrixf
	carbon.MultTransposeMatrixd = gl.MultTransposeMatrixd
	carbon.MultTransposeMatrixf = gl.MultTransposeMatrixf

	carbon.MultiDrawArrays = gl.MultiDrawArrays

	carbon.MultiDrawArraysIndirect = gl.MultiDrawArraysIndirect
	carbon.MultiDrawArraysIndirectBindlessCountNV = gl.MultiDrawArraysIndirectBindlessCountNV
	carbon.MultiDrawArraysIndirectBindlessNV = gl.MultiDrawArraysIndirectBindlessNV
	carbon.MultiDrawArraysIndirectCount = gl.MultiDrawArraysIndirectCount
	carbon.MultiDrawArraysIndirectCountARB = gl.MultiDrawArraysIndirectCountARB

	carbon.MultiDrawElements = gl.MultiDrawElements

	carbon.MultiDrawElementsBaseVertex = gl.MultiDrawElementsBaseVertex

	carbon.MultiDrawElementsIndirect = gl.MultiDrawElementsIndirect
	carbon.MultiDrawElementsIndirectBindlessCountNV = gl.MultiDrawElementsIndirectBindlessCountNV
	carbon.MultiDrawElementsIndirectBindlessNV = gl.MultiDrawElementsIndirectBindlessNV
	carbon.MultiDrawElementsIndirectCount = gl.MultiDrawElementsIndirectCount
	carbon.MultiDrawElementsIndirectCountARB = gl.MultiDrawElementsIndirectCountARB
	carbon.MultiTexBufferEXT = gl.MultiTexBufferEXT
	carbon.MultiTexCoord1d = gl.MultiTexCoord1d
	carbon.MultiTexCoord1dv = gl.MultiTexCoord1dv
	carbon.MultiTexCoord1f = gl.MultiTexCoord1f
	carbon.MultiTexCoord1fv = gl.MultiTexCoord1fv
	carbon.MultiTexCoord1i = gl.MultiTexCoord1i
	carbon.MultiTexCoord1iv = gl.MultiTexCoord1iv
	carbon.MultiTexCoord1s = gl.MultiTexCoord1s
	carbon.MultiTexCoord1sv = gl.MultiTexCoord1sv
	carbon.MultiTexCoord2d = gl.MultiTexCoord2d
	carbon.MultiTexCoord2dv = gl.MultiTexCoord2dv
	carbon.MultiTexCoord2f = gl.MultiTexCoord2f
	carbon.MultiTexCoord2fv = gl.MultiTexCoord2fv
	carbon.MultiTexCoord2i = gl.MultiTexCoord2i
	carbon.MultiTexCoord2iv = gl.MultiTexCoord2iv
	carbon.MultiTexCoord2s = gl.MultiTexCoord2s
	carbon.MultiTexCoord2sv = gl.MultiTexCoord2sv
	carbon.MultiTexCoord3d = gl.MultiTexCoord3d
	carbon.MultiTexCoord3dv = gl.MultiTexCoord3dv
	carbon.MultiTexCoord3f = gl.MultiTexCoord3f
	carbon.MultiTexCoord3fv = gl.MultiTexCoord3fv
	carbon.MultiTexCoord3i = gl.MultiTexCoord3i
	carbon.MultiTexCoord3iv = gl.MultiTexCoord3iv
	carbon.MultiTexCoord3s = gl.MultiTexCoord3s
	carbon.MultiTexCoord3sv = gl.MultiTexCoord3sv
	carbon.MultiTexCoord4d = gl.MultiTexCoord4d
	carbon.MultiTexCoord4dv = gl.MultiTexCoord4dv
	carbon.MultiTexCoord4f = gl.MultiTexCoord4f
	carbon.MultiTexCoord4fv = gl.MultiTexCoord4fv
	carbon.MultiTexCoord4i = gl.MultiTexCoord4i
	carbon.MultiTexCoord4iv = gl.MultiTexCoord4iv
	carbon.MultiTexCoord4s = gl.MultiTexCoord4s
	carbon.MultiTexCoord4sv = gl.MultiTexCoord4sv
	carbon.MultiTexCoordPointerEXT = gl.MultiTexCoordPointerEXT
	carbon.MultiTexEnvfEXT = gl.MultiTexEnvfEXT
	carbon.MultiTexEnvfvEXT = gl.MultiTexEnvfvEXT
	carbon.MultiTexEnviEXT = gl.MultiTexEnviEXT
	carbon.MultiTexEnvivEXT = gl.MultiTexEnvivEXT
	carbon.MultiTexGendEXT = gl.MultiTexGendEXT
	carbon.MultiTexGendvEXT = gl.MultiTexGendvEXT
	carbon.MultiTexGenfEXT = gl.MultiTexGenfEXT
	carbon.MultiTexGenfvEXT = gl.MultiTexGenfvEXT
	carbon.MultiTexGeniEXT = gl.MultiTexGeniEXT
	carbon.MultiTexGenivEXT = gl.MultiTexGenivEXT
	carbon.MultiTexImage1DEXT = gl.MultiTexImage1DEXT
	carbon.MultiTexImage2DEXT = gl.MultiTexImage2DEXT
	carbon.MultiTexImage3DEXT = gl.MultiTexImage3DEXT
	carbon.MultiTexParameterIivEXT = gl.MultiTexParameterIivEXT
	carbon.MultiTexParameterIuivEXT = gl.MultiTexParameterIuivEXT
	carbon.MultiTexParameterfEXT = gl.MultiTexParameterfEXT
	carbon.MultiTexParameterfvEXT = gl.MultiTexParameterfvEXT
	carbon.MultiTexParameteriEXT = gl.MultiTexParameteriEXT
	carbon.MultiTexParameterivEXT = gl.MultiTexParameterivEXT
	carbon.MultiTexRenderbufferEXT = gl.MultiTexRenderbufferEXT
	carbon.MultiTexSubImage1DEXT = gl.MultiTexSubImage1DEXT
	carbon.MultiTexSubImage2DEXT = gl.MultiTexSubImage2DEXT
	carbon.MultiTexSubImage3DEXT = gl.MultiTexSubImage3DEXT

	carbon.NamedBufferData = gl.NamedBufferData
	carbon.NamedBufferDataEXT = gl.NamedBufferDataEXT
	carbon.NamedBufferPageCommitmentARB = gl.NamedBufferPageCommitmentARB
	carbon.NamedBufferPageCommitmentEXT = gl.NamedBufferPageCommitmentEXT

	carbon.NamedBufferStorage = gl.NamedBufferStorage
	carbon.NamedBufferStorageEXT = gl.NamedBufferStorageEXT

	carbon.NamedBufferSubData = gl.NamedBufferSubData
	carbon.NamedBufferSubDataEXT = gl.NamedBufferSubDataEXT
	carbon.NamedCopyBufferSubDataEXT = gl.NamedCopyBufferSubDataEXT

	carbon.NamedFramebufferDrawBuffer = gl.NamedFramebufferDrawBuffer

	carbon.NamedFramebufferDrawBuffers = gl.NamedFramebufferDrawBuffers

	carbon.NamedFramebufferParameteri = gl.NamedFramebufferParameteri
	carbon.NamedFramebufferParameteriEXT = gl.NamedFramebufferParameteriEXT

	carbon.NamedFramebufferReadBuffer = gl.NamedFramebufferReadBuffer

	carbon.NamedFramebufferRenderbuffer = gl.NamedFramebufferRenderbuffer
	carbon.NamedFramebufferRenderbufferEXT = gl.NamedFramebufferRenderbufferEXT
	carbon.NamedFramebufferSampleLocationsfvARB = gl.NamedFramebufferSampleLocationsfvARB
	carbon.NamedFramebufferSampleLocationsfvNV = gl.NamedFramebufferSampleLocationsfvNV
	carbon.NamedFramebufferTexture = gl.NamedFramebufferTexture
	carbon.NamedFramebufferTexture1DEXT = gl.NamedFramebufferTexture1DEXT
	carbon.NamedFramebufferTexture2DEXT = gl.NamedFramebufferTexture2DEXT
	carbon.NamedFramebufferTexture3DEXT = gl.NamedFramebufferTexture3DEXT
	carbon.NamedFramebufferTextureEXT = gl.NamedFramebufferTextureEXT
	carbon.NamedFramebufferTextureFaceEXT = gl.NamedFramebufferTextureFaceEXT

	carbon.NamedFramebufferTextureLayer = gl.NamedFramebufferTextureLayer
	carbon.NamedFramebufferTextureLayerEXT = gl.NamedFramebufferTextureLayerEXT
	carbon.NamedProgramLocalParameter4dEXT = gl.NamedProgramLocalParameter4dEXT
	carbon.NamedProgramLocalParameter4dvEXT = gl.NamedProgramLocalParameter4dvEXT
	carbon.NamedProgramLocalParameter4fEXT = gl.NamedProgramLocalParameter4fEXT
	carbon.NamedProgramLocalParameter4fvEXT = gl.NamedProgramLocalParameter4fvEXT
	carbon.NamedProgramLocalParameterI4iEXT = gl.NamedProgramLocalParameterI4iEXT
	carbon.NamedProgramLocalParameterI4ivEXT = gl.NamedProgramLocalParameterI4ivEXT
	carbon.NamedProgramLocalParameterI4uiEXT = gl.NamedProgramLocalParameterI4uiEXT
	carbon.NamedProgramLocalParameterI4uivEXT = gl.NamedProgramLocalParameterI4uivEXT
	carbon.NamedProgramLocalParameters4fvEXT = gl.NamedProgramLocalParameters4fvEXT
	carbon.NamedProgramLocalParametersI4ivEXT = gl.NamedProgramLocalParametersI4ivEXT
	carbon.NamedProgramLocalParametersI4uivEXT = gl.NamedProgramLocalParametersI4uivEXT
	carbon.NamedProgramStringEXT = gl.NamedProgramStringEXT

	carbon.NamedRenderbufferStorage = gl.NamedRenderbufferStorage
	carbon.NamedRenderbufferStorageEXT = gl.NamedRenderbufferStorageEXT

	carbon.NamedRenderbufferStorageMultisample = gl.NamedRenderbufferStorageMultisample
	carbon.NamedRenderbufferStorageMultisampleCoverageEXT = gl.NamedRenderbufferStorageMultisampleCoverageEXT
	carbon.NamedRenderbufferStorageMultisampleEXT = gl.NamedRenderbufferStorageMultisampleEXT
	carbon.NamedStringARB = gl.NamedStringARB

	carbon.NewList = gl.NewList
	carbon.Normal3b = gl.Normal3b
	carbon.Normal3bv = gl.Normal3bv
	carbon.Normal3d = gl.Normal3d
	carbon.Normal3dv = gl.Normal3dv
	carbon.Normal3f = gl.Normal3f
	carbon.Normal3fv = gl.Normal3fv
	carbon.Normal3i = gl.Normal3i
	carbon.Normal3iv = gl.Normal3iv
	carbon.Normal3s = gl.Normal3s
	carbon.Normal3sv = gl.Normal3sv
	carbon.NormalFormatNV = gl.NormalFormatNV

	carbon.NormalPointer = gl.NormalPointer

	carbon.ObjectLabel = gl.ObjectLabel
	carbon.ObjectLabelKHR = gl.ObjectLabelKHR

	carbon.ObjectPtrLabel = gl.ObjectPtrLabel
	carbon.ObjectPtrLabelKHR = gl.ObjectPtrLabelKHR

	carbon.Ortho = gl.Ortho

	carbon.PassThrough = gl.PassThrough
	carbon.PatchParameterfv = gl.PatchParameterfv

	carbon.PatchParameteri = gl.PatchParameteri
	carbon.PathCommandsNV = gl.PathCommandsNV
	carbon.PathCoordsNV = gl.PathCoordsNV
	carbon.PathCoverDepthFuncNV = gl.PathCoverDepthFuncNV
	carbon.PathDashArrayNV = gl.PathDashArrayNV
	carbon.PathGlyphIndexArrayNV = gl.PathGlyphIndexArrayNV
	carbon.PathGlyphIndexRangeNV = gl.PathGlyphIndexRangeNV
	carbon.PathGlyphRangeNV = gl.PathGlyphRangeNV
	carbon.PathGlyphsNV = gl.PathGlyphsNV
	carbon.PathMemoryGlyphIndexArrayNV = gl.PathMemoryGlyphIndexArrayNV
	carbon.PathParameterfNV = gl.PathParameterfNV
	carbon.PathParameterfvNV = gl.PathParameterfvNV
	carbon.PathParameteriNV = gl.PathParameteriNV
	carbon.PathParameterivNV = gl.PathParameterivNV
	carbon.PathStencilDepthOffsetNV = gl.PathStencilDepthOffsetNV
	carbon.PathStencilFuncNV = gl.PathStencilFuncNV
	carbon.PathStringNV = gl.PathStringNV
	carbon.PathSubCommandsNV = gl.PathSubCommandsNV
	carbon.PathSubCoordsNV = gl.PathSubCoordsNV

	carbon.PauseTransformFeedback = gl.PauseTransformFeedback
	carbon.PixelMapfv = gl.PixelMapfv
	carbon.PixelMapuiv = gl.PixelMapuiv
	carbon.PixelMapusv = gl.PixelMapusv
	carbon.PixelStoref = gl.PixelStoref

	carbon.PixelStorei = gl.PixelStorei
	carbon.PixelTransferf = gl.PixelTransferf
	carbon.PixelTransferi = gl.PixelTransferi

	carbon.PixelZoom = gl.PixelZoom
	carbon.PointAlongPathNV = gl.PointAlongPathNV
	carbon.PointParameterf = gl.PointParameterf
	carbon.PointParameterfv = gl.PointParameterfv
	carbon.PointParameteri = gl.PointParameteri
	carbon.PointParameteriv = gl.PointParameteriv

	carbon.PointSize = gl.PointSize

	carbon.PolygonMode = gl.PolygonMode

	carbon.PolygonOffset = gl.PolygonOffset
	carbon.PolygonOffsetClamp = gl.PolygonOffsetClamp
	carbon.PolygonOffsetClampEXT = gl.PolygonOffsetClampEXT

	carbon.PolygonStipple = gl.PolygonStipple
	carbon.PopAttrib = gl.PopAttrib
	carbon.PopClientAttrib = gl.PopClientAttrib

	carbon.PopDebugGroup = gl.PopDebugGroup
	carbon.PopDebugGroupKHR = gl.PopDebugGroupKHR
	carbon.PopGroupMarkerEXT = gl.PopGroupMarkerEXT
	carbon.PopMatrix = gl.PopMatrix
	carbon.PopName = gl.PopName
	carbon.PrimitiveBoundingBoxARB = gl.PrimitiveBoundingBoxARB

	carbon.PrimitiveRestartIndex = gl.PrimitiveRestartIndex

	carbon.PrioritizeTextures = gl.PrioritizeTextures

	carbon.ProgramBinary = gl.ProgramBinary

	carbon.ProgramParameteri = gl.ProgramParameteri
	carbon.ProgramParameteriARB = gl.ProgramParameteriARB
	carbon.ProgramParameteriEXT = gl.ProgramParameteriEXT
	carbon.ProgramPathFragmentInputGenNV = gl.ProgramPathFragmentInputGenNV
	carbon.ProgramUniform1d = gl.ProgramUniform1d
	carbon.ProgramUniform1dEXT = gl.ProgramUniform1dEXT
	carbon.ProgramUniform1dv = gl.ProgramUniform1dv
	carbon.ProgramUniform1dvEXT = gl.ProgramUniform1dvEXT

	carbon.ProgramUniform1f = gl.ProgramUniform1f
	carbon.ProgramUniform1fEXT = gl.ProgramUniform1fEXT

	carbon.ProgramUniform1fv = gl.ProgramUniform1fv
	carbon.ProgramUniform1fvEXT = gl.ProgramUniform1fvEXT

	carbon.ProgramUniform1i = gl.ProgramUniform1i
	carbon.ProgramUniform1i64ARB = gl.ProgramUniform1i64ARB
	carbon.ProgramUniform1i64NV = gl.ProgramUniform1i64NV
	carbon.ProgramUniform1i64vARB = gl.ProgramUniform1i64vARB
	carbon.ProgramUniform1i64vNV = gl.ProgramUniform1i64vNV
	carbon.ProgramUniform1iEXT = gl.ProgramUniform1iEXT

	carbon.ProgramUniform1iv = gl.ProgramUniform1iv
	carbon.ProgramUniform1ivEXT = gl.ProgramUniform1ivEXT

	carbon.ProgramUniform1ui = gl.ProgramUniform1ui
	carbon.ProgramUniform1ui64ARB = gl.ProgramUniform1ui64ARB
	carbon.ProgramUniform1ui64NV = gl.ProgramUniform1ui64NV
	carbon.ProgramUniform1ui64vARB = gl.ProgramUniform1ui64vARB
	carbon.ProgramUniform1ui64vNV = gl.ProgramUniform1ui64vNV
	carbon.ProgramUniform1uiEXT = gl.ProgramUniform1uiEXT

	carbon.ProgramUniform1uiv = gl.ProgramUniform1uiv
	carbon.ProgramUniform1uivEXT = gl.ProgramUniform1uivEXT
	carbon.ProgramUniform2d = gl.ProgramUniform2d
	carbon.ProgramUniform2dEXT = gl.ProgramUniform2dEXT
	carbon.ProgramUniform2dv = gl.ProgramUniform2dv
	carbon.ProgramUniform2dvEXT = gl.ProgramUniform2dvEXT

	carbon.ProgramUniform2f = gl.ProgramUniform2f
	carbon.ProgramUniform2fEXT = gl.ProgramUniform2fEXT

	carbon.ProgramUniform2fv = gl.ProgramUniform2fv
	carbon.ProgramUniform2fvEXT = gl.ProgramUniform2fvEXT

	carbon.ProgramUniform2i = gl.ProgramUniform2i
	carbon.ProgramUniform2i64ARB = gl.ProgramUniform2i64ARB
	carbon.ProgramUniform2i64NV = gl.ProgramUniform2i64NV
	carbon.ProgramUniform2i64vARB = gl.ProgramUniform2i64vARB
	carbon.ProgramUniform2i64vNV = gl.ProgramUniform2i64vNV
	carbon.ProgramUniform2iEXT = gl.ProgramUniform2iEXT

	carbon.ProgramUniform2iv = gl.ProgramUniform2iv
	carbon.ProgramUniform2ivEXT = gl.ProgramUniform2ivEXT

	carbon.ProgramUniform2ui = gl.ProgramUniform2ui
	carbon.ProgramUniform2ui64ARB = gl.ProgramUniform2ui64ARB
	carbon.ProgramUniform2ui64NV = gl.ProgramUniform2ui64NV
	carbon.ProgramUniform2ui64vARB = gl.ProgramUniform2ui64vARB
	carbon.ProgramUniform2ui64vNV = gl.ProgramUniform2ui64vNV
	carbon.ProgramUniform2uiEXT = gl.ProgramUniform2uiEXT

	carbon.ProgramUniform2uiv = gl.ProgramUniform2uiv
	carbon.ProgramUniform2uivEXT = gl.ProgramUniform2uivEXT
	carbon.ProgramUniform3d = gl.ProgramUniform3d
	carbon.ProgramUniform3dEXT = gl.ProgramUniform3dEXT
	carbon.ProgramUniform3dv = gl.ProgramUniform3dv
	carbon.ProgramUniform3dvEXT = gl.ProgramUniform3dvEXT

	carbon.ProgramUniform3f = gl.ProgramUniform3f
	carbon.ProgramUniform3fEXT = gl.ProgramUniform3fEXT

	carbon.ProgramUniform3fv = gl.ProgramUniform3fv
	carbon.ProgramUniform3fvEXT = gl.ProgramUniform3fvEXT

	carbon.ProgramUniform3i = gl.ProgramUniform3i
	carbon.ProgramUniform3i64ARB = gl.ProgramUniform3i64ARB
	carbon.ProgramUniform3i64NV = gl.ProgramUniform3i64NV
	carbon.ProgramUniform3i64vARB = gl.ProgramUniform3i64vARB
	carbon.ProgramUniform3i64vNV = gl.ProgramUniform3i64vNV
	carbon.ProgramUniform3iEXT = gl.ProgramUniform3iEXT

	carbon.ProgramUniform3iv = gl.ProgramUniform3iv
	carbon.ProgramUniform3ivEXT = gl.ProgramUniform3ivEXT

	carbon.ProgramUniform3ui = gl.ProgramUniform3ui
	carbon.ProgramUniform3ui64ARB = gl.ProgramUniform3ui64ARB
	carbon.ProgramUniform3ui64NV = gl.ProgramUniform3ui64NV
	carbon.ProgramUniform3ui64vARB = gl.ProgramUniform3ui64vARB
	carbon.ProgramUniform3ui64vNV = gl.ProgramUniform3ui64vNV
	carbon.ProgramUniform3uiEXT = gl.ProgramUniform3uiEXT

	carbon.ProgramUniform3uiv = gl.ProgramUniform3uiv
	carbon.ProgramUniform3uivEXT = gl.ProgramUniform3uivEXT
	carbon.ProgramUniform4d = gl.ProgramUniform4d
	carbon.ProgramUniform4dEXT = gl.ProgramUniform4dEXT
	carbon.ProgramUniform4dv = gl.ProgramUniform4dv
	carbon.ProgramUniform4dvEXT = gl.ProgramUniform4dvEXT

	carbon.ProgramUniform4f = gl.ProgramUniform4f
	carbon.ProgramUniform4fEXT = gl.ProgramUniform4fEXT

	carbon.ProgramUniform4fv = gl.ProgramUniform4fv
	carbon.ProgramUniform4fvEXT = gl.ProgramUniform4fvEXT

	carbon.ProgramUniform4i = gl.ProgramUniform4i
	carbon.ProgramUniform4i64ARB = gl.ProgramUniform4i64ARB
	carbon.ProgramUniform4i64NV = gl.ProgramUniform4i64NV
	carbon.ProgramUniform4i64vARB = gl.ProgramUniform4i64vARB
	carbon.ProgramUniform4i64vNV = gl.ProgramUniform4i64vNV
	carbon.ProgramUniform4iEXT = gl.ProgramUniform4iEXT

	carbon.ProgramUniform4iv = gl.ProgramUniform4iv
	carbon.ProgramUniform4ivEXT = gl.ProgramUniform4ivEXT

	carbon.ProgramUniform4ui = gl.ProgramUniform4ui
	carbon.ProgramUniform4ui64ARB = gl.ProgramUniform4ui64ARB
	carbon.ProgramUniform4ui64NV = gl.ProgramUniform4ui64NV
	carbon.ProgramUniform4ui64vARB = gl.ProgramUniform4ui64vARB
	carbon.ProgramUniform4ui64vNV = gl.ProgramUniform4ui64vNV
	carbon.ProgramUniform4uiEXT = gl.ProgramUniform4uiEXT

	carbon.ProgramUniform4uiv = gl.ProgramUniform4uiv
	carbon.ProgramUniform4uivEXT = gl.ProgramUniform4uivEXT
	carbon.ProgramUniformHandleui64ARB = gl.ProgramUniformHandleui64ARB
	carbon.ProgramUniformHandleui64NV = gl.ProgramUniformHandleui64NV
	carbon.ProgramUniformHandleui64vARB = gl.ProgramUniformHandleui64vARB
	carbon.ProgramUniformHandleui64vNV = gl.ProgramUniformHandleui64vNV
	carbon.ProgramUniformMatrix2dv = gl.ProgramUniformMatrix2dv
	carbon.ProgramUniformMatrix2dvEXT = gl.ProgramUniformMatrix2dvEXT

	carbon.ProgramUniformMatrix2fv = gl.ProgramUniformMatrix2fv
	carbon.ProgramUniformMatrix2fvEXT = gl.ProgramUniformMatrix2fvEXT
	carbon.ProgramUniformMatrix2x3dv = gl.ProgramUniformMatrix2x3dv
	carbon.ProgramUniformMatrix2x3dvEXT = gl.ProgramUniformMatrix2x3dvEXT

	carbon.ProgramUniformMatrix2x3fv = gl.ProgramUniformMatrix2x3fv
	carbon.ProgramUniformMatrix2x3fvEXT = gl.ProgramUniformMatrix2x3fvEXT
	carbon.ProgramUniformMatrix2x4dv = gl.ProgramUniformMatrix2x4dv
	carbon.ProgramUniformMatrix2x4dvEXT = gl.ProgramUniformMatrix2x4dvEXT

	carbon.ProgramUniformMatrix2x4fv = gl.ProgramUniformMatrix2x4fv
	carbon.ProgramUniformMatrix2x4fvEXT = gl.ProgramUniformMatrix2x4fvEXT
	carbon.ProgramUniformMatrix3dv = gl.ProgramUniformMatrix3dv
	carbon.ProgramUniformMatrix3dvEXT = gl.ProgramUniformMatrix3dvEXT

	carbon.ProgramUniformMatrix3fv = gl.ProgramUniformMatrix3fv
	carbon.ProgramUniformMatrix3fvEXT = gl.ProgramUniformMatrix3fvEXT
	carbon.ProgramUniformMatrix3x2dv = gl.ProgramUniformMatrix3x2dv
	carbon.ProgramUniformMatrix3x2dvEXT = gl.ProgramUniformMatrix3x2dvEXT

	carbon.ProgramUniformMatrix3x2fv = gl.ProgramUniformMatrix3x2fv
	carbon.ProgramUniformMatrix3x2fvEXT = gl.ProgramUniformMatrix3x2fvEXT
	carbon.ProgramUniformMatrix3x4dv = gl.ProgramUniformMatrix3x4dv
	carbon.ProgramUniformMatrix3x4dvEXT = gl.ProgramUniformMatrix3x4dvEXT

	carbon.ProgramUniformMatrix3x4fv = gl.ProgramUniformMatrix3x4fv
	carbon.ProgramUniformMatrix3x4fvEXT = gl.ProgramUniformMatrix3x4fvEXT
	carbon.ProgramUniformMatrix4dv = gl.ProgramUniformMatrix4dv
	carbon.ProgramUniformMatrix4dvEXT = gl.ProgramUniformMatrix4dvEXT

	carbon.ProgramUniformMatrix4fv = gl.ProgramUniformMatrix4fv
	carbon.ProgramUniformMatrix4fvEXT = gl.ProgramUniformMatrix4fvEXT
	carbon.ProgramUniformMatrix4x2dv = gl.ProgramUniformMatrix4x2dv
	carbon.ProgramUniformMatrix4x2dvEXT = gl.ProgramUniformMatrix4x2dvEXT

	carbon.ProgramUniformMatrix4x2fv = gl.ProgramUniformMatrix4x2fv
	carbon.ProgramUniformMatrix4x2fvEXT = gl.ProgramUniformMatrix4x2fvEXT
	carbon.ProgramUniformMatrix4x3dv = gl.ProgramUniformMatrix4x3dv
	carbon.ProgramUniformMatrix4x3dvEXT = gl.ProgramUniformMatrix4x3dvEXT

	carbon.ProgramUniformMatrix4x3fv = gl.ProgramUniformMatrix4x3fv
	carbon.ProgramUniformMatrix4x3fvEXT = gl.ProgramUniformMatrix4x3fvEXT
	carbon.ProgramUniformui64NV = gl.ProgramUniformui64NV
	carbon.ProgramUniformui64vNV = gl.ProgramUniformui64vNV

	carbon.ProvokingVertex = gl.ProvokingVertex

	carbon.PushAttrib = gl.PushAttrib

	carbon.PushClientAttrib = gl.PushClientAttrib
	carbon.PushClientAttribDefaultEXT = gl.PushClientAttribDefaultEXT

	carbon.PushDebugGroup = gl.PushDebugGroup
	carbon.PushDebugGroupKHR = gl.PushDebugGroupKHR
	carbon.PushGroupMarkerEXT = gl.PushGroupMarkerEXT

	carbon.PushMatrix = gl.PushMatrix

	carbon.PushName = gl.PushName

	carbon.QueryCounter = gl.QueryCounter
	carbon.RasterPos2d = gl.RasterPos2d
	carbon.RasterPos2dv = gl.RasterPos2dv
	carbon.RasterPos2f = gl.RasterPos2f
	carbon.RasterPos2fv = gl.RasterPos2fv
	carbon.RasterPos2i = gl.RasterPos2i
	carbon.RasterPos2iv = gl.RasterPos2iv
	carbon.RasterPos2s = gl.RasterPos2s
	carbon.RasterPos2sv = gl.RasterPos2sv
	carbon.RasterPos3d = gl.RasterPos3d
	carbon.RasterPos3dv = gl.RasterPos3dv
	carbon.RasterPos3f = gl.RasterPos3f
	carbon.RasterPos3fv = gl.RasterPos3fv
	carbon.RasterPos3i = gl.RasterPos3i
	carbon.RasterPos3iv = gl.RasterPos3iv
	carbon.RasterPos3s = gl.RasterPos3s
	carbon.RasterPos3sv = gl.RasterPos3sv
	carbon.RasterPos4d = gl.RasterPos4d
	carbon.RasterPos4dv = gl.RasterPos4dv
	carbon.RasterPos4f = gl.RasterPos4f
	carbon.RasterPos4fv = gl.RasterPos4fv
	carbon.RasterPos4i = gl.RasterPos4i
	carbon.RasterPos4iv = gl.RasterPos4iv
	carbon.RasterPos4s = gl.RasterPos4s
	carbon.RasterPos4sv = gl.RasterPos4sv
	carbon.RasterSamplesEXT = gl.RasterSamplesEXT

	carbon.ReadBuffer = gl.ReadBuffer

	carbon.ReadPixels = gl.ReadPixels

	carbon.ReadnPixels = gl.ReadnPixels
	carbon.ReadnPixelsARB = gl.ReadnPixelsARB
	carbon.ReadnPixelsKHR = gl.ReadnPixelsKHR
	carbon.Rectd = gl.Rectd
	carbon.Rectdv = gl.Rectdv
	carbon.Rectf = gl.Rectf
	carbon.Rectfv = gl.Rectfv
	carbon.Recti = gl.Recti
	carbon.Rectiv = gl.Rectiv
	carbon.Rects = gl.Rects
	carbon.Rectsv = gl.Rectsv

	carbon.ReleaseShaderCompiler = gl.ReleaseShaderCompiler

	carbon.RenderMode = gl.RenderMode

	carbon.RenderbufferStorage = gl.RenderbufferStorage

	carbon.RenderbufferStorageMultisample = gl.RenderbufferStorageMultisample
	carbon.RenderbufferStorageMultisampleCoverageNV = gl.RenderbufferStorageMultisampleCoverageNV
	carbon.ResolveDepthValuesNV = gl.ResolveDepthValuesNV

	carbon.ResumeTransformFeedback = gl.ResumeTransformFeedback
	carbon.Rotated = gl.Rotated
	carbon.Rotatef = gl.Rotatef

	carbon.SampleCoverage = gl.SampleCoverage

	carbon.SampleMaski = gl.SampleMaski
	carbon.SamplerParameterIiv = gl.SamplerParameterIiv
	carbon.SamplerParameterIuiv = gl.SamplerParameterIuiv
	carbon.SamplerParameterf = gl.SamplerParameterf
	carbon.SamplerParameterfv = gl.SamplerParameterfv
	carbon.SamplerParameteri = gl.SamplerParameteri
	carbon.SamplerParameteriv = gl.SamplerParameteriv
	carbon.Scaled = gl.Scaled
	carbon.Scalef = gl.Scalef

	carbon.Scissor = gl.Scissor
	carbon.ScissorArrayv = gl.ScissorArrayv

	carbon.ScissorIndexed = gl.ScissorIndexed
	carbon.ScissorIndexedv = gl.ScissorIndexedv
	carbon.SecondaryColor3b = gl.SecondaryColor3b
	carbon.SecondaryColor3bv = gl.SecondaryColor3bv
	carbon.SecondaryColor3d = gl.SecondaryColor3d
	carbon.SecondaryColor3dv = gl.SecondaryColor3dv
	carbon.SecondaryColor3f = gl.SecondaryColor3f
	carbon.SecondaryColor3fv = gl.SecondaryColor3fv
	carbon.SecondaryColor3i = gl.SecondaryColor3i
	carbon.SecondaryColor3iv = gl.SecondaryColor3iv
	carbon.SecondaryColor3s = gl.SecondaryColor3s
	carbon.SecondaryColor3sv = gl.SecondaryColor3sv
	carbon.SecondaryColor3ub = gl.SecondaryColor3ub
	carbon.SecondaryColor3ubv = gl.SecondaryColor3ubv
	carbon.SecondaryColor3ui = gl.SecondaryColor3ui
	carbon.SecondaryColor3uiv = gl.SecondaryColor3uiv
	carbon.SecondaryColor3us = gl.SecondaryColor3us
	carbon.SecondaryColor3usv = gl.SecondaryColor3usv
	carbon.SecondaryColorFormatNV = gl.SecondaryColorFormatNV

	carbon.SecondaryColorPointer = gl.SecondaryColorPointer

	carbon.SelectBuffer = gl.SelectBuffer
	carbon.SelectPerfMonitorCountersAMD = gl.SelectPerfMonitorCountersAMD

	carbon.ShadeModel = gl.ShadeModel

	carbon.ShaderBinary = gl.ShaderBinary

	carbon.ShaderSource = gl.ShaderSource

	carbon.ShaderStorageBlockBinding = gl.ShaderStorageBlockBinding
	carbon.SignalVkFenceNV = gl.SignalVkFenceNV
	carbon.SignalVkSemaphoreNV = gl.SignalVkSemaphoreNV
	carbon.SpecializeShader = gl.SpecializeShader
	carbon.SpecializeShaderARB = gl.SpecializeShaderARB
	carbon.StateCaptureNV = gl.StateCaptureNV
	carbon.StencilFillPathInstancedNV = gl.StencilFillPathInstancedNV
	carbon.StencilFillPathNV = gl.StencilFillPathNV

	carbon.StencilFunc = gl.StencilFunc

	carbon.StencilFuncSeparate = gl.StencilFuncSeparate

	carbon.StencilMask = gl.StencilMask

	carbon.StencilMaskSeparate = gl.StencilMaskSeparate

	carbon.StencilOp = gl.StencilOp

	carbon.StencilOpSeparate = gl.StencilOpSeparate
	carbon.StencilStrokePathInstancedNV = gl.StencilStrokePathInstancedNV
	carbon.StencilStrokePathNV = gl.StencilStrokePathNV
	carbon.StencilThenCoverFillPathInstancedNV = gl.StencilThenCoverFillPathInstancedNV
	carbon.StencilThenCoverFillPathNV = gl.StencilThenCoverFillPathNV
	carbon.StencilThenCoverStrokePathInstancedNV = gl.StencilThenCoverStrokePathInstancedNV
	carbon.StencilThenCoverStrokePathNV = gl.StencilThenCoverStrokePathNV
	carbon.SubpixelPrecisionBiasNV = gl.SubpixelPrecisionBiasNV

	carbon.TexBuffer = gl.TexBuffer
	carbon.TexBufferARB = gl.TexBufferARB

	carbon.TexBufferRange = gl.TexBufferRange
	carbon.TexCoord1d = gl.TexCoord1d
	carbon.TexCoord1dv = gl.TexCoord1dv
	carbon.TexCoord1f = gl.TexCoord1f
	carbon.TexCoord1fv = gl.TexCoord1fv
	carbon.TexCoord1i = gl.TexCoord1i
	carbon.TexCoord1iv = gl.TexCoord1iv
	carbon.TexCoord1s = gl.TexCoord1s
	carbon.TexCoord1sv = gl.TexCoord1sv
	carbon.TexCoord2d = gl.TexCoord2d
	carbon.TexCoord2dv = gl.TexCoord2dv
	carbon.TexCoord2f = gl.TexCoord2f
	carbon.TexCoord2fv = gl.TexCoord2fv
	carbon.TexCoord2i = gl.TexCoord2i
	carbon.TexCoord2iv = gl.TexCoord2iv
	carbon.TexCoord2s = gl.TexCoord2s
	carbon.TexCoord2sv = gl.TexCoord2sv
	carbon.TexCoord3d = gl.TexCoord3d
	carbon.TexCoord3dv = gl.TexCoord3dv
	carbon.TexCoord3f = gl.TexCoord3f
	carbon.TexCoord3fv = gl.TexCoord3fv
	carbon.TexCoord3i = gl.TexCoord3i
	carbon.TexCoord3iv = gl.TexCoord3iv
	carbon.TexCoord3s = gl.TexCoord3s
	carbon.TexCoord3sv = gl.TexCoord3sv
	carbon.TexCoord4d = gl.TexCoord4d
	carbon.TexCoord4dv = gl.TexCoord4dv
	carbon.TexCoord4f = gl.TexCoord4f
	carbon.TexCoord4fv = gl.TexCoord4fv
	carbon.TexCoord4i = gl.TexCoord4i
	carbon.TexCoord4iv = gl.TexCoord4iv
	carbon.TexCoord4s = gl.TexCoord4s
	carbon.TexCoord4sv = gl.TexCoord4sv
	carbon.TexCoordFormatNV = gl.TexCoordFormatNV

	carbon.TexCoordPointer = gl.TexCoordPointer
	carbon.TexEnvf = gl.TexEnvf
	carbon.TexEnvfv = gl.TexEnvfv
	carbon.TexEnvi = gl.TexEnvi
	carbon.TexEnviv = gl.TexEnviv
	carbon.TexGend = gl.TexGend
	carbon.TexGendv = gl.TexGendv
	carbon.TexGenf = gl.TexGenf
	carbon.TexGenfv = gl.TexGenfv
	carbon.TexGeni = gl.TexGeni
	carbon.TexGeniv = gl.TexGeniv

	carbon.TexImage1D = gl.TexImage1D

	carbon.TexImage2D = gl.TexImage2D

	carbon.TexImage2DMultisample = gl.TexImage2DMultisample

	carbon.TexImage3D = gl.TexImage3D

	carbon.TexImage3DMultisample = gl.TexImage3DMultisample
	carbon.TexPageCommitmentARB = gl.TexPageCommitmentARB
	carbon.TexParameterIiv = gl.TexParameterIiv
	carbon.TexParameterIuiv = gl.TexParameterIuiv
	carbon.TexParameterf = gl.TexParameterf
	carbon.TexParameterfv = gl.TexParameterfv
	carbon.TexParameteri = gl.TexParameteri
	carbon.TexParameteriv = gl.TexParameteriv

	carbon.TexStorage1D = gl.TexStorage1D

	carbon.TexStorage2D = gl.TexStorage2D

	carbon.TexStorage2DMultisample = gl.TexStorage2DMultisample

	carbon.TexStorage3D = gl.TexStorage3D

	carbon.TexStorage3DMultisample = gl.TexStorage3DMultisample

	carbon.TexSubImage1D = gl.TexSubImage1D

	carbon.TexSubImage2D = gl.TexSubImage2D

	carbon.TexSubImage3D = gl.TexSubImage3D

	carbon.TextureBarrier = gl.TextureBarrier
	carbon.TextureBarrierNV = gl.TextureBarrierNV

	carbon.TextureBuffer = gl.TextureBuffer
	carbon.TextureBufferEXT = gl.TextureBufferEXT

	carbon.TextureBufferRange = gl.TextureBufferRange
	carbon.TextureBufferRangeEXT = gl.TextureBufferRangeEXT
	carbon.TextureImage1DEXT = gl.TextureImage1DEXT
	carbon.TextureImage2DEXT = gl.TextureImage2DEXT
	carbon.TextureImage3DEXT = gl.TextureImage3DEXT
	carbon.TexturePageCommitmentEXT = gl.TexturePageCommitmentEXT
	carbon.TextureParameterIiv = gl.TextureParameterIiv
	carbon.TextureParameterIivEXT = gl.TextureParameterIivEXT
	carbon.TextureParameterIuiv = gl.TextureParameterIuiv
	carbon.TextureParameterIuivEXT = gl.TextureParameterIuivEXT
	carbon.TextureParameterf = gl.TextureParameterf
	carbon.TextureParameterfEXT = gl.TextureParameterfEXT
	carbon.TextureParameterfv = gl.TextureParameterfv
	carbon.TextureParameterfvEXT = gl.TextureParameterfvEXT
	carbon.TextureParameteri = gl.TextureParameteri
	carbon.TextureParameteriEXT = gl.TextureParameteriEXT
	carbon.TextureParameteriv = gl.TextureParameteriv
	carbon.TextureParameterivEXT = gl.TextureParameterivEXT
	carbon.TextureRenderbufferEXT = gl.TextureRenderbufferEXT

	carbon.TextureStorage1D = gl.TextureStorage1D
	carbon.TextureStorage1DEXT = gl.TextureStorage1DEXT

	carbon.TextureStorage2D = gl.TextureStorage2D
	carbon.TextureStorage2DEXT = gl.TextureStorage2DEXT

	carbon.TextureStorage2DMultisample = gl.TextureStorage2DMultisample
	carbon.TextureStorage2DMultisampleEXT = gl.TextureStorage2DMultisampleEXT

	carbon.TextureStorage3D = gl.TextureStorage3D
	carbon.TextureStorage3DEXT = gl.TextureStorage3DEXT

	carbon.TextureStorage3DMultisample = gl.TextureStorage3DMultisample
	carbon.TextureStorage3DMultisampleEXT = gl.TextureStorage3DMultisampleEXT

	carbon.TextureSubImage1D = gl.TextureSubImage1D
	carbon.TextureSubImage1DEXT = gl.TextureSubImage1DEXT

	carbon.TextureSubImage2D = gl.TextureSubImage2D
	carbon.TextureSubImage2DEXT = gl.TextureSubImage2DEXT

	carbon.TextureSubImage3D = gl.TextureSubImage3D
	carbon.TextureSubImage3DEXT = gl.TextureSubImage3DEXT

	carbon.TextureView = gl.TextureView

	carbon.TransformFeedbackBufferBase = gl.TransformFeedbackBufferBase

	carbon.TransformFeedbackBufferRange = gl.TransformFeedbackBufferRange

	carbon.TransformFeedbackVaryings = gl.TransformFeedbackVaryings
	carbon.TransformPathNV = gl.TransformPathNV
	carbon.Translated = gl.Translated
	carbon.Translatef = gl.Translatef
	carbon.Uniform1d = gl.Uniform1d
	carbon.Uniform1dv = gl.Uniform1dv

	carbon.Uniform1f = gl.Uniform1f

	carbon.Uniform1fv = gl.Uniform1fv

	carbon.Uniform1i = gl.Uniform1i
	carbon.Uniform1i64ARB = gl.Uniform1i64ARB
	carbon.Uniform1i64NV = gl.Uniform1i64NV
	carbon.Uniform1i64vARB = gl.Uniform1i64vARB
	carbon.Uniform1i64vNV = gl.Uniform1i64vNV

	carbon.Uniform1iv = gl.Uniform1iv

	carbon.Uniform1ui = gl.Uniform1ui
	carbon.Uniform1ui64ARB = gl.Uniform1ui64ARB
	carbon.Uniform1ui64NV = gl.Uniform1ui64NV
	carbon.Uniform1ui64vARB = gl.Uniform1ui64vARB
	carbon.Uniform1ui64vNV = gl.Uniform1ui64vNV

	carbon.Uniform1uiv = gl.Uniform1uiv
	carbon.Uniform2d = gl.Uniform2d
	carbon.Uniform2dv = gl.Uniform2dv

	carbon.Uniform2f = gl.Uniform2f

	carbon.Uniform2fv = gl.Uniform2fv

	carbon.Uniform2i = gl.Uniform2i
	carbon.Uniform2i64ARB = gl.Uniform2i64ARB
	carbon.Uniform2i64NV = gl.Uniform2i64NV
	carbon.Uniform2i64vARB = gl.Uniform2i64vARB
	carbon.Uniform2i64vNV = gl.Uniform2i64vNV

	carbon.Uniform2iv = gl.Uniform2iv

	carbon.Uniform2ui = gl.Uniform2ui
	carbon.Uniform2ui64ARB = gl.Uniform2ui64ARB
	carbon.Uniform2ui64NV = gl.Uniform2ui64NV
	carbon.Uniform2ui64vARB = gl.Uniform2ui64vARB
	carbon.Uniform2ui64vNV = gl.Uniform2ui64vNV

	carbon.Uniform2uiv = gl.Uniform2uiv
	carbon.Uniform3d = gl.Uniform3d
	carbon.Uniform3dv = gl.Uniform3dv

	carbon.Uniform3f = gl.Uniform3f

	carbon.Uniform3fv = gl.Uniform3fv

	carbon.Uniform3i = gl.Uniform3i
	carbon.Uniform3i64ARB = gl.Uniform3i64ARB
	carbon.Uniform3i64NV = gl.Uniform3i64NV
	carbon.Uniform3i64vARB = gl.Uniform3i64vARB
	carbon.Uniform3i64vNV = gl.Uniform3i64vNV

	carbon.Uniform3iv = gl.Uniform3iv

	carbon.Uniform3ui = gl.Uniform3ui
	carbon.Uniform3ui64ARB = gl.Uniform3ui64ARB
	carbon.Uniform3ui64NV = gl.Uniform3ui64NV
	carbon.Uniform3ui64vARB = gl.Uniform3ui64vARB
	carbon.Uniform3ui64vNV = gl.Uniform3ui64vNV

	carbon.Uniform3uiv = gl.Uniform3uiv
	carbon.Uniform4d = gl.Uniform4d
	carbon.Uniform4dv = gl.Uniform4dv

	carbon.Uniform4f = gl.Uniform4f

	carbon.Uniform4fv = gl.Uniform4fv

	carbon.Uniform4i = gl.Uniform4i
	carbon.Uniform4i64ARB = gl.Uniform4i64ARB
	carbon.Uniform4i64NV = gl.Uniform4i64NV
	carbon.Uniform4i64vARB = gl.Uniform4i64vARB
	carbon.Uniform4i64vNV = gl.Uniform4i64vNV

	carbon.Uniform4iv = gl.Uniform4iv

	carbon.Uniform4ui = gl.Uniform4ui
	carbon.Uniform4ui64ARB = gl.Uniform4ui64ARB
	carbon.Uniform4ui64NV = gl.Uniform4ui64NV
	carbon.Uniform4ui64vARB = gl.Uniform4ui64vARB
	carbon.Uniform4ui64vNV = gl.Uniform4ui64vNV

	carbon.Uniform4uiv = gl.Uniform4uiv

	carbon.UniformBlockBinding = gl.UniformBlockBinding
	carbon.UniformHandleui64ARB = gl.UniformHandleui64ARB
	carbon.UniformHandleui64NV = gl.UniformHandleui64NV
	carbon.UniformHandleui64vARB = gl.UniformHandleui64vARB
	carbon.UniformHandleui64vNV = gl.UniformHandleui64vNV
	carbon.UniformMatrix2dv = gl.UniformMatrix2dv

	carbon.UniformMatrix2fv = gl.UniformMatrix2fv
	carbon.UniformMatrix2x3dv = gl.UniformMatrix2x3dv

	carbon.UniformMatrix2x3fv = gl.UniformMatrix2x3fv
	carbon.UniformMatrix2x4dv = gl.UniformMatrix2x4dv

	carbon.UniformMatrix2x4fv = gl.UniformMatrix2x4fv
	carbon.UniformMatrix3dv = gl.UniformMatrix3dv

	carbon.UniformMatrix3fv = gl.UniformMatrix3fv
	carbon.UniformMatrix3x2dv = gl.UniformMatrix3x2dv

	carbon.UniformMatrix3x2fv = gl.UniformMatrix3x2fv
	carbon.UniformMatrix3x4dv = gl.UniformMatrix3x4dv

	carbon.UniformMatrix3x4fv = gl.UniformMatrix3x4fv
	carbon.UniformMatrix4dv = gl.UniformMatrix4dv

	carbon.UniformMatrix4fv = gl.UniformMatrix4fv
	carbon.UniformMatrix4x2dv = gl.UniformMatrix4x2dv

	carbon.UniformMatrix4x2fv = gl.UniformMatrix4x2fv
	carbon.UniformMatrix4x3dv = gl.UniformMatrix4x3dv

	carbon.UniformMatrix4x3fv = gl.UniformMatrix4x3fv
	carbon.UniformSubroutinesuiv = gl.UniformSubroutinesuiv
	carbon.Uniformui64NV = gl.Uniformui64NV
	carbon.Uniformui64vNV = gl.Uniformui64vNV

	carbon.UnmapBuffer = gl.UnmapBuffer

	carbon.UnmapNamedBuffer = gl.UnmapNamedBuffer
	carbon.UnmapNamedBufferEXT = gl.UnmapNamedBufferEXT

	carbon.UseProgram = gl.UseProgram

	carbon.UseProgramStages = gl.UseProgramStages
	carbon.UseProgramStagesEXT = gl.UseProgramStagesEXT
	carbon.UseShaderProgramEXT = gl.UseShaderProgramEXT

	carbon.ValidateProgram = gl.ValidateProgram

	carbon.ValidateProgramPipeline = gl.ValidateProgramPipeline
	carbon.ValidateProgramPipelineEXT = gl.ValidateProgramPipelineEXT
	carbon.Vertex2d = gl.Vertex2d
	carbon.Vertex2dv = gl.Vertex2dv
	carbon.Vertex2f = gl.Vertex2f
	carbon.Vertex2fv = gl.Vertex2fv
	carbon.Vertex2i = gl.Vertex2i
	carbon.Vertex2iv = gl.Vertex2iv
	carbon.Vertex2s = gl.Vertex2s
	carbon.Vertex2sv = gl.Vertex2sv
	carbon.Vertex3d = gl.Vertex3d
	carbon.Vertex3dv = gl.Vertex3dv
	carbon.Vertex3f = gl.Vertex3f
	carbon.Vertex3fv = gl.Vertex3fv
	carbon.Vertex3i = gl.Vertex3i
	carbon.Vertex3iv = gl.Vertex3iv
	carbon.Vertex3s = gl.Vertex3s
	carbon.Vertex3sv = gl.Vertex3sv
	carbon.Vertex4d = gl.Vertex4d
	carbon.Vertex4dv = gl.Vertex4dv
	carbon.Vertex4f = gl.Vertex4f
	carbon.Vertex4fv = gl.Vertex4fv
	carbon.Vertex4i = gl.Vertex4i
	carbon.Vertex4iv = gl.Vertex4iv
	carbon.Vertex4s = gl.Vertex4s
	carbon.Vertex4sv = gl.Vertex4sv
	carbon.VertexArrayAttribBinding = gl.VertexArrayAttribBinding

	carbon.VertexArrayAttribFormat = gl.VertexArrayAttribFormat
	carbon.VertexArrayAttribIFormat = gl.VertexArrayAttribIFormat
	carbon.VertexArrayAttribLFormat = gl.VertexArrayAttribLFormat
	carbon.VertexArrayBindVertexBufferEXT = gl.VertexArrayBindVertexBufferEXT

	carbon.VertexArrayBindingDivisor = gl.VertexArrayBindingDivisor
	carbon.VertexArrayColorOffsetEXT = gl.VertexArrayColorOffsetEXT
	carbon.VertexArrayEdgeFlagOffsetEXT = gl.VertexArrayEdgeFlagOffsetEXT

	carbon.VertexArrayElementBuffer = gl.VertexArrayElementBuffer
	carbon.VertexArrayFogCoordOffsetEXT = gl.VertexArrayFogCoordOffsetEXT
	carbon.VertexArrayIndexOffsetEXT = gl.VertexArrayIndexOffsetEXT
	carbon.VertexArrayMultiTexCoordOffsetEXT = gl.VertexArrayMultiTexCoordOffsetEXT
	carbon.VertexArrayNormalOffsetEXT = gl.VertexArrayNormalOffsetEXT
	carbon.VertexArraySecondaryColorOffsetEXT = gl.VertexArraySecondaryColorOffsetEXT
	carbon.VertexArrayTexCoordOffsetEXT = gl.VertexArrayTexCoordOffsetEXT
	carbon.VertexArrayVertexAttribBindingEXT = gl.VertexArrayVertexAttribBindingEXT
	carbon.VertexArrayVertexAttribDivisorEXT = gl.VertexArrayVertexAttribDivisorEXT
	carbon.VertexArrayVertexAttribFormatEXT = gl.VertexArrayVertexAttribFormatEXT
	carbon.VertexArrayVertexAttribIFormatEXT = gl.VertexArrayVertexAttribIFormatEXT
	carbon.VertexArrayVertexAttribIOffsetEXT = gl.VertexArrayVertexAttribIOffsetEXT
	carbon.VertexArrayVertexAttribLFormatEXT = gl.VertexArrayVertexAttribLFormatEXT
	carbon.VertexArrayVertexAttribLOffsetEXT = gl.VertexArrayVertexAttribLOffsetEXT
	carbon.VertexArrayVertexAttribOffsetEXT = gl.VertexArrayVertexAttribOffsetEXT
	carbon.VertexArrayVertexBindingDivisorEXT = gl.VertexArrayVertexBindingDivisorEXT

	carbon.VertexArrayVertexBuffer = gl.VertexArrayVertexBuffer

	carbon.VertexArrayVertexBuffers = gl.VertexArrayVertexBuffers
	carbon.VertexArrayVertexOffsetEXT = gl.VertexArrayVertexOffsetEXT
	carbon.VertexAttrib1d = gl.VertexAttrib1d
	carbon.VertexAttrib1dv = gl.VertexAttrib1dv
	carbon.VertexAttrib1f = gl.VertexAttrib1f
	carbon.VertexAttrib1fv = gl.VertexAttrib1fv
	carbon.VertexAttrib1s = gl.VertexAttrib1s
	carbon.VertexAttrib1sv = gl.VertexAttrib1sv
	carbon.VertexAttrib2d = gl.VertexAttrib2d
	carbon.VertexAttrib2dv = gl.VertexAttrib2dv
	carbon.VertexAttrib2f = gl.VertexAttrib2f
	carbon.VertexAttrib2fv = gl.VertexAttrib2fv
	carbon.VertexAttrib2s = gl.VertexAttrib2s
	carbon.VertexAttrib2sv = gl.VertexAttrib2sv
	carbon.VertexAttrib3d = gl.VertexAttrib3d
	carbon.VertexAttrib3dv = gl.VertexAttrib3dv
	carbon.VertexAttrib3f = gl.VertexAttrib3f
	carbon.VertexAttrib3fv = gl.VertexAttrib3fv
	carbon.VertexAttrib3s = gl.VertexAttrib3s
	carbon.VertexAttrib3sv = gl.VertexAttrib3sv
	carbon.VertexAttrib4Nbv = gl.VertexAttrib4Nbv
	carbon.VertexAttrib4Niv = gl.VertexAttrib4Niv
	carbon.VertexAttrib4Nsv = gl.VertexAttrib4Nsv
	carbon.VertexAttrib4Nub = gl.VertexAttrib4Nub
	carbon.VertexAttrib4Nubv = gl.VertexAttrib4Nubv
	carbon.VertexAttrib4Nuiv = gl.VertexAttrib4Nuiv
	carbon.VertexAttrib4Nusv = gl.VertexAttrib4Nusv
	carbon.VertexAttrib4bv = gl.VertexAttrib4bv
	carbon.VertexAttrib4d = gl.VertexAttrib4d
	carbon.VertexAttrib4dv = gl.VertexAttrib4dv
	carbon.VertexAttrib4f = gl.VertexAttrib4f
	carbon.VertexAttrib4fv = gl.VertexAttrib4fv
	carbon.VertexAttrib4iv = gl.VertexAttrib4iv
	carbon.VertexAttrib4s = gl.VertexAttrib4s
	carbon.VertexAttrib4sv = gl.VertexAttrib4sv
	carbon.VertexAttrib4ubv = gl.VertexAttrib4ubv
	carbon.VertexAttrib4uiv = gl.VertexAttrib4uiv
	carbon.VertexAttrib4usv = gl.VertexAttrib4usv

	carbon.VertexAttribBinding = gl.VertexAttribBinding

	carbon.VertexAttribDivisor = gl.VertexAttribDivisor
	carbon.VertexAttribDivisorARB = gl.VertexAttribDivisorARB

	carbon.VertexAttribFormat = gl.VertexAttribFormat
	carbon.VertexAttribFormatNV = gl.VertexAttribFormatNV
	carbon.VertexAttribI1i = gl.VertexAttribI1i
	carbon.VertexAttribI1iv = gl.VertexAttribI1iv
	carbon.VertexAttribI1ui = gl.VertexAttribI1ui
	carbon.VertexAttribI1uiv = gl.VertexAttribI1uiv
	carbon.VertexAttribI2i = gl.VertexAttribI2i
	carbon.VertexAttribI2iv = gl.VertexAttribI2iv
	carbon.VertexAttribI2ui = gl.VertexAttribI2ui
	carbon.VertexAttribI2uiv = gl.VertexAttribI2uiv
	carbon.VertexAttribI3i = gl.VertexAttribI3i
	carbon.VertexAttribI3iv = gl.VertexAttribI3iv
	carbon.VertexAttribI3ui = gl.VertexAttribI3ui
	carbon.VertexAttribI3uiv = gl.VertexAttribI3uiv
	carbon.VertexAttribI4bv = gl.VertexAttribI4bv
	carbon.VertexAttribI4i = gl.VertexAttribI4i
	carbon.VertexAttribI4iv = gl.VertexAttribI4iv
	carbon.VertexAttribI4sv = gl.VertexAttribI4sv
	carbon.VertexAttribI4ubv = gl.VertexAttribI4ubv
	carbon.VertexAttribI4ui = gl.VertexAttribI4ui
	carbon.VertexAttribI4uiv = gl.VertexAttribI4uiv
	carbon.VertexAttribI4usv = gl.VertexAttribI4usv
	carbon.VertexAttribIFormat = gl.VertexAttribIFormat
	carbon.VertexAttribIFormatNV = gl.VertexAttribIFormatNV
	carbon.VertexAttribIPointer = gl.VertexAttribIPointer
	carbon.VertexAttribL1d = gl.VertexAttribL1d
	carbon.VertexAttribL1dv = gl.VertexAttribL1dv
	carbon.VertexAttribL1i64NV = gl.VertexAttribL1i64NV
	carbon.VertexAttribL1i64vNV = gl.VertexAttribL1i64vNV
	carbon.VertexAttribL1ui64ARB = gl.VertexAttribL1ui64ARB
	carbon.VertexAttribL1ui64NV = gl.VertexAttribL1ui64NV
	carbon.VertexAttribL1ui64vARB = gl.VertexAttribL1ui64vARB
	carbon.VertexAttribL1ui64vNV = gl.VertexAttribL1ui64vNV
	carbon.VertexAttribL2d = gl.VertexAttribL2d
	carbon.VertexAttribL2dv = gl.VertexAttribL2dv
	carbon.VertexAttribL2i64NV = gl.VertexAttribL2i64NV
	carbon.VertexAttribL2i64vNV = gl.VertexAttribL2i64vNV
	carbon.VertexAttribL2ui64NV = gl.VertexAttribL2ui64NV
	carbon.VertexAttribL2ui64vNV = gl.VertexAttribL2ui64vNV
	carbon.VertexAttribL3d = gl.VertexAttribL3d
	carbon.VertexAttribL3dv = gl.VertexAttribL3dv
	carbon.VertexAttribL3i64NV = gl.VertexAttribL3i64NV
	carbon.VertexAttribL3i64vNV = gl.VertexAttribL3i64vNV
	carbon.VertexAttribL3ui64NV = gl.VertexAttribL3ui64NV
	carbon.VertexAttribL3ui64vNV = gl.VertexAttribL3ui64vNV
	carbon.VertexAttribL4d = gl.VertexAttribL4d
	carbon.VertexAttribL4dv = gl.VertexAttribL4dv
	carbon.VertexAttribL4i64NV = gl.VertexAttribL4i64NV
	carbon.VertexAttribL4i64vNV = gl.VertexAttribL4i64vNV
	carbon.VertexAttribL4ui64NV = gl.VertexAttribL4ui64NV
	carbon.VertexAttribL4ui64vNV = gl.VertexAttribL4ui64vNV
	carbon.VertexAttribLFormat = gl.VertexAttribLFormat
	carbon.VertexAttribLFormatNV = gl.VertexAttribLFormatNV
	carbon.VertexAttribLPointer = gl.VertexAttribLPointer
	carbon.VertexAttribP1ui = gl.VertexAttribP1ui
	carbon.VertexAttribP1uiv = gl.VertexAttribP1uiv
	carbon.VertexAttribP2ui = gl.VertexAttribP2ui
	carbon.VertexAttribP2uiv = gl.VertexAttribP2uiv
	carbon.VertexAttribP3ui = gl.VertexAttribP3ui
	carbon.VertexAttribP3uiv = gl.VertexAttribP3uiv
	carbon.VertexAttribP4ui = gl.VertexAttribP4ui
	carbon.VertexAttribP4uiv = gl.VertexAttribP4uiv

	carbon.VertexAttribPointer = gl.VertexAttribPointer

	carbon.VertexBindingDivisor = gl.VertexBindingDivisor
	carbon.VertexFormatNV = gl.VertexFormatNV

	carbon.VertexPointer = gl.VertexPointer

	carbon.Viewport = gl.Viewport
	carbon.ViewportArrayv = gl.ViewportArrayv
	carbon.ViewportIndexedf = gl.ViewportIndexedf
	carbon.ViewportIndexedfv = gl.ViewportIndexedfv
	carbon.ViewportPositionWScaleNV = gl.ViewportPositionWScaleNV
	carbon.ViewportSwizzleNV = gl.ViewportSwizzleNV

	carbon.WaitSync = gl.WaitSync
	carbon.WaitVkSemaphoreNV = gl.WaitVkSemaphoreNV
	carbon.WeightPathsNV = gl.WeightPathsNV
	carbon.WindowPos2d = gl.WindowPos2d
	carbon.WindowPos2dv = gl.WindowPos2dv
	carbon.WindowPos2f = gl.WindowPos2f
	carbon.WindowPos2fv = gl.WindowPos2fv
	carbon.WindowPos2i = gl.WindowPos2i
	carbon.WindowPos2iv = gl.WindowPos2iv
	carbon.WindowPos2s = gl.WindowPos2s
	carbon.WindowPos2sv = gl.WindowPos2sv
	carbon.WindowPos3d = gl.WindowPos3d
	carbon.WindowPos3dv = gl.WindowPos3dv
	carbon.WindowPos3f = gl.WindowPos3f
	carbon.WindowPos3fv = gl.WindowPos3fv
	carbon.WindowPos3i = gl.WindowPos3i
	carbon.WindowPos3iv = gl.WindowPos3iv
	carbon.WindowPos3s = gl.WindowPos3s
	carbon.WindowPos3sv = gl.WindowPos3sv
	carbon.WindowRectanglesEXT = gl.WindowRectanglesEXT

	carbon.Init = gl.Init
}
