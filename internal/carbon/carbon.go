package carbon

import "C"
import (
	"unsafe"
)

const (
	GL_2D                                                      = 0x0600
	GL_2_BYTES                                                 = 0x1407
	GL_3D                                                      = 0x0601
	GL_3D_COLOR                                                = 0x0602
	GL_3D_COLOR_TEXTURE                                        = 0x0603
	GL_3_BYTES                                                 = 0x1408
	GL_4D_COLOR_TEXTURE                                        = 0x0604
	GL_4_BYTES                                                 = 0x1409
	ACCUM                                                      = 0x0100
	ACCUM_ADJACENT_PAIRS_NV                                    = 0x90AD
	ACCUM_ALPHA_BITS                                           = 0x0D5B
	ACCUM_BLUE_BITS                                            = 0x0D5A
	ACCUM_BUFFER_BIT                                           = 0x00000200
	ACCUM_CLEAR_VALUE                                          = 0x0B80
	ACCUM_GREEN_BITS                                           = 0x0D59
	ACCUM_RED_BITS                                             = 0x0D58
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	ACTIVE_PROGRAM                                             = 0x8259
	ACTIVE_PROGRAM_EXT                                         = 0x8B8D
	ACTIVE_RESOURCES                                           = 0x92F5
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	ACTIVE_TEXTURE                                             = 0x84E0
	ACTIVE_UNIFORMS                                            = 0x8B86
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	ACTIVE_VARIABLES                                           = 0x9305
	ADD                                                        = 0x0104
	ADD_SIGNED                                                 = 0x8574
	ADJACENT_PAIRS_NV                                          = 0x90AE
	AFFINE_2D_NV                                               = 0x9092
	AFFINE_3D_NV                                               = 0x9094
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	ALIASED_POINT_SIZE_RANGE                                   = 0x846D
	ALL_ATTRIB_BITS                                            = 0xFFFFFFFF
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	ALL_SHADER_BITS_EXT                                        = 0xFFFFFFFF
	ALPHA                                                      = 0x1906
	ALPHA12                                                    = 0x803D
	ALPHA16                                                    = 0x803E
	ALPHA4                                                     = 0x803B
	ALPHA8                                                     = 0x803C
	ALPHA_BIAS                                                 = 0x0D1D
	ALPHA_BITS                                                 = 0x0D55
	ALPHA_INTEGER                                              = 0x8D97
	ALPHA_REF_COMMAND_NV                                       = 0x000F
	ALPHA_SCALE                                                = 0x0D1C
	ALPHA_TEST                                                 = 0x0BC0
	ALPHA_TEST_FUNC                                            = 0x0BC1
	ALPHA_TEST_REF                                             = 0x0BC2
	ALREADY_SIGNALED                                           = 0x911A
	ALWAYS                                                     = 0x0207
	AMBIENT                                                    = 0x1200
	AMBIENT_AND_DIFFUSE                                        = 0x1602
	AND                                                        = 0x1501
	AND_INVERTED                                               = 0x1504
	AND_REVERSE                                                = 0x1502
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	ARC_TO_NV                                                  = 0xFE
	ARRAY_BUFFER                                               = 0x8892
	ARRAY_BUFFER_BINDING                                       = 0x8894
	ARRAY_SIZE                                                 = 0x92FB
	ARRAY_STRIDE                                               = 0x92FE
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	ATTACHED_SHADERS                                           = 0x8B85
	ATTRIBUTE_ADDRESS_COMMAND_NV                               = 0x0009
	ATTRIB_STACK_DEPTH                                         = 0x0BB0
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	AUTO_NORMAL                                                = 0x0D80
	AUX0                                                       = 0x0409
	AUX1                                                       = 0x040A
	AUX2                                                       = 0x040B
	AUX3                                                       = 0x040C
	AUX_BUFFERS                                                = 0x0C00
	BACK                                                       = 0x0405
	BACK_LEFT                                                  = 0x0402
	BACK_RIGHT                                                 = 0x0403
	BEVEL_NV                                                   = 0x90A6
	BGR                                                        = 0x80E0
	BGRA                                                       = 0x80E1
	BGRA_INTEGER                                               = 0x8D9B
	BGR_INTEGER                                                = 0x8D9A
	BITMAP                                                     = 0x1A00
	BITMAP_TOKEN                                               = 0x0704
	BLACKHOLE_RENDER_INTEL                                     = 0x83FC
	BLEND                                                      = 0x0BE2
	BLEND_ADVANCED_COHERENT_KHR                                = 0x9285
	BLEND_ADVANCED_COHERENT_NV                                 = 0x9285
	BLEND_COLOR                                                = 0x8005
	BLEND_COLOR_COMMAND_NV                                     = 0x000B
	BLEND_DST                                                  = 0x0BE0
	BLEND_DST_ALPHA                                            = 0x80CA
	BLEND_DST_RGB                                              = 0x80C8
	BLEND_EQUATION                                             = 0x8009
	BLEND_EQUATION_ALPHA                                       = 0x883D
	BLEND_EQUATION_RGB                                         = 0x8009
	BLEND_OVERLAP_NV                                           = 0x9281
	BLEND_PREMULTIPLIED_SRC_NV                                 = 0x9280
	BLEND_SRC                                                  = 0x0BE1
	BLEND_SRC_ALPHA                                            = 0x80CB
	BLEND_SRC_RGB                                              = 0x80C9
	BLOCK_INDEX                                                = 0x92FD
	BLUE                                                       = 0x1905
	BLUE_BIAS                                                  = 0x0D1B
	BLUE_BITS                                                  = 0x0D54
	BLUE_INTEGER                                               = 0x8D96
	BLUE_NV                                                    = 0x1905
	BLUE_SCALE                                                 = 0x0D1A
	BOLD_BIT_NV                                                = 0x01
	BOOL                                                       = 0x8B56
	BOOL_VEC2                                                  = 0x8B57
	BOOL_VEC3                                                  = 0x8B58
	BOOL_VEC4                                                  = 0x8B59
	BOUNDING_BOX_NV                                            = 0x908D
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                          = 0x909C
	BUFFER                                                     = 0x82E0
	BUFFER_ACCESS                                              = 0x88BB
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	BUFFER_BINDING                                             = 0x9302
	BUFFER_DATA_SIZE                                           = 0x9303
	BUFFER_GPU_ADDRESS_NV                                      = 0x8F1D
	BUFFER_IMMUTABLE_STORAGE                                   = 0x821F
	BUFFER_KHR                                                 = 0x82E0
	BUFFER_MAPPED                                              = 0x88BC
	BUFFER_MAP_LENGTH                                          = 0x9120
	BUFFER_MAP_OFFSET                                          = 0x9121
	BUFFER_MAP_POINTER                                         = 0x88BD
	BUFFER_OBJECT_EXT                                          = 0x9151
	BUFFER_SIZE                                                = 0x8764
	BUFFER_STORAGE_FLAGS                                       = 0x8220
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	BUFFER_USAGE                                               = 0x8765
	BUFFER_VARIABLE                                            = 0x92E5
	BYTE                                                       = 0x1400
	C3F_V3F                                                    = 0x2A24
	C4F_N3F_V3F                                                = 0x2A26
	C4UB_V2F                                                   = 0x2A22
	C4UB_V3F                                                   = 0x2A23
	CAVEAT_SUPPORT                                             = 0x82B8
	CCW                                                        = 0x0901
	CIRCULAR_CCW_ARC_TO_NV                                     = 0xF8
	CIRCULAR_CW_ARC_TO_NV                                      = 0xFA
	CIRCULAR_TANGENT_ARC_TO_NV                                 = 0xFC
	CLAMP                                                      = 0x2900
	CLAMP_FRAGMENT_COLOR                                       = 0x891B
	CLAMP_READ_COLOR                                           = 0x891C
	CLAMP_TO_BORDER                                            = 0x812D
	CLAMP_TO_BORDER_ARB                                        = 0x812D
	CLAMP_TO_EDGE                                              = 0x812F
	CLAMP_VERTEX_COLOR                                         = 0x891A
	CLEAR                                                      = 0x1500
	CLEAR_BUFFER                                               = 0x82B4
	CLEAR_TEXTURE                                              = 0x9365
	CLIENT_ACTIVE_TEXTURE                                      = 0x84E1
	CLIENT_ALL_ATTRIB_BITS                                     = 0xFFFFFFFF
	CLIENT_ATTRIB_STACK_DEPTH                                  = 0x0BB1
	CLIENT_MAPPED_BUFFER_BARRIER_BIT                           = 0x00004000
	CLIENT_PIXEL_STORE_BIT                                     = 0x00000001
	CLIENT_STORAGE_BIT                                         = 0x0200
	CLIENT_VERTEX_ARRAY_BIT                                    = 0x00000002
	CLIPPING_INPUT_PRIMITIVES                                  = 0x82F6
	CLIPPING_INPUT_PRIMITIVES_ARB                              = 0x82F6
	CLIPPING_OUTPUT_PRIMITIVES                                 = 0x82F7
	CLIPPING_OUTPUT_PRIMITIVES_ARB                             = 0x82F7
	CLIP_DEPTH_MODE                                            = 0x935D
	CLIP_DISTANCE0                                             = 0x3000
	CLIP_DISTANCE1                                             = 0x3001
	CLIP_DISTANCE2                                             = 0x3002
	CLIP_DISTANCE3                                             = 0x3003
	CLIP_DISTANCE4                                             = 0x3004
	CLIP_DISTANCE5                                             = 0x3005
	CLIP_DISTANCE6                                             = 0x3006
	CLIP_DISTANCE7                                             = 0x3007
	CLIP_ORIGIN                                                = 0x935C
	CLIP_PLANE0                                                = 0x3000
	CLIP_PLANE1                                                = 0x3001
	CLIP_PLANE2                                                = 0x3002
	CLIP_PLANE3                                                = 0x3003
	CLIP_PLANE4                                                = 0x3004
	CLIP_PLANE5                                                = 0x3005
	CLOSE_PATH_NV                                              = 0x00
	COEFF                                                      = 0x0A00
	COLOR                                                      = 0x1800
	COLORBURN_KHR                                              = 0x929A
	COLORBURN_NV                                               = 0x929A
	COLORDODGE_KHR                                             = 0x9299
	COLORDODGE_NV                                              = 0x9299
	COLOR_ARRAY                                                = 0x8076
	COLOR_ARRAY_ADDRESS_NV                                     = 0x8F23
	COLOR_ARRAY_BUFFER_BINDING                                 = 0x8898
	COLOR_ARRAY_LENGTH_NV                                      = 0x8F2D
	COLOR_ARRAY_POINTER                                        = 0x8090
	COLOR_ARRAY_SIZE                                           = 0x8081
	COLOR_ARRAY_STRIDE                                         = 0x8083
	COLOR_ARRAY_TYPE                                           = 0x8082
	COLOR_ATTACHMENT0                                          = 0x8CE0
	COLOR_ATTACHMENT1                                          = 0x8CE1
	COLOR_ATTACHMENT10                                         = 0x8CEA
	COLOR_ATTACHMENT11                                         = 0x8CEB
	COLOR_ATTACHMENT12                                         = 0x8CEC
	COLOR_ATTACHMENT13                                         = 0x8CED
	COLOR_ATTACHMENT14                                         = 0x8CEE
	COLOR_ATTACHMENT15                                         = 0x8CEF
	COLOR_ATTACHMENT16                                         = 0x8CF0
	COLOR_ATTACHMENT17                                         = 0x8CF1
	COLOR_ATTACHMENT18                                         = 0x8CF2
	COLOR_ATTACHMENT19                                         = 0x8CF3
	COLOR_ATTACHMENT2                                          = 0x8CE2
	COLOR_ATTACHMENT20                                         = 0x8CF4
	COLOR_ATTACHMENT21                                         = 0x8CF5
	COLOR_ATTACHMENT22                                         = 0x8CF6
	COLOR_ATTACHMENT23                                         = 0x8CF7
	COLOR_ATTACHMENT24                                         = 0x8CF8
	COLOR_ATTACHMENT25                                         = 0x8CF9
	COLOR_ATTACHMENT26                                         = 0x8CFA
	COLOR_ATTACHMENT27                                         = 0x8CFB
	COLOR_ATTACHMENT28                                         = 0x8CFC
	COLOR_ATTACHMENT29                                         = 0x8CFD
	COLOR_ATTACHMENT3                                          = 0x8CE3
	COLOR_ATTACHMENT30                                         = 0x8CFE
	COLOR_ATTACHMENT31                                         = 0x8CFF
	COLOR_ATTACHMENT4                                          = 0x8CE4
	COLOR_ATTACHMENT5                                          = 0x8CE5
	COLOR_ATTACHMENT6                                          = 0x8CE6
	COLOR_ATTACHMENT7                                          = 0x8CE7
	COLOR_ATTACHMENT8                                          = 0x8CE8
	COLOR_ATTACHMENT9                                          = 0x8CE9
	COLOR_BUFFER_BIT                                           = 0x00004000
	COLOR_CLEAR_VALUE                                          = 0x0C22
	COLOR_COMPONENTS                                           = 0x8283
	COLOR_ENCODING                                             = 0x8296
	COLOR_INDEX                                                = 0x1900
	COLOR_INDEXES                                              = 0x1603
	COLOR_LOGIC_OP                                             = 0x0BF2
	COLOR_MATERIAL                                             = 0x0B57
	COLOR_MATERIAL_FACE                                        = 0x0B55
	COLOR_MATERIAL_PARAMETER                                   = 0x0B56
	COLOR_RENDERABLE                                           = 0x8286
	COLOR_SAMPLES_NV                                           = 0x8E20
	COLOR_SUM                                                  = 0x8458
	COLOR_WRITEMASK                                            = 0x0C23
	COMBINE                                                    = 0x8570
	COMBINE_ALPHA                                              = 0x8572
	COMBINE_RGB                                                = 0x8571
	COMMAND_BARRIER_BIT                                        = 0x00000040
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	COMPARE_R_TO_TEXTURE                                       = 0x884E
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	COMPILE                                                    = 0x1300
	COMPILE_AND_EXECUTE                                        = 0x1301
	COMPILE_STATUS                                             = 0x8B81
	COMPLETION_STATUS_ARB                                      = 0x91B1
	COMPLETION_STATUS_KHR                                      = 0x91B1
	COMPRESSED_ALPHA                                           = 0x84E9
	COMPRESSED_INTENSITY                                       = 0x84EC
	COMPRESSED_LUMINANCE                                       = 0x84EA
	COMPRESSED_LUMINANCE_ALPHA                                 = 0x84EB
	COMPRESSED_R11_EAC                                         = 0x9270
	COMPRESSED_RED                                             = 0x8225
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	COMPRESSED_RG                                              = 0x8226
	COMPRESSED_RG11_EAC                                        = 0x9272
	COMPRESSED_RGB                                             = 0x84ED
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	COMPRESSED_RGBA                                            = 0x84EE
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	COMPRESSED_RGBA_ASTC_10x10_KHR                             = 0x93BB
	COMPRESSED_RGBA_ASTC_10x5_KHR                              = 0x93B8
	COMPRESSED_RGBA_ASTC_10x6_KHR                              = 0x93B9
	COMPRESSED_RGBA_ASTC_10x8_KHR                              = 0x93BA
	COMPRESSED_RGBA_ASTC_12x10_KHR                             = 0x93BC
	COMPRESSED_RGBA_ASTC_12x12_KHR                             = 0x93BD
	COMPRESSED_RGBA_ASTC_4x4_KHR                               = 0x93B0
	COMPRESSED_RGBA_ASTC_5x4_KHR                               = 0x93B1
	COMPRESSED_RGBA_ASTC_5x5_KHR                               = 0x93B2
	COMPRESSED_RGBA_ASTC_6x5_KHR                               = 0x93B3
	COMPRESSED_RGBA_ASTC_6x6_KHR                               = 0x93B4
	COMPRESSED_RGBA_ASTC_8x5_KHR                               = 0x93B5
	COMPRESSED_RGBA_ASTC_8x6_KHR                               = 0x93B6
	COMPRESSED_RGBA_ASTC_8x8_KHR                               = 0x93B7
	COMPRESSED_RGBA_BPTC_UNORM                                 = 0x8E8C
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             = 0x8E8C
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT                           = 0x8E8E
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       = 0x8E8E
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT                         = 0x8E8F
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     = 0x8E8F
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	COMPRESSED_SLUMINANCE                                      = 0x8C4A
	COMPRESSED_SLUMINANCE_ALPHA                                = 0x8C4B
	COMPRESSED_SRGB                                            = 0x8C48
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     = 0x93DB
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      = 0x93D8
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      = 0x93D9
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      = 0x93DA
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     = 0x93DC
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     = 0x93DD
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       = 0x93D0
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       = 0x93D1
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       = 0x93D2
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       = 0x93D3
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       = 0x93D4
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       = 0x93D5
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       = 0x93D6
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       = 0x93D7
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM                           = 0x8E8D
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       = 0x8E8D
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	COMPUTE_SHADER                                             = 0x91B9
	COMPUTE_SHADER_BIT                                         = 0x00000020
	COMPUTE_SHADER_INVOCATIONS                                 = 0x82F5
	COMPUTE_SHADER_INVOCATIONS_ARB                             = 0x82F5
	COMPUTE_SUBROUTINE                                         = 0x92ED
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	COMPUTE_TEXTURE                                            = 0x82A0
	COMPUTE_WORK_GROUP_SIZE                                    = 0x8267
	CONDITION_SATISFIED                                        = 0x911C
	CONFORMANT_NV                                              = 0x9374
	CONIC_CURVE_TO_NV                                          = 0x1A
	CONJOINT_NV                                                = 0x9284
	CONSERVATIVE_RASTERIZATION_INTEL                           = 0x83FE
	CONSERVATIVE_RASTERIZATION_NV                              = 0x9346
	CONSERVATIVE_RASTER_DILATE_GRANULARITY_NV                  = 0x937B
	CONSERVATIVE_RASTER_DILATE_NV                              = 0x9379
	CONSERVATIVE_RASTER_DILATE_RANGE_NV                        = 0x937A
	CONSERVATIVE_RASTER_MODE_NV                                = 0x954D
	CONSERVATIVE_RASTER_MODE_POST_SNAP_NV                      = 0x954E
	CONSERVATIVE_RASTER_MODE_PRE_SNAP_NV                       = 0x9550
	CONSERVATIVE_RASTER_MODE_PRE_SNAP_TRIANGLES_NV             = 0x954F
	CONSTANT                                                   = 0x8576
	CONSTANT_ALPHA                                             = 0x8003
	CONSTANT_ATTENUATION                                       = 0x1207
	CONSTANT_COLOR                                             = 0x8001
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	CONTEXT_FLAGS                                              = 0x821E
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	CONTEXT_FLAG_DEBUG_BIT_KHR                                 = 0x00000002
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x00000001
	CONTEXT_FLAG_NO_ERROR_BIT                                  = 0x00000008
	CONTEXT_FLAG_NO_ERROR_BIT_KHR                              = 0x00000008
	CONTEXT_FLAG_ROBUST_ACCESS_BIT                             = 0x00000004
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         = 0x00000004
	CONTEXT_LOST                                               = 0x0507
	CONTEXT_LOST_KHR                                           = 0x0507
	CONTEXT_PROFILE_MASK                                       = 0x9126
	CONTEXT_RELEASE_BEHAVIOR                                   = 0x82FB
	CONTEXT_RELEASE_BEHAVIOR_FLUSH                             = 0x82FC
	CONTEXT_RELEASE_BEHAVIOR_FLUSH_KHR                         = 0x82FC
	CONTEXT_RELEASE_BEHAVIOR_KHR                               = 0x82FB
	CONTEXT_ROBUST_ACCESS                                      = 0x90F3
	CONTEXT_ROBUST_ACCESS_KHR                                  = 0x90F3
	CONTRAST_NV                                                = 0x92A1
	CONVEX_HULL_NV                                             = 0x908B
	COORD_REPLACE                                              = 0x8862
	COPY                                                       = 0x1503
	COPY_INVERTED                                              = 0x150C
	COPY_PIXEL_TOKEN                                           = 0x0706
	COPY_READ_BUFFER                                           = 0x8F36
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	COPY_WRITE_BUFFER                                          = 0x8F37
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	COUNTER_RANGE_AMD                                          = 0x8BC1
	COUNTER_TYPE_AMD                                           = 0x8BC0
	COUNT_DOWN_NV                                              = 0x9089
	COUNT_UP_NV                                                = 0x9088
	COVERAGE_MODULATION_NV                                     = 0x9332
	COVERAGE_MODULATION_TABLE_NV                               = 0x9331
	COVERAGE_MODULATION_TABLE_SIZE_NV                          = 0x9333
	CUBIC_CURVE_TO_NV                                          = 0x0C
	CULL_FACE                                                  = 0x0B44
	CULL_FACE_MODE                                             = 0x0B45
	CURRENT_BIT                                                = 0x00000001
	CURRENT_COLOR                                              = 0x0B00
	CURRENT_FOG_COORD                                          = 0x8453
	CURRENT_FOG_COORDINATE                                     = 0x8453
	CURRENT_INDEX                                              = 0x0B01
	CURRENT_NORMAL                                             = 0x0B02
	CURRENT_PROGRAM                                            = 0x8B8D
	CURRENT_QUERY                                              = 0x8865
	CURRENT_RASTER_COLOR                                       = 0x0B04
	CURRENT_RASTER_DISTANCE                                    = 0x0B09
	CURRENT_RASTER_INDEX                                       = 0x0B05
	CURRENT_RASTER_POSITION                                    = 0x0B07
	CURRENT_RASTER_POSITION_VALID                              = 0x0B08
	CURRENT_RASTER_SECONDARY_COLOR                             = 0x845F
	CURRENT_RASTER_TEXTURE_COORDS                              = 0x0B06
	CURRENT_SECONDARY_COLOR                                    = 0x8459
	CURRENT_TEXTURE_COORDS                                     = 0x0B03
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	CW                                                         = 0x0900
	DARKEN_KHR                                                 = 0x9297
	DARKEN_NV                                                  = 0x9297
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	DEBUG_CALLBACK_FUNCTION_ARB                                = 0x8244
	DEBUG_CALLBACK_FUNCTION_KHR                                = 0x8244
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	DEBUG_CALLBACK_USER_PARAM_ARB                              = 0x8245
	DEBUG_CALLBACK_USER_PARAM_KHR                              = 0x8245
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	DEBUG_GROUP_STACK_DEPTH_KHR                                = 0x826D
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	DEBUG_LOGGED_MESSAGES_ARB                                  = 0x9145
	DEBUG_LOGGED_MESSAGES_KHR                                  = 0x9145
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       = 0x8243
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_KHR                       = 0x8243
	DEBUG_OUTPUT                                               = 0x92E0
	DEBUG_OUTPUT_KHR                                           = 0x92E0
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               = 0x8242
	DEBUG_OUTPUT_SYNCHRONOUS_KHR                               = 0x8242
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	DEBUG_SEVERITY_HIGH_ARB                                    = 0x9146
	DEBUG_SEVERITY_HIGH_KHR                                    = 0x9146
	DEBUG_SEVERITY_LOW                                         = 0x9148
	DEBUG_SEVERITY_LOW_ARB                                     = 0x9148
	DEBUG_SEVERITY_LOW_KHR                                     = 0x9148
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	DEBUG_SEVERITY_MEDIUM_ARB                                  = 0x9147
	DEBUG_SEVERITY_MEDIUM_KHR                                  = 0x9147
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	DEBUG_SEVERITY_NOTIFICATION_KHR                            = 0x826B
	DEBUG_SOURCE_API                                           = 0x8246
	DEBUG_SOURCE_API_ARB                                       = 0x8246
	DEBUG_SOURCE_API_KHR                                       = 0x8246
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	DEBUG_SOURCE_APPLICATION_ARB                               = 0x824A
	DEBUG_SOURCE_APPLICATION_KHR                               = 0x824A
	DEBUG_SOURCE_OTHER                                         = 0x824B
	DEBUG_SOURCE_OTHER_ARB                                     = 0x824B
	DEBUG_SOURCE_OTHER_KHR                                     = 0x824B
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           = 0x8248
	DEBUG_SOURCE_SHADER_COMPILER_KHR                           = 0x8248
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	DEBUG_SOURCE_THIRD_PARTY_ARB                               = 0x8249
	DEBUG_SOURCE_THIRD_PARTY_KHR                               = 0x8249
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             = 0x8247
	DEBUG_SOURCE_WINDOW_SYSTEM_KHR                             = 0x8247
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         = 0x824D
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_KHR                         = 0x824D
	DEBUG_TYPE_ERROR                                           = 0x824C
	DEBUG_TYPE_ERROR_ARB                                       = 0x824C
	DEBUG_TYPE_ERROR_KHR                                       = 0x824C
	DEBUG_TYPE_MARKER                                          = 0x8268
	DEBUG_TYPE_MARKER_KHR                                      = 0x8268
	DEBUG_TYPE_OTHER                                           = 0x8251
	DEBUG_TYPE_OTHER_ARB                                       = 0x8251
	DEBUG_TYPE_OTHER_KHR                                       = 0x8251
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	DEBUG_TYPE_PERFORMANCE_ARB                                 = 0x8250
	DEBUG_TYPE_PERFORMANCE_KHR                                 = 0x8250
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	DEBUG_TYPE_POP_GROUP_KHR                                   = 0x826A
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	DEBUG_TYPE_PORTABILITY_ARB                                 = 0x824F
	DEBUG_TYPE_PORTABILITY_KHR                                 = 0x824F
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	DEBUG_TYPE_PUSH_GROUP_KHR                                  = 0x8269
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          = 0x824E
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_KHR                          = 0x824E
	DECAL                                                      = 0x2101
	DECODE_EXT                                                 = 0x8A49
	DECR                                                       = 0x1E03
	DECR_WRAP                                                  = 0x8508
	DELETE_STATUS                                              = 0x8B80
	DEPTH                                                      = 0x1801
	DEPTH24_STENCIL8                                           = 0x88F0
	DEPTH32F_STENCIL8                                          = 0x8CAD
	DEPTH_ATTACHMENT                                           = 0x8D00
	DEPTH_BIAS                                                 = 0x0D1F
	DEPTH_BITS                                                 = 0x0D56
	DEPTH_BUFFER_BIT                                           = 0x00000100
	DEPTH_CLAMP                                                = 0x864F
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	DEPTH_COMPONENT                                            = 0x1902
	DEPTH_COMPONENT16                                          = 0x81A5
	DEPTH_COMPONENT24                                          = 0x81A6
	DEPTH_COMPONENT32                                          = 0x81A7
	DEPTH_COMPONENT32F                                         = 0x8CAC
	DEPTH_COMPONENTS                                           = 0x8284
	DEPTH_FUNC                                                 = 0x0B74
	DEPTH_RANGE                                                = 0x0B70
	DEPTH_RENDERABLE                                           = 0x8287
	DEPTH_SAMPLES_NV                                           = 0x932D
	DEPTH_SCALE                                                = 0x0D1E
	DEPTH_STENCIL                                              = 0x84F9
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	DEPTH_TEST                                                 = 0x0B71
	DEPTH_TEXTURE_MODE                                         = 0x884B
	DEPTH_WRITEMASK                                            = 0x0B72
	DIFFERENCE_KHR                                             = 0x929E
	DIFFERENCE_NV                                              = 0x929E
	DIFFUSE                                                    = 0x1201
	DISJOINT_NV                                                = 0x9283
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	DITHER                                                     = 0x0BD0
	DOMAIN                                                     = 0x0A02
	DONT_CARE                                                  = 0x1100
	DOT3_RGB                                                   = 0x86AE
	DOT3_RGBA                                                  = 0x86AF
	DOUBLE                                                     = 0x140A
	DOUBLEBUFFER                                               = 0x0C32
	DOUBLE_MAT2                                                = 0x8F46
	DOUBLE_MAT2x3                                              = 0x8F49
	DOUBLE_MAT2x4                                              = 0x8F4A
	DOUBLE_MAT3                                                = 0x8F47
	DOUBLE_MAT3x2                                              = 0x8F4B
	DOUBLE_MAT3x4                                              = 0x8F4C
	DOUBLE_MAT4                                                = 0x8F48
	DOUBLE_MAT4x2                                              = 0x8F4D
	DOUBLE_MAT4x3                                              = 0x8F4E
	DOUBLE_VEC2                                                = 0x8FFC
	DOUBLE_VEC3                                                = 0x8FFD
	DOUBLE_VEC4                                                = 0x8FFE
	DRAW_ARRAYS_COMMAND_NV                                     = 0x0003
	DRAW_ARRAYS_INSTANCED_COMMAND_NV                           = 0x0007
	DRAW_ARRAYS_STRIP_COMMAND_NV                               = 0x0005
	DRAW_BUFFER                                                = 0x0C01
	DRAW_BUFFER0                                               = 0x8825
	DRAW_BUFFER1                                               = 0x8826
	DRAW_BUFFER10                                              = 0x882F
	DRAW_BUFFER11                                              = 0x8830
	DRAW_BUFFER12                                              = 0x8831
	DRAW_BUFFER13                                              = 0x8832
	DRAW_BUFFER14                                              = 0x8833
	DRAW_BUFFER15                                              = 0x8834
	DRAW_BUFFER2                                               = 0x8827
	DRAW_BUFFER3                                               = 0x8828
	DRAW_BUFFER4                                               = 0x8829
	DRAW_BUFFER5                                               = 0x882A
	DRAW_BUFFER6                                               = 0x882B
	DRAW_BUFFER7                                               = 0x882C
	DRAW_BUFFER8                                               = 0x882D
	DRAW_BUFFER9                                               = 0x882E
	DRAW_ELEMENTS_COMMAND_NV                                   = 0x0002
	DRAW_ELEMENTS_INSTANCED_COMMAND_NV                         = 0x0006
	DRAW_ELEMENTS_STRIP_COMMAND_NV                             = 0x0004
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	DRAW_FRAMEBUFFER_BINDING                                   = 0x8CA6
	DRAW_INDIRECT_ADDRESS_NV                                   = 0x8F41
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	DRAW_INDIRECT_LENGTH_NV                                    = 0x8F42
	DRAW_INDIRECT_UNIFIED_NV                                   = 0x8F40
	DRAW_PIXEL_TOKEN                                           = 0x0705
	DST_ALPHA                                                  = 0x0304
	DST_ATOP_NV                                                = 0x928F
	DST_COLOR                                                  = 0x0306
	DST_IN_NV                                                  = 0x928B
	DST_NV                                                     = 0x9287
	DST_OUT_NV                                                 = 0x928D
	DST_OVER_NV                                                = 0x9289
	DUP_FIRST_CUBIC_CURVE_TO_NV                                = 0xF2
	DUP_LAST_CUBIC_CURVE_TO_NV                                 = 0xF4
	DYNAMIC_COPY                                               = 0x88EA
	DYNAMIC_DRAW                                               = 0x88E8
	DYNAMIC_READ                                               = 0x88E9
	DYNAMIC_STORAGE_BIT                                        = 0x0100
	EDGE_FLAG                                                  = 0x0B43
	EDGE_FLAG_ARRAY                                            = 0x8079
	EDGE_FLAG_ARRAY_ADDRESS_NV                                 = 0x8F26
	EDGE_FLAG_ARRAY_BUFFER_BINDING                             = 0x889B
	EDGE_FLAG_ARRAY_LENGTH_NV                                  = 0x8F30
	EDGE_FLAG_ARRAY_POINTER                                    = 0x8093
	EDGE_FLAG_ARRAY_STRIDE                                     = 0x808C
	EFFECTIVE_RASTER_SAMPLES_EXT                               = 0x932C
	ELEMENT_ADDRESS_COMMAND_NV                                 = 0x0008
	ELEMENT_ARRAY_ADDRESS_NV                                   = 0x8F29
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	ELEMENT_ARRAY_LENGTH_NV                                    = 0x8F33
	ELEMENT_ARRAY_UNIFIED_NV                                   = 0x8F1F
	EMISSION                                                   = 0x1600
	ENABLE_BIT                                                 = 0x00002000
	EQUAL                                                      = 0x0202
	EQUIV                                                      = 0x1509
	EVAL_BIT                                                   = 0x00010000
	EXCLUSION_KHR                                              = 0x92A0
	EXCLUSION_NV                                               = 0x92A0
	EXCLUSIVE_EXT                                              = 0x8F11
	EXP                                                        = 0x0800
	EXP2                                                       = 0x0801
	EXTENSIONS                                                 = 0x1F03
	EYE_LINEAR                                                 = 0x2400
	EYE_PLANE                                                  = 0x2502
	FACTOR_MAX_AMD                                             = 0x901D
	FACTOR_MIN_AMD                                             = 0x901C
	FALSE                                                      = 0
	FASTEST                                                    = 0x1101
	FEEDBACK                                                   = 0x1C01
	FEEDBACK_BUFFER_POINTER                                    = 0x0DF0
	FEEDBACK_BUFFER_SIZE                                       = 0x0DF1
	FEEDBACK_BUFFER_TYPE                                       = 0x0DF2
	FILE_NAME_NV                                               = 0x9074
	FILL                                                       = 0x1B02
	FILL_RECTANGLE_NV                                          = 0x933C
	FILTER                                                     = 0x829A
	FIRST_TO_REST_NV                                           = 0x90AF
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	FIXED                                                      = 0x140C
	FIXED_ONLY                                                 = 0x891D
	FLAT                                                       = 0x1D00
	FLOAT                                                      = 0x1406
	FLOAT16_NV                                                 = 0x8FF8
	FLOAT16_VEC2_NV                                            = 0x8FF9
	FLOAT16_VEC3_NV                                            = 0x8FFA
	FLOAT16_VEC4_NV                                            = 0x8FFB
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	FLOAT_MAT2                                                 = 0x8B5A
	FLOAT_MAT2x3                                               = 0x8B65
	FLOAT_MAT2x4                                               = 0x8B66
	FLOAT_MAT3                                                 = 0x8B5B
	FLOAT_MAT3x2                                               = 0x8B67
	FLOAT_MAT3x4                                               = 0x8B68
	FLOAT_MAT4                                                 = 0x8B5C
	FLOAT_MAT4x2                                               = 0x8B69
	FLOAT_MAT4x3                                               = 0x8B6A
	FLOAT_VEC2                                                 = 0x8B50
	FLOAT_VEC3                                                 = 0x8B51
	FLOAT_VEC4                                                 = 0x8B52
	FOG                                                        = 0x0B60
	FOG_BIT                                                    = 0x00000080
	FOG_COLOR                                                  = 0x0B66
	FOG_COORD                                                  = 0x8451
	FOG_COORDINATE                                             = 0x8451
	FOG_COORDINATE_ARRAY                                       = 0x8457
	FOG_COORDINATE_ARRAY_BUFFER_BINDING                        = 0x889D
	FOG_COORDINATE_ARRAY_POINTER                               = 0x8456
	FOG_COORDINATE_ARRAY_STRIDE                                = 0x8455
	FOG_COORDINATE_ARRAY_TYPE                                  = 0x8454
	FOG_COORDINATE_SOURCE                                      = 0x8450
	FOG_COORD_ARRAY                                            = 0x8457
	FOG_COORD_ARRAY_ADDRESS_NV                                 = 0x8F28
	FOG_COORD_ARRAY_BUFFER_BINDING                             = 0x889D
	FOG_COORD_ARRAY_LENGTH_NV                                  = 0x8F32
	FOG_COORD_ARRAY_POINTER                                    = 0x8456
	FOG_COORD_ARRAY_STRIDE                                     = 0x8455
	FOG_COORD_ARRAY_TYPE                                       = 0x8454
	FOG_COORD_SRC                                              = 0x8450
	FOG_DENSITY                                                = 0x0B62
	FOG_END                                                    = 0x0B64
	FOG_HINT                                                   = 0x0C54
	FOG_INDEX                                                  = 0x0B61
	FOG_MODE                                                   = 0x0B65
	FOG_START                                                  = 0x0B63
	FONT_ASCENDER_BIT_NV                                       = 0x00200000
	FONT_DESCENDER_BIT_NV                                      = 0x00400000
	FONT_GLYPHS_AVAILABLE_NV                                   = 0x9368
	FONT_HAS_KERNING_BIT_NV                                    = 0x10000000
	FONT_HEIGHT_BIT_NV                                         = 0x00800000
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                             = 0x02000000
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                              = 0x01000000
	FONT_NUM_GLYPH_INDICES_BIT_NV                              = 0x20000000
	FONT_TARGET_UNAVAILABLE_NV                                 = 0x9369
	FONT_UNAVAILABLE_NV                                        = 0x936A
	FONT_UNDERLINE_POSITION_BIT_NV                             = 0x04000000
	FONT_UNDERLINE_THICKNESS_BIT_NV                            = 0x08000000
	FONT_UNINTELLIGIBLE_NV                                     = 0x936B
	FONT_UNITS_PER_EM_BIT_NV                                   = 0x00100000
	FONT_X_MAX_BOUNDS_BIT_NV                                   = 0x00040000
	FONT_X_MIN_BOUNDS_BIT_NV                                   = 0x00010000
	FONT_Y_MAX_BOUNDS_BIT_NV                                   = 0x00080000
	FONT_Y_MIN_BOUNDS_BIT_NV                                   = 0x00020000
	FRACTIONAL_EVEN                                            = 0x8E7C
	FRACTIONAL_ODD                                             = 0x8E7B
	FRAGMENT_COVERAGE_COLOR_NV                                 = 0x92DE
	FRAGMENT_COVERAGE_TO_COLOR_NV                              = 0x92DD
	FRAGMENT_DEPTH                                             = 0x8452
	FRAGMENT_INPUT_NV                                          = 0x936D
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	FRAGMENT_SHADER                                            = 0x8B30
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	FRAGMENT_SHADER_BIT_EXT                                    = 0x00000002
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                       = 0x8A52
	FRAGMENT_SHADER_INVOCATIONS                                = 0x82F4
	FRAGMENT_SHADER_INVOCATIONS_ARB                            = 0x82F4
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	FRAGMENT_TEXTURE                                           = 0x829F
	FRAMEBUFFER                                                = 0x8D40
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	FRAMEBUFFER_ATTACHMENT_LAYERED_ARB                         = 0x8DA7
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	FRAMEBUFFER_ATTACHMENT_TEXTURE_BASE_VIEW_INDEX_OVR         = 0x9632
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	FRAMEBUFFER_ATTACHMENT_TEXTURE_NUM_VIEWS_OVR               = 0x9630
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	FRAMEBUFFER_BLEND                                          = 0x828B
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB                     = 0x8DA9
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB                   = 0x8DA8
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	FRAMEBUFFER_INCOMPLETE_VIEW_TARGETS_OVR                    = 0x9633
	FRAMEBUFFER_PROGRAMMABLE_SAMPLE_LOCATIONS_ARB              = 0x9342
	FRAMEBUFFER_PROGRAMMABLE_SAMPLE_LOCATIONS_NV               = 0x9342
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	FRAMEBUFFER_SAMPLE_LOCATION_PIXEL_GRID_ARB                 = 0x9343
	FRAMEBUFFER_SAMPLE_LOCATION_PIXEL_GRID_NV                  = 0x9343
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	FRONT                                                      = 0x0404
	FRONT_AND_BACK                                             = 0x0408
	FRONT_FACE                                                 = 0x0B46
	FRONT_FACE_COMMAND_NV                                      = 0x0012
	FRONT_LEFT                                                 = 0x0400
	FRONT_RIGHT                                                = 0x0401
	FULL_SUPPORT                                               = 0x82B7
	FUNC_ADD                                                   = 0x8006
	FUNC_REVERSE_SUBTRACT                                      = 0x800B
	FUNC_SUBTRACT                                              = 0x800A
	GENERATE_MIPMAP                                            = 0x8191
	GENERATE_MIPMAP_HINT                                       = 0x8192
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	GEOMETRY_INPUT_TYPE_ARB                                    = 0x8DDB
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	GEOMETRY_OUTPUT_TYPE_ARB                                   = 0x8DDC
	GEOMETRY_SHADER                                            = 0x8DD9
	GEOMETRY_SHADER_ARB                                        = 0x8DD9
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	GEOMETRY_SHADER_PRIMITIVES_EMITTED                         = 0x82F3
	GEOMETRY_SHADER_PRIMITIVES_EMITTED_ARB                     = 0x82F3
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	GEOMETRY_TEXTURE                                           = 0x829E
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	GEOMETRY_VERTICES_OUT_ARB                                  = 0x8DDA
	GEQUAL                                                     = 0x0206
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	GLYPH_HAS_KERNING_BIT_NV                                   = 0x100
	GLYPH_HEIGHT_BIT_NV                                        = 0x02
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                    = 0x10
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                          = 0x04
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                          = 0x08
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                      = 0x80
	GLYPH_VERTICAL_BEARING_X_BIT_NV                            = 0x20
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                            = 0x40
	GLYPH_WIDTH_BIT_NV                                         = 0x01
	GPU_ADDRESS_NV                                             = 0x8F34
	GREATER                                                    = 0x0204
	GREEN                                                      = 0x1904
	GREEN_BIAS                                                 = 0x0D19
	GREEN_BITS                                                 = 0x0D53
	GREEN_INTEGER                                              = 0x8D95
	GREEN_NV                                                   = 0x1904
	GREEN_SCALE                                                = 0x0D18
	GUILTY_CONTEXT_RESET                                       = 0x8253
	GUILTY_CONTEXT_RESET_ARB                                   = 0x8253
	GUILTY_CONTEXT_RESET_KHR                                   = 0x8253
	HALF_FLOAT                                                 = 0x140B
	HARDLIGHT_KHR                                              = 0x929B
	HARDLIGHT_NV                                               = 0x929B
	HARDMIX_NV                                                 = 0x92A9
	HIGH_FLOAT                                                 = 0x8DF2
	HIGH_INT                                                   = 0x8DF5
	HINT_BIT                                                   = 0x00008000
	HORIZONTAL_LINE_TO_NV                                      = 0x06
	HSL_COLOR_KHR                                              = 0x92AF
	HSL_COLOR_NV                                               = 0x92AF
	HSL_HUE_KHR                                                = 0x92AD
	HSL_HUE_NV                                                 = 0x92AD
	HSL_LUMINOSITY_KHR                                         = 0x92B0
	HSL_LUMINOSITY_NV                                          = 0x92B0
	HSL_SATURATION_KHR                                         = 0x92AE
	HSL_SATURATION_NV                                          = 0x92AE
	IMAGE_1D                                                   = 0x904C
	IMAGE_1D_ARRAY                                             = 0x9052
	IMAGE_2D                                                   = 0x904D
	IMAGE_2D_ARRAY                                             = 0x9053
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	IMAGE_2D_RECT                                              = 0x904F
	IMAGE_3D                                                   = 0x904E
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	IMAGE_BINDING_FORMAT                                       = 0x906E
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	IMAGE_BINDING_NAME                                         = 0x8F3A
	IMAGE_BUFFER                                               = 0x9051
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	IMAGE_CUBE                                                 = 0x9050
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	INCLUSIVE_EXT                                              = 0x8F10
	INCR                                                       = 0x1E02
	INCR_WRAP                                                  = 0x8507
	INDEX_ARRAY                                                = 0x8077
	INDEX_ARRAY_ADDRESS_NV                                     = 0x8F24
	INDEX_ARRAY_BUFFER_BINDING                                 = 0x8899
	INDEX_ARRAY_LENGTH_NV                                      = 0x8F2E
	INDEX_ARRAY_POINTER                                        = 0x8091
	INDEX_ARRAY_STRIDE                                         = 0x8086
	INDEX_ARRAY_TYPE                                           = 0x8085
	INDEX_BITS                                                 = 0x0D51
	INDEX_CLEAR_VALUE                                          = 0x0C20
	INDEX_LOGIC_OP                                             = 0x0BF1
	INDEX_MODE                                                 = 0x0C30
	INDEX_OFFSET                                               = 0x0D13
	INDEX_SHIFT                                                = 0x0D12
	INDEX_WRITEMASK                                            = 0x0C21
	INFO_LOG_LENGTH                                            = 0x8B84
	INNOCENT_CONTEXT_RESET                                     = 0x8254
	INNOCENT_CONTEXT_RESET_ARB                                 = 0x8254
	INNOCENT_CONTEXT_RESET_KHR                                 = 0x8254
	INT                                                        = 0x1404
	INT16_NV                                                   = 0x8FE4
	INT16_VEC2_NV                                              = 0x8FE5
	INT16_VEC3_NV                                              = 0x8FE6
	INT16_VEC4_NV                                              = 0x8FE7
	INT64_ARB                                                  = 0x140E
	INT64_NV                                                   = 0x140E
	INT64_VEC2_ARB                                             = 0x8FE9
	INT64_VEC2_NV                                              = 0x8FE9
	INT64_VEC3_ARB                                             = 0x8FEA
	INT64_VEC3_NV                                              = 0x8FEA
	INT64_VEC4_ARB                                             = 0x8FEB
	INT64_VEC4_NV                                              = 0x8FEB
	INT8_NV                                                    = 0x8FE0
	INT8_VEC2_NV                                               = 0x8FE1
	INT8_VEC3_NV                                               = 0x8FE2
	INT8_VEC4_NV                                               = 0x8FE3
	INTENSITY                                                  = 0x8049
	INTENSITY12                                                = 0x804C
	INTENSITY16                                                = 0x804D
	INTENSITY4                                                 = 0x804A
	INTENSITY8                                                 = 0x804B
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	INTERPOLATE                                                = 0x8575
	INT_2_10_10_10_REV                                         = 0x8D9F
	INT_IMAGE_1D                                               = 0x9057
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	INT_IMAGE_2D                                               = 0x9058
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	INT_IMAGE_2D_RECT                                          = 0x905A
	INT_IMAGE_3D                                               = 0x9059
	INT_IMAGE_BUFFER                                           = 0x905C
	INT_IMAGE_CUBE                                             = 0x905B
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	INT_SAMPLER_1D                                             = 0x8DC9
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	INT_SAMPLER_2D                                             = 0x8DCA
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	INT_SAMPLER_3D                                             = 0x8DCB
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	INT_SAMPLER_CUBE                                           = 0x8DCC
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	INT_SAMPLER_CUBE_MAP_ARRAY_ARB                             = 0x900E
	INT_VEC2                                                   = 0x8B53
	INT_VEC3                                                   = 0x8B54
	INT_VEC4                                                   = 0x8B55
	INVALID_ENUM                                               = 0x0500
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	INVALID_INDEX                                              = 0xFFFFFFFF
	INVALID_OPERATION                                          = 0x0502
	INVALID_VALUE                                              = 0x0501
	INVERT                                                     = 0x150A
	INVERT_OVG_NV                                              = 0x92B4
	INVERT_RGB_NV                                              = 0x92A3
	ISOLINES                                                   = 0x8E7A
	IS_PER_PATCH                                               = 0x92E7
	IS_ROW_MAJOR                                               = 0x9300
	ITALIC_BIT_NV                                              = 0x02
	KEEP                                                       = 0x1E00
	LARGE_CCW_ARC_TO_NV                                        = 0x16
	LARGE_CW_ARC_TO_NV                                         = 0x18
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	LEFT                                                       = 0x0406
	LEQUAL                                                     = 0x0203
	LESS                                                       = 0x0201
	LIGHT0                                                     = 0x4000
	LIGHT1                                                     = 0x4001
	LIGHT2                                                     = 0x4002
	LIGHT3                                                     = 0x4003
	LIGHT4                                                     = 0x4004
	LIGHT5                                                     = 0x4005
	LIGHT6                                                     = 0x4006
	LIGHT7                                                     = 0x4007
	LIGHTEN_KHR                                                = 0x9298
	LIGHTEN_NV                                                 = 0x9298
	LIGHTING                                                   = 0x0B50
	LIGHTING_BIT                                               = 0x00000040
	LIGHT_MODEL_AMBIENT                                        = 0x0B53
	LIGHT_MODEL_COLOR_CONTROL                                  = 0x81F8
	LIGHT_MODEL_LOCAL_VIEWER                                   = 0x0B51
	LIGHT_MODEL_TWO_SIDE                                       = 0x0B52
	LINE                                                       = 0x1B01
	LINEAR                                                     = 0x2601
	LINEARBURN_NV                                              = 0x92A5
	LINEARDODGE_NV                                             = 0x92A4
	LINEARLIGHT_NV                                             = 0x92A7
	LINEAR_ATTENUATION                                         = 0x1208
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	LINES                                                      = 0x0001
	LINES_ADJACENCY                                            = 0x000A
	LINES_ADJACENCY_ARB                                        = 0x000A
	LINE_BIT                                                   = 0x00000004
	LINE_LOOP                                                  = 0x0002
	LINE_RESET_TOKEN                                           = 0x0707
	LINE_SMOOTH                                                = 0x0B20
	LINE_SMOOTH_HINT                                           = 0x0C52
	LINE_STIPPLE                                               = 0x0B24
	LINE_STIPPLE_PATTERN                                       = 0x0B25
	LINE_STIPPLE_REPEAT                                        = 0x0B26
	LINE_STRIP                                                 = 0x0003
	LINE_STRIP_ADJACENCY                                       = 0x000B
	LINE_STRIP_ADJACENCY_ARB                                   = 0x000B
	LINE_TOKEN                                                 = 0x0702
	LINE_TO_NV                                                 = 0x04
	LINE_WIDTH                                                 = 0x0B21
	LINE_WIDTH_COMMAND_NV                                      = 0x000D
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	LINE_WIDTH_RANGE                                           = 0x0B22
	LINK_STATUS                                                = 0x8B82
	LIST_BASE                                                  = 0x0B32
	LIST_BIT                                                   = 0x00020000
	LIST_INDEX                                                 = 0x0B33
	LIST_MODE                                                  = 0x0B30
	LOAD                                                       = 0x0101
	LOCATION                                                   = 0x930E
	LOCATION_COMPONENT                                         = 0x934A
	LOCATION_INDEX                                             = 0x930F
	LOGIC_OP                                                   = 0x0BF1
	LOGIC_OP_MODE                                              = 0x0BF0
	LOSE_CONTEXT_ON_RESET                                      = 0x8252
	LOSE_CONTEXT_ON_RESET_ARB                                  = 0x8252
	LOSE_CONTEXT_ON_RESET_KHR                                  = 0x8252
	LOWER_LEFT                                                 = 0x8CA1
	LOW_FLOAT                                                  = 0x8DF0
	LOW_INT                                                    = 0x8DF3
	LUMINANCE                                                  = 0x1909
	LUMINANCE12                                                = 0x8041
	LUMINANCE12_ALPHA12                                        = 0x8047
	LUMINANCE12_ALPHA4                                         = 0x8046
	LUMINANCE16                                                = 0x8042
	LUMINANCE16_ALPHA16                                        = 0x8048
	LUMINANCE4                                                 = 0x803F
	LUMINANCE4_ALPHA4                                          = 0x8043
	LUMINANCE6_ALPHA2                                          = 0x8044
	LUMINANCE8                                                 = 0x8040
	LUMINANCE8_ALPHA8                                          = 0x8045
	LUMINANCE_ALPHA                                            = 0x190A
	MAJOR_VERSION                                              = 0x821B
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	MAP1_COLOR_4                                               = 0x0D90
	MAP1_GRID_DOMAIN                                           = 0x0DD0
	MAP1_GRID_SEGMENTS                                         = 0x0DD1
	MAP1_INDEX                                                 = 0x0D91
	MAP1_NORMAL                                                = 0x0D92
	MAP1_TEXTURE_COORD_1                                       = 0x0D93
	MAP1_TEXTURE_COORD_2                                       = 0x0D94
	MAP1_TEXTURE_COORD_3                                       = 0x0D95
	MAP1_TEXTURE_COORD_4                                       = 0x0D96
	MAP1_VERTEX_3                                              = 0x0D97
	MAP1_VERTEX_4                                              = 0x0D98
	MAP2_COLOR_4                                               = 0x0DB0
	MAP2_GRID_DOMAIN                                           = 0x0DD2
	MAP2_GRID_SEGMENTS                                         = 0x0DD3
	MAP2_INDEX                                                 = 0x0DB1
	MAP2_NORMAL                                                = 0x0DB2
	MAP2_TEXTURE_COORD_1                                       = 0x0DB3
	MAP2_TEXTURE_COORD_2                                       = 0x0DB4
	MAP2_TEXTURE_COORD_3                                       = 0x0DB5
	MAP2_TEXTURE_COORD_4                                       = 0x0DB6
	MAP2_VERTEX_3                                              = 0x0DB7
	MAP2_VERTEX_4                                              = 0x0DB8
	MAP_COHERENT_BIT                                           = 0x0080
	MAP_COLOR                                                  = 0x0D10
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	MAP_PERSISTENT_BIT                                         = 0x0040
	MAP_READ_BIT                                               = 0x0001
	MAP_STENCIL                                                = 0x0D11
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	MAP_WRITE_BIT                                              = 0x0002
	MATRIX_MODE                                                = 0x0BA0
	MATRIX_STRIDE                                              = 0x92FF
	MAX                                                        = 0x8008
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	MAX_ATTRIB_STACK_DEPTH                                     = 0x0D35
	MAX_CLIENT_ATTRIB_STACK_DEPTH                              = 0x0D3B
	MAX_CLIP_DISTANCES                                         = 0x0D32
	MAX_CLIP_PLANES                                            = 0x0D32
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	MAX_COMBINED_CLIP_AND_CULL_DISTANCES                       = 0x82FA
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   = 0x8A32
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = 0x8F39
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	MAX_COMPUTE_FIXED_GROUP_INVOCATIONS_ARB                    = 0x90EB
	MAX_COMPUTE_FIXED_GROUP_SIZE_ARB                           = 0x91BF
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	MAX_COMPUTE_VARIABLE_GROUP_INVOCATIONS_ARB                 = 0x9344
	MAX_COMPUTE_VARIABLE_GROUP_SIZE_ARB                        = 0x9345
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	MAX_COMPUTE_WORK_GROUP_INVOCATIONS                         = 0x90EB
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	MAX_CULL_DISTANCES                                         = 0x82F9
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	MAX_DEBUG_GROUP_STACK_DEPTH_KHR                            = 0x826C
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              = 0x9144
	MAX_DEBUG_LOGGED_MESSAGES_KHR                              = 0x9144
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               = 0x9143
	MAX_DEBUG_MESSAGE_LENGTH_KHR                               = 0x9143
	MAX_DEPTH                                                  = 0x8280
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	MAX_DRAW_BUFFERS                                           = 0x8824
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	MAX_EVAL_ORDER                                             = 0x0D30
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	MAX_GEOMETRY_OUTPUT_VERTICES_ARB                           = 0x8DE0
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB                       = 0x8C29
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB                   = 0x8DE1
	MAX_GEOMETRY_UNIFORM_BLOCKS                                = 0x8A2C
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB                        = 0x8DDF
	MAX_GEOMETRY_VARYING_COMPONENTS_ARB                        = 0x8DDD
	MAX_HEIGHT                                                 = 0x827F
	MAX_IMAGE_SAMPLES                                          = 0x906D
	MAX_IMAGE_UNITS                                            = 0x8F38
	MAX_INTEGER_SAMPLES                                        = 0x9110
	MAX_LABEL_LENGTH                                           = 0x82E8
	MAX_LABEL_LENGTH_KHR                                       = 0x82E8
	MAX_LAYERS                                                 = 0x8281
	MAX_LIGHTS                                                 = 0x0D31
	MAX_LIST_NESTING                                           = 0x0B31
	MAX_MODELVIEW_STACK_DEPTH                                  = 0x0D36
	MAX_MULTISAMPLE_COVERAGE_MODES_NV                          = 0x8E11
	MAX_NAME_LENGTH                                            = 0x92F6
	MAX_NAME_STACK_DEPTH                                       = 0x0D37
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	MAX_PATCH_VERTICES                                         = 0x8E7D
	MAX_PIXEL_MAP_TABLE                                        = 0x0D34
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  = 0x8F9F
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5F
	MAX_PROJECTION_STACK_DEPTH                                 = 0x0D38
	MAX_RASTER_SAMPLES_EXT                                     = 0x9329
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	MAX_SAMPLES                                                = 0x8D57
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	MAX_SHADER_BUFFER_ADDRESS_NV                               = 0x8F35
	MAX_SHADER_COMPILER_THREADS_ARB                            = 0x91B0
	MAX_SHADER_COMPILER_THREADS_KHR                            = 0x91B0
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	MAX_SPARSE_3D_TEXTURE_SIZE_ARB                             = 0x9199
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS_ARB                        = 0x919A
	MAX_SPARSE_TEXTURE_SIZE_ARB                                = 0x9198
	MAX_SUBPIXEL_PRECISION_BIAS_BITS_NV                        = 0x9349
	MAX_SUBROUTINES                                            = 0x8DE7
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	MAX_TEXTURE_BUFFER_SIZE_ARB                                = 0x8C2B
	MAX_TEXTURE_COORDS                                         = 0x8871
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	MAX_TEXTURE_MAX_ANISOTROPY                                 = 0x84FF
	MAX_TEXTURE_SIZE                                           = 0x0D33
	MAX_TEXTURE_STACK_DEPTH                                    = 0x0D39
	MAX_TEXTURE_UNITS                                          = 0x84E2
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	MAX_VARYING_FLOATS                                         = 0x8B4B
	MAX_VARYING_VECTORS                                        = 0x8DFC
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	MAX_VERTEX_ATTRIB_STRIDE                                   = 0x82E5
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	MAX_VERTEX_STREAMS                                         = 0x8E71
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	MAX_VERTEX_VARYING_COMPONENTS_ARB                          = 0x8DDE
	MAX_VIEWPORTS                                              = 0x825B
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	MAX_VIEWS_OVR                                              = 0x9631
	MAX_WIDTH                                                  = 0x827E
	MAX_WINDOW_RECTANGLES_EXT                                  = 0x8F14
	MEDIUM_FLOAT                                               = 0x8DF1
	MEDIUM_INT                                                 = 0x8DF4
	MIN                                                        = 0x8007
	MINOR_VERSION                                              = 0x821C
	MINUS_CLAMPED_NV                                           = 0x92B3
	MINUS_NV                                                   = 0x929F
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5E
	MIN_SAMPLE_SHADING_VALUE                                   = 0x8C37
	MIN_SAMPLE_SHADING_VALUE_ARB                               = 0x8C37
	MIPMAP                                                     = 0x8293
	MIRRORED_REPEAT                                            = 0x8370
	MIRRORED_REPEAT_ARB                                        = 0x8370
	MIRROR_CLAMP_TO_EDGE                                       = 0x8743
	MITER_REVERT_NV                                            = 0x90A7
	MITER_TRUNCATE_NV                                          = 0x90A8
	MIXED_DEPTH_SAMPLES_SUPPORTED_NV                           = 0x932F
	MIXED_STENCIL_SAMPLES_SUPPORTED_NV                         = 0x9330
	MODELVIEW                                                  = 0x1700
	MODELVIEW_MATRIX                                           = 0x0BA6
	MODELVIEW_STACK_DEPTH                                      = 0x0BA3
	MODULATE                                                   = 0x2100
	MOVE_TO_CONTINUES_NV                                       = 0x90B6
	MOVE_TO_NV                                                 = 0x02
	MOVE_TO_RESETS_NV                                          = 0x90B5
	MULT                                                       = 0x0103
	MULTIPLY_KHR                                               = 0x9294
	MULTIPLY_NV                                                = 0x9294
	MULTISAMPLE                                                = 0x809D
	MULTISAMPLES_NV                                            = 0x9371
	MULTISAMPLE_BIT                                            = 0x20000000
	MULTISAMPLE_COVERAGE_MODES_NV                              = 0x8E12
	MULTISAMPLE_LINE_WIDTH_GRANULARITY_ARB                     = 0x9382
	MULTISAMPLE_LINE_WIDTH_RANGE_ARB                           = 0x9381
	MULTISAMPLE_RASTERIZATION_ALLOWED_EXT                      = 0x932B
	N3F_V3F                                                    = 0x2A25
	NAMED_STRING_LENGTH_ARB                                    = 0x8DE9
	NAMED_STRING_TYPE_ARB                                      = 0x8DEA
	NAME_LENGTH                                                = 0x92F9
	NAME_STACK_DEPTH                                           = 0x0D70
	NAND                                                       = 0x150E
	NEAREST                                                    = 0x2600
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	NEGATIVE_ONE_TO_ONE                                        = 0x935E
	NEVER                                                      = 0x0200
	NICEST                                                     = 0x1102
	NONE                                                       = 0
	NOOP                                                       = 0x1505
	NOP_COMMAND_NV                                             = 0x0001
	NOR                                                        = 0x1508
	NORMALIZE                                                  = 0x0BA1
	NORMAL_ARRAY                                               = 0x8075
	NORMAL_ARRAY_ADDRESS_NV                                    = 0x8F22
	NORMAL_ARRAY_BUFFER_BINDING                                = 0x8897
	NORMAL_ARRAY_LENGTH_NV                                     = 0x8F2C
	NORMAL_ARRAY_POINTER                                       = 0x808F
	NORMAL_ARRAY_STRIDE                                        = 0x807F
	NORMAL_ARRAY_TYPE                                          = 0x807E
	NORMAL_MAP                                                 = 0x8511
	NOTEQUAL                                                   = 0x0205
	NO_ERROR                                                   = 0
	NO_RESET_NOTIFICATION                                      = 0x8261
	NO_RESET_NOTIFICATION_ARB                                  = 0x8261
	NO_RESET_NOTIFICATION_KHR                                  = 0x8261
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	NUM_EXTENSIONS                                             = 0x821D
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	NUM_SAMPLE_COUNTS                                          = 0x9380
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	NUM_SPARSE_LEVELS_ARB                                      = 0x91AA
	NUM_SPIR_V_EXTENSIONS                                      = 0x9554
	NUM_VIRTUAL_PAGE_SIZES_ARB                                 = 0x91A8
	NUM_WINDOW_RECTANGLES_EXT                                  = 0x8F15
	OBJECT_LINEAR                                              = 0x2401
	OBJECT_PLANE                                               = 0x2501
	OBJECT_TYPE                                                = 0x9112
	OFFSET                                                     = 0x92FC
	ONE                                                        = 1
	ONE_MINUS_CONSTANT_ALPHA                                   = 0x8004
	ONE_MINUS_CONSTANT_COLOR                                   = 0x8002
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	ONE_MINUS_DST_COLOR                                        = 0x0307
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	OPERAND0_ALPHA                                             = 0x8598
	OPERAND0_RGB                                               = 0x8590
	OPERAND1_ALPHA                                             = 0x8599
	OPERAND1_RGB                                               = 0x8591
	OPERAND2_ALPHA                                             = 0x859A
	OPERAND2_RGB                                               = 0x8592
	OR                                                         = 0x1507
	ORDER                                                      = 0x0A01
	OR_INVERTED                                                = 0x150D
	OR_REVERSE                                                 = 0x150B
	OUT_OF_MEMORY                                              = 0x0505
	OVERLAY_KHR                                                = 0x9296
	OVERLAY_NV                                                 = 0x9296
	PACK_ALIGNMENT                                             = 0x0D05
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	PACK_IMAGE_HEIGHT                                          = 0x806C
	PACK_LSB_FIRST                                             = 0x0D01
	PACK_ROW_LENGTH                                            = 0x0D02
	PACK_SKIP_IMAGES                                           = 0x806B
	PACK_SKIP_PIXELS                                           = 0x0D04
	PACK_SKIP_ROWS                                             = 0x0D03
	PACK_SWAP_BYTES                                            = 0x0D00
	PARAMETER_BUFFER                                           = 0x80EE
	PARAMETER_BUFFER_ARB                                       = 0x80EE
	PARAMETER_BUFFER_BINDING                                   = 0x80EF
	PARAMETER_BUFFER_BINDING_ARB                               = 0x80EF
	PASS_THROUGH_TOKEN                                         = 0x0700
	PATCHES                                                    = 0x000E
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	PATCH_VERTICES                                             = 0x8E72
	PATH_CLIENT_LENGTH_NV                                      = 0x907F
	PATH_COMMAND_COUNT_NV                                      = 0x909D
	PATH_COMPUTED_LENGTH_NV                                    = 0x90A0
	PATH_COORD_COUNT_NV                                        = 0x909E
	PATH_COVER_DEPTH_FUNC_NV                                   = 0x90BF
	PATH_DASH_ARRAY_COUNT_NV                                   = 0x909F
	PATH_DASH_CAPS_NV                                          = 0x907B
	PATH_DASH_OFFSET_NV                                        = 0x907E
	PATH_DASH_OFFSET_RESET_NV                                  = 0x90B4
	PATH_END_CAPS_NV                                           = 0x9076
	PATH_ERROR_POSITION_NV                                     = 0x90AB
	PATH_FILL_BOUNDING_BOX_NV                                  = 0x90A1
	PATH_FILL_COVER_MODE_NV                                    = 0x9082
	PATH_FILL_MASK_NV                                          = 0x9081
	PATH_FILL_MODE_NV                                          = 0x9080
	PATH_FORMAT_PS_NV                                          = 0x9071
	PATH_FORMAT_SVG_NV                                         = 0x9070
	PATH_GEN_COEFF_NV                                          = 0x90B1
	PATH_GEN_COMPONENTS_NV                                     = 0x90B3
	PATH_GEN_MODE_NV                                           = 0x90B0
	PATH_INITIAL_DASH_CAP_NV                                   = 0x907C
	PATH_INITIAL_END_CAP_NV                                    = 0x9077
	PATH_JOIN_STYLE_NV                                         = 0x9079
	PATH_MAX_MODELVIEW_STACK_DEPTH_NV                          = 0x0D36
	PATH_MAX_PROJECTION_STACK_DEPTH_NV                         = 0x0D38
	PATH_MITER_LIMIT_NV                                        = 0x907A
	PATH_MODELVIEW_MATRIX_NV                                   = 0x0BA6
	PATH_MODELVIEW_NV                                          = 0x1700
	PATH_MODELVIEW_STACK_DEPTH_NV                              = 0x0BA3
	PATH_OBJECT_BOUNDING_BOX_NV                                = 0x908A
	PATH_PROJECTION_MATRIX_NV                                  = 0x0BA7
	PATH_PROJECTION_NV                                         = 0x1701
	PATH_PROJECTION_STACK_DEPTH_NV                             = 0x0BA4
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                        = 0x90BD
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                         = 0x90BE
	PATH_STENCIL_FUNC_NV                                       = 0x90B7
	PATH_STENCIL_REF_NV                                        = 0x90B8
	PATH_STENCIL_VALUE_MASK_NV                                 = 0x90B9
	PATH_STROKE_BOUNDING_BOX_NV                                = 0x90A2
	PATH_STROKE_COVER_MODE_NV                                  = 0x9083
	PATH_STROKE_MASK_NV                                        = 0x9084
	PATH_STROKE_WIDTH_NV                                       = 0x9075
	PATH_TERMINAL_DASH_CAP_NV                                  = 0x907D
	PATH_TERMINAL_END_CAP_NV                                   = 0x9078
	PATH_TRANSPOSE_MODELVIEW_MATRIX_NV                         = 0x84E3
	PATH_TRANSPOSE_PROJECTION_MATRIX_NV                        = 0x84E4
	PERCENTAGE_AMD                                             = 0x8BC3
	PERFMON_RESULT_AMD                                         = 0x8BC6
	PERFMON_RESULT_AVAILABLE_AMD                               = 0x8BC4
	PERFMON_RESULT_SIZE_AMD                                    = 0x8BC5
	PERFQUERY_COUNTER_DATA_BOOL32_INTEL                        = 0x94FC
	PERFQUERY_COUNTER_DATA_DOUBLE_INTEL                        = 0x94FB
	PERFQUERY_COUNTER_DATA_FLOAT_INTEL                         = 0x94FA
	PERFQUERY_COUNTER_DATA_UINT32_INTEL                        = 0x94F8
	PERFQUERY_COUNTER_DATA_UINT64_INTEL                        = 0x94F9
	PERFQUERY_COUNTER_DESC_LENGTH_MAX_INTEL                    = 0x94FF
	PERFQUERY_COUNTER_DURATION_NORM_INTEL                      = 0x94F1
	PERFQUERY_COUNTER_DURATION_RAW_INTEL                       = 0x94F2
	PERFQUERY_COUNTER_EVENT_INTEL                              = 0x94F0
	PERFQUERY_COUNTER_NAME_LENGTH_MAX_INTEL                    = 0x94FE
	PERFQUERY_COUNTER_RAW_INTEL                                = 0x94F4
	PERFQUERY_COUNTER_THROUGHPUT_INTEL                         = 0x94F3
	PERFQUERY_COUNTER_TIMESTAMP_INTEL                          = 0x94F5
	PERFQUERY_DONOT_FLUSH_INTEL                                = 0x83F9
	PERFQUERY_FLUSH_INTEL                                      = 0x83FA
	PERFQUERY_GLOBAL_CONTEXT_INTEL                             = 0x00000001
	PERFQUERY_GPA_EXTENDED_COUNTERS_INTEL                      = 0x9500
	PERFQUERY_QUERY_NAME_LENGTH_MAX_INTEL                      = 0x94FD
	PERFQUERY_SINGLE_CONTEXT_INTEL                             = 0x00000000
	PERFQUERY_WAIT_INTEL                                       = 0x83FB
	PERSPECTIVE_CORRECTION_HINT                                = 0x0C50
	PINLIGHT_NV                                                = 0x92A8
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	PIXEL_MAP_A_TO_A                                           = 0x0C79
	PIXEL_MAP_A_TO_A_SIZE                                      = 0x0CB9
	PIXEL_MAP_B_TO_B                                           = 0x0C78
	PIXEL_MAP_B_TO_B_SIZE                                      = 0x0CB8
	PIXEL_MAP_G_TO_G                                           = 0x0C77
	PIXEL_MAP_G_TO_G_SIZE                                      = 0x0CB7
	PIXEL_MAP_I_TO_A                                           = 0x0C75
	PIXEL_MAP_I_TO_A_SIZE                                      = 0x0CB5
	PIXEL_MAP_I_TO_B                                           = 0x0C74
	PIXEL_MAP_I_TO_B_SIZE                                      = 0x0CB4
	PIXEL_MAP_I_TO_G                                           = 0x0C73
	PIXEL_MAP_I_TO_G_SIZE                                      = 0x0CB3
	PIXEL_MAP_I_TO_I                                           = 0x0C70
	PIXEL_MAP_I_TO_I_SIZE                                      = 0x0CB0
	PIXEL_MAP_I_TO_R                                           = 0x0C72
	PIXEL_MAP_I_TO_R_SIZE                                      = 0x0CB2
	PIXEL_MAP_R_TO_R                                           = 0x0C76
	PIXEL_MAP_R_TO_R_SIZE                                      = 0x0CB6
	PIXEL_MAP_S_TO_S                                           = 0x0C71
	PIXEL_MAP_S_TO_S_SIZE                                      = 0x0CB1
	PIXEL_MODE_BIT                                             = 0x00000020
	PIXEL_PACK_BUFFER                                          = 0x88EB
	PIXEL_PACK_BUFFER_ARB                                      = 0x88EB
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	PIXEL_PACK_BUFFER_BINDING_ARB                              = 0x88ED
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	PIXEL_UNPACK_BUFFER_ARB                                    = 0x88EC
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	PIXEL_UNPACK_BUFFER_BINDING_ARB                            = 0x88EF
	PLUS_CLAMPED_ALPHA_NV                                      = 0x92B2
	PLUS_CLAMPED_NV                                            = 0x92B1
	PLUS_DARKER_NV                                             = 0x9292
	PLUS_NV                                                    = 0x9291
	POINT                                                      = 0x1B00
	POINTS                                                     = 0x0000
	POINT_BIT                                                  = 0x00000002
	POINT_DISTANCE_ATTENUATION                                 = 0x8129
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	POINT_SIZE                                                 = 0x0B11
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	POINT_SIZE_MAX                                             = 0x8127
	POINT_SIZE_MIN                                             = 0x8126
	POINT_SIZE_RANGE                                           = 0x0B12
	POINT_SMOOTH                                               = 0x0B10
	POINT_SMOOTH_HINT                                          = 0x0C51
	POINT_SPRITE                                               = 0x8861
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	POINT_TOKEN                                                = 0x0701
	POLYGON                                                    = 0x0009
	POLYGON_BIT                                                = 0x00000008
	POLYGON_MODE                                               = 0x0B40
	POLYGON_OFFSET_CLAMP                                       = 0x8E1B
	POLYGON_OFFSET_CLAMP_EXT                                   = 0x8E1B
	POLYGON_OFFSET_COMMAND_NV                                  = 0x000E
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	POLYGON_OFFSET_FILL                                        = 0x8037
	POLYGON_OFFSET_LINE                                        = 0x2A02
	POLYGON_OFFSET_POINT                                       = 0x2A01
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	POLYGON_SMOOTH                                             = 0x0B41
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	POLYGON_STIPPLE                                            = 0x0B42
	POLYGON_STIPPLE_BIT                                        = 0x00000010
	POLYGON_TOKEN                                              = 0x0703
	POSITION                                                   = 0x1203
	PREVIOUS                                                   = 0x8578
	PRIMARY_COLOR                                              = 0x8577
	PRIMITIVES_GENERATED                                       = 0x8C87
	PRIMITIVES_SUBMITTED                                       = 0x82EF
	PRIMITIVES_SUBMITTED_ARB                                   = 0x82EF
	PRIMITIVE_BOUNDING_BOX_ARB                                 = 0x92BE
	PRIMITIVE_RESTART                                          = 0x8F9D
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	PRIMITIVE_RESTART_FOR_PATCHES_SUPPORTED                    = 0x8221
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	PROGRAM                                                    = 0x82E2
	PROGRAMMABLE_SAMPLE_LOCATION_ARB                           = 0x9341
	PROGRAMMABLE_SAMPLE_LOCATION_NV                            = 0x9341
	PROGRAMMABLE_SAMPLE_LOCATION_TABLE_SIZE_ARB                = 0x9340
	PROGRAMMABLE_SAMPLE_LOCATION_TABLE_SIZE_NV                 = 0x9340
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	PROGRAM_INPUT                                              = 0x92E3
	PROGRAM_KHR                                                = 0x82E2
	PROGRAM_MATRIX_EXT                                         = 0x8E2D
	PROGRAM_MATRIX_STACK_DEPTH_EXT                             = 0x8E2F
	PROGRAM_OBJECT_EXT                                         = 0x8B40
	PROGRAM_OUTPUT                                             = 0x92E4
	PROGRAM_PIPELINE                                           = 0x82E4
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	PROGRAM_PIPELINE_BINDING_EXT                               = 0x825A
	PROGRAM_PIPELINE_KHR                                       = 0x82E4
	PROGRAM_PIPELINE_OBJECT_EXT                                = 0x8A4F
	PROGRAM_POINT_SIZE                                         = 0x8642
	PROGRAM_POINT_SIZE_ARB                                     = 0x8642
	PROGRAM_SEPARABLE                                          = 0x8258
	PROGRAM_SEPARABLE_EXT                                      = 0x8258
	PROJECTION                                                 = 0x1701
	PROJECTION_MATRIX                                          = 0x0BA7
	PROJECTION_STACK_DEPTH                                     = 0x0BA4
	PROVOKING_VERTEX                                           = 0x8E4F
	PROXY_TEXTURE_1D                                           = 0x8063
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	PROXY_TEXTURE_2D                                           = 0x8064
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	PROXY_TEXTURE_3D                                           = 0x8070
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	PROXY_TEXTURE_CUBE_MAP_ARRAY_ARB                           = 0x900B
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	Q                                                          = 0x2003
	QUADRATIC_ATTENUATION                                      = 0x1209
	QUADRATIC_CURVE_TO_NV                                      = 0x0A
	QUADS                                                      = 0x0007
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	QUAD_STRIP                                                 = 0x0008
	QUERY                                                      = 0x82E3
	QUERY_BUFFER                                               = 0x9192
	QUERY_BUFFER_BARRIER_BIT                                   = 0x00008000
	QUERY_BUFFER_BINDING                                       = 0x9193
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	QUERY_BY_REGION_NO_WAIT_INVERTED                           = 0x8E1A
	QUERY_BY_REGION_NO_WAIT_NV                                 = 0x8E16
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	QUERY_BY_REGION_WAIT_INVERTED                              = 0x8E19
	QUERY_BY_REGION_WAIT_NV                                    = 0x8E15
	QUERY_COUNTER_BITS                                         = 0x8864
	QUERY_KHR                                                  = 0x82E3
	QUERY_NO_WAIT                                              = 0x8E14
	QUERY_NO_WAIT_INVERTED                                     = 0x8E18
	QUERY_NO_WAIT_NV                                           = 0x8E14
	QUERY_OBJECT_EXT                                           = 0x9153
	QUERY_RESULT                                               = 0x8866
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	QUERY_RESULT_NO_WAIT                                       = 0x9194
	QUERY_TARGET                                               = 0x82EA
	QUERY_WAIT                                                 = 0x8E13
	QUERY_WAIT_INVERTED                                        = 0x8E17
	QUERY_WAIT_NV                                              = 0x8E13
	R                                                          = 0x2002
	R11F_G11F_B10F                                             = 0x8C3A
	R16                                                        = 0x822A
	R16F                                                       = 0x822D
	R16I                                                       = 0x8233
	R16UI                                                      = 0x8234
	R16_SNORM                                                  = 0x8F98
	R32F                                                       = 0x822E
	R32I                                                       = 0x8235
	R32UI                                                      = 0x8236
	R3_G3_B2                                                   = 0x2A10
	R8                                                         = 0x8229
	R8I                                                        = 0x8231
	R8UI                                                       = 0x8232
	R8_SNORM                                                   = 0x8F94
	RASTERIZER_DISCARD                                         = 0x8C89
	RASTER_FIXED_SAMPLE_LOCATIONS_EXT                          = 0x932A
	RASTER_MULTISAMPLE_EXT                                     = 0x9327
	RASTER_SAMPLES_EXT                                         = 0x9328
	READ_BUFFER                                                = 0x0C02
	READ_FRAMEBUFFER                                           = 0x8CA8
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	READ_ONLY                                                  = 0x88B8
	READ_PIXELS                                                = 0x828C
	READ_PIXELS_FORMAT                                         = 0x828D
	READ_PIXELS_TYPE                                           = 0x828E
	READ_WRITE                                                 = 0x88BA
	RECT_NV                                                    = 0xF6
	RED                                                        = 0x1903
	RED_BIAS                                                   = 0x0D15
	RED_BITS                                                   = 0x0D52
	RED_INTEGER                                                = 0x8D94
	RED_NV                                                     = 0x1903
	RED_SCALE                                                  = 0x0D14
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	REFLECTION_MAP                                             = 0x8512
	RELATIVE_ARC_TO_NV                                         = 0xFF
	RELATIVE_CONIC_CURVE_TO_NV                                 = 0x1B
	RELATIVE_CUBIC_CURVE_TO_NV                                 = 0x0D
	RELATIVE_HORIZONTAL_LINE_TO_NV                             = 0x07
	RELATIVE_LARGE_CCW_ARC_TO_NV                               = 0x17
	RELATIVE_LARGE_CW_ARC_TO_NV                                = 0x19
	RELATIVE_LINE_TO_NV                                        = 0x05
	RELATIVE_MOVE_TO_NV                                        = 0x03
	RELATIVE_QUADRATIC_CURVE_TO_NV                             = 0x0B
	RELATIVE_RECT_NV                                           = 0xF7
	RELATIVE_ROUNDED_RECT2_NV                                  = 0xEB
	RELATIVE_ROUNDED_RECT4_NV                                  = 0xED
	RELATIVE_ROUNDED_RECT8_NV                                  = 0xEF
	RELATIVE_ROUNDED_RECT_NV                                   = 0xE9
	RELATIVE_SMALL_CCW_ARC_TO_NV                               = 0x13
	RELATIVE_SMALL_CW_ARC_TO_NV                                = 0x15
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                          = 0x11
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                      = 0x0F
	RELATIVE_VERTICAL_LINE_TO_NV                               = 0x09
	RENDER                                                     = 0x1C00
	RENDERBUFFER                                               = 0x8D41
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	RENDERBUFFER_BINDING                                       = 0x8CA7
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	RENDERBUFFER_COLOR_SAMPLES_NV                              = 0x8E10
	RENDERBUFFER_COVERAGE_SAMPLES_NV                           = 0x8CAB
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	RENDERBUFFER_WIDTH                                         = 0x8D42
	RENDERER                                                   = 0x1F01
	RENDER_MODE                                                = 0x0C40
	REPEAT                                                     = 0x2901
	REPLACE                                                    = 0x1E01
	RESCALE_NORMAL                                             = 0x803A
	RESET_NOTIFICATION_STRATEGY                                = 0x8256
	RESET_NOTIFICATION_STRATEGY_ARB                            = 0x8256
	RESET_NOTIFICATION_STRATEGY_KHR                            = 0x8256
	RESTART_PATH_NV                                            = 0xF0
	RETURN                                                     = 0x0102
	RG                                                         = 0x8227
	RG16                                                       = 0x822C
	RG16F                                                      = 0x822F
	RG16I                                                      = 0x8239
	RG16UI                                                     = 0x823A
	RG16_SNORM                                                 = 0x8F99
	RG32F                                                      = 0x8230
	RG32I                                                      = 0x823B
	RG32UI                                                     = 0x823C
	RG8                                                        = 0x822B
	RG8I                                                       = 0x8237
	RG8UI                                                      = 0x8238
	RG8_SNORM                                                  = 0x8F95
	RGB                                                        = 0x1907
	RGB10                                                      = 0x8052
	RGB10_A2                                                   = 0x8059
	RGB10_A2UI                                                 = 0x906F
	RGB12                                                      = 0x8053
	RGB16                                                      = 0x8054
	RGB16F                                                     = 0x881B
	RGB16I                                                     = 0x8D89
	RGB16UI                                                    = 0x8D77
	RGB16_SNORM                                                = 0x8F9A
	RGB32F                                                     = 0x8815
	RGB32I                                                     = 0x8D83
	RGB32UI                                                    = 0x8D71
	RGB4                                                       = 0x804F
	RGB5                                                       = 0x8050
	RGB565                                                     = 0x8D62
	RGB5_A1                                                    = 0x8057
	RGB8                                                       = 0x8051
	RGB8I                                                      = 0x8D8F
	RGB8UI                                                     = 0x8D7D
	RGB8_SNORM                                                 = 0x8F96
	RGB9_E5                                                    = 0x8C3D
	RGBA                                                       = 0x1908
	RGBA12                                                     = 0x805A
	RGBA16                                                     = 0x805B
	RGBA16F                                                    = 0x881A
	RGBA16I                                                    = 0x8D88
	RGBA16UI                                                   = 0x8D76
	RGBA16_SNORM                                               = 0x8F9B
	RGBA2                                                      = 0x8055
	RGBA32F                                                    = 0x8814
	RGBA32I                                                    = 0x8D82
	RGBA32UI                                                   = 0x8D70
	RGBA4                                                      = 0x8056
	RGBA8                                                      = 0x8058
	RGBA8I                                                     = 0x8D8E
	RGBA8UI                                                    = 0x8D7C
	RGBA8_SNORM                                                = 0x8F97
	RGBA_INTEGER                                               = 0x8D99
	RGBA_MODE                                                  = 0x0C31
	RGB_422_APPLE                                              = 0x8A1F
	RGB_INTEGER                                                = 0x8D98
	RGB_RAW_422_APPLE                                          = 0x8A51
	RGB_SCALE                                                  = 0x8573
	RG_INTEGER                                                 = 0x8228
	RIGHT                                                      = 0x0407
	ROUNDED_RECT2_NV                                           = 0xEA
	ROUNDED_RECT4_NV                                           = 0xEC
	ROUNDED_RECT8_NV                                           = 0xEE
	ROUNDED_RECT_NV                                            = 0xE8
	ROUND_NV                                                   = 0x90A4
	S                                                          = 0x2000
	SAMPLER                                                    = 0x82E6
	SAMPLER_1D                                                 = 0x8B5D
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	SAMPLER_1D_SHADOW                                          = 0x8B61
	SAMPLER_2D                                                 = 0x8B5E
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	SAMPLER_2D_RECT                                            = 0x8B63
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	SAMPLER_2D_SHADOW                                          = 0x8B62
	SAMPLER_3D                                                 = 0x8B5F
	SAMPLER_BINDING                                            = 0x8919
	SAMPLER_BUFFER                                             = 0x8DC2
	SAMPLER_CUBE                                               = 0x8B60
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	SAMPLER_CUBE_MAP_ARRAY_ARB                                 = 0x900C
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	SAMPLER_CUBE_MAP_ARRAY_SHADOW_ARB                          = 0x900D
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	SAMPLER_KHR                                                = 0x82E6
	SAMPLES                                                    = 0x80A9
	SAMPLES_PASSED                                             = 0x8914
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	SAMPLE_BUFFERS                                             = 0x80A8
	SAMPLE_COVERAGE                                            = 0x80A0
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	SAMPLE_LOCATION_ARB                                        = 0x8E50
	SAMPLE_LOCATION_NV                                         = 0x8E50
	SAMPLE_LOCATION_PIXEL_GRID_HEIGHT_ARB                      = 0x933F
	SAMPLE_LOCATION_PIXEL_GRID_HEIGHT_NV                       = 0x933F
	SAMPLE_LOCATION_PIXEL_GRID_WIDTH_ARB                       = 0x933E
	SAMPLE_LOCATION_PIXEL_GRID_WIDTH_NV                        = 0x933E
	SAMPLE_LOCATION_SUBPIXEL_BITS_ARB                          = 0x933D
	SAMPLE_LOCATION_SUBPIXEL_BITS_NV                           = 0x933D
	SAMPLE_MASK                                                = 0x8E51
	SAMPLE_MASK_VALUE                                          = 0x8E52
	SAMPLE_POSITION                                            = 0x8E50
	SAMPLE_SHADING                                             = 0x8C36
	SAMPLE_SHADING_ARB                                         = 0x8C36
	SCISSOR_BIT                                                = 0x00080000
	SCISSOR_BOX                                                = 0x0C10
	SCISSOR_COMMAND_NV                                         = 0x0011
	SCISSOR_TEST                                               = 0x0C11
	SCREEN_KHR                                                 = 0x9295
	SCREEN_NV                                                  = 0x9295
	SECONDARY_COLOR_ARRAY                                      = 0x845E
	SECONDARY_COLOR_ARRAY_ADDRESS_NV                           = 0x8F27
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING                       = 0x889C
	SECONDARY_COLOR_ARRAY_LENGTH_NV                            = 0x8F31
	SECONDARY_COLOR_ARRAY_POINTER                              = 0x845D
	SECONDARY_COLOR_ARRAY_SIZE                                 = 0x845A
	SECONDARY_COLOR_ARRAY_STRIDE                               = 0x845C
	SECONDARY_COLOR_ARRAY_TYPE                                 = 0x845B
	SELECT                                                     = 0x1C02
	SELECTION_BUFFER_POINTER                                   = 0x0DF3
	SELECTION_BUFFER_SIZE                                      = 0x0DF4
	SEPARATE_ATTRIBS                                           = 0x8C8D
	SEPARATE_SPECULAR_COLOR                                    = 0x81FA
	SET                                                        = 0x150F
	SHADER                                                     = 0x82E1
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	SHADER_BINARY_FORMAT_SPIR_V                                = 0x9551
	SHADER_BINARY_FORMAT_SPIR_V_ARB                            = 0x9551
	SHADER_COMPILER                                            = 0x8DFA
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV                        = 0x00000010
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	SHADER_IMAGE_LOAD                                          = 0x82A4
	SHADER_IMAGE_STORE                                         = 0x82A5
	SHADER_INCLUDE_ARB                                         = 0x8DAE
	SHADER_KHR                                                 = 0x82E1
	SHADER_OBJECT_EXT                                          = 0x8B48
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	SHADER_STORAGE_BARRIER_BIT                                 = 0x00002000
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	SHADER_TYPE                                                = 0x8B4F
	SHADE_MODEL                                                = 0x0B54
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	SHARED_EDGE_NV                                             = 0xC0
	SHININESS                                                  = 0x1601
	SHORT                                                      = 0x1402
	SIGNALED                                                   = 0x9119
	SIGNED_NORMALIZED                                          = 0x8F9C
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	SINGLE_COLOR                                               = 0x81F9
	SKIP_DECODE_EXT                                            = 0x8A4A
	SKIP_MISSING_GLYPH_NV                                      = 0x90A9
	SLUMINANCE                                                 = 0x8C46
	SLUMINANCE8                                                = 0x8C47
	SLUMINANCE8_ALPHA8                                         = 0x8C45
	SLUMINANCE_ALPHA                                           = 0x8C44
	SMALL_CCW_ARC_TO_NV                                        = 0x12
	SMALL_CW_ARC_TO_NV                                         = 0x14
	SMOOTH                                                     = 0x1D01
	SMOOTH_CUBIC_CURVE_TO_NV                                   = 0x10
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	SMOOTH_QUADRATIC_CURVE_TO_NV                               = 0x0E
	SM_COUNT_NV                                                = 0x933B
	SOFTLIGHT_KHR                                              = 0x929C
	SOFTLIGHT_NV                                               = 0x929C
	SOURCE0_ALPHA                                              = 0x8588
	SOURCE0_RGB                                                = 0x8580
	SOURCE1_ALPHA                                              = 0x8589
	SOURCE1_RGB                                                = 0x8581
	SOURCE2_ALPHA                                              = 0x858A
	SOURCE2_RGB                                                = 0x8582
	SPARSE_BUFFER_PAGE_SIZE_ARB                                = 0x82F8
	SPARSE_STORAGE_BIT_ARB                                     = 0x0400
	SPARSE_TEXTURE_FULL_ARRAY_CUBE_MIPMAPS_ARB                 = 0x91A9
	SPECULAR                                                   = 0x1202
	SPHERE_MAP                                                 = 0x2402
	SPIR_V_BINARY                                              = 0x9552
	SPIR_V_BINARY_ARB                                          = 0x9552
	SPIR_V_EXTENSIONS                                          = 0x9553
	SPOT_CUTOFF                                                = 0x1206
	SPOT_DIRECTION                                             = 0x1204
	SPOT_EXPONENT                                              = 0x1205
	SQUARE_NV                                                  = 0x90A3
	SRC0_ALPHA                                                 = 0x8588
	SRC0_RGB                                                   = 0x8580
	SRC1_ALPHA                                                 = 0x8589
	SRC1_COLOR                                                 = 0x88F9
	SRC1_RGB                                                   = 0x8581
	SRC2_ALPHA                                                 = 0x858A
	SRC2_RGB                                                   = 0x8582
	SRC_ALPHA                                                  = 0x0302
	SRC_ALPHA_SATURATE                                         = 0x0308
	SRC_ATOP_NV                                                = 0x928E
	SRC_COLOR                                                  = 0x0300
	SRC_IN_NV                                                  = 0x928A
	SRC_NV                                                     = 0x9286
	SRC_OUT_NV                                                 = 0x928C
	SRC_OVER_NV                                                = 0x9288
	SRGB                                                       = 0x8C40
	SRGB8                                                      = 0x8C41
	SRGB8_ALPHA8                                               = 0x8C43
	SRGB_ALPHA                                                 = 0x8C42
	SRGB_DECODE_ARB                                            = 0x8299
	SRGB_READ                                                  = 0x8297
	SRGB_WRITE                                                 = 0x8298
	STACK_OVERFLOW                                             = 0x0503
	STACK_OVERFLOW_KHR                                         = 0x0503
	STACK_UNDERFLOW                                            = 0x0504
	STACK_UNDERFLOW_KHR                                        = 0x0504
	STANDARD_FONT_FORMAT_NV                                    = 0x936C
	STANDARD_FONT_NAME_NV                                      = 0x9072
	STATIC_COPY                                                = 0x88E6
	STATIC_DRAW                                                = 0x88E4
	STATIC_READ                                                = 0x88E5
	STENCIL                                                    = 0x1802
	STENCIL_ATTACHMENT                                         = 0x8D20
	STENCIL_BACK_FAIL                                          = 0x8801
	STENCIL_BACK_FUNC                                          = 0x8800
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	STENCIL_BACK_REF                                           = 0x8CA3
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	STENCIL_BITS                                               = 0x0D57
	STENCIL_BUFFER_BIT                                         = 0x00000400
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	STENCIL_COMPONENTS                                         = 0x8285
	STENCIL_FAIL                                               = 0x0B94
	STENCIL_FUNC                                               = 0x0B92
	STENCIL_INDEX                                              = 0x1901
	STENCIL_INDEX1                                             = 0x8D46
	STENCIL_INDEX16                                            = 0x8D49
	STENCIL_INDEX4                                             = 0x8D47
	STENCIL_INDEX8                                             = 0x8D48
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	STENCIL_REF                                                = 0x0B97
	STENCIL_REF_COMMAND_NV                                     = 0x000C
	STENCIL_RENDERABLE                                         = 0x8288
	STENCIL_SAMPLES_NV                                         = 0x932E
	STENCIL_TEST                                               = 0x0B90
	STENCIL_VALUE_MASK                                         = 0x0B93
	STENCIL_WRITEMASK                                          = 0x0B98
	STEREO                                                     = 0x0C33
	STREAM_COPY                                                = 0x88E2
	STREAM_DRAW                                                = 0x88E0
	STREAM_READ                                                = 0x88E1
	SUBPIXEL_BITS                                              = 0x0D50
	SUBPIXEL_PRECISION_BIAS_X_BITS_NV                          = 0x9347
	SUBPIXEL_PRECISION_BIAS_Y_BITS_NV                          = 0x9348
	SUBTRACT                                                   = 0x84E7
	SUPERSAMPLE_SCALE_X_NV                                     = 0x9372
	SUPERSAMPLE_SCALE_Y_NV                                     = 0x9373
	SYNC_CL_EVENT_ARB                                          = 0x8240
	SYNC_CL_EVENT_COMPLETE_ARB                                 = 0x8241
	SYNC_CONDITION                                             = 0x9113
	SYNC_FENCE                                                 = 0x9116
	SYNC_FLAGS                                                 = 0x9115
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	SYNC_STATUS                                                = 0x9114
	SYSTEM_FONT_NAME_NV                                        = 0x9073
	T                                                          = 0x2001
	T2F_C3F_V3F                                                = 0x2A2A
	T2F_C4F_N3F_V3F                                            = 0x2A2C
	T2F_C4UB_V3F                                               = 0x2A29
	T2F_N3F_V3F                                                = 0x2A2B
	T2F_V3F                                                    = 0x2A27
	T4F_C4F_N3F_V4F                                            = 0x2A2D
	T4F_V4F                                                    = 0x2A28
	TERMINATE_SEQUENCE_COMMAND_NV                              = 0x0000
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	TESS_CONTROL_SHADER                                        = 0x8E88
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	TESS_CONTROL_SHADER_PATCHES                                = 0x82F1
	TESS_CONTROL_SHADER_PATCHES_ARB                            = 0x82F1
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	TESS_CONTROL_TEXTURE                                       = 0x829C
	TESS_EVALUATION_SHADER                                     = 0x8E87
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	TESS_EVALUATION_SHADER_INVOCATIONS                         = 0x82F2
	TESS_EVALUATION_SHADER_INVOCATIONS_ARB                     = 0x82F2
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	TESS_GEN_MODE                                              = 0x8E76
	TESS_GEN_POINT_MODE                                        = 0x8E79
	TESS_GEN_SPACING                                           = 0x8E77
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	TEXTURE                                                    = 0x1702
	TEXTURE0                                                   = 0x84C0
	TEXTURE1                                                   = 0x84C1
	TEXTURE10                                                  = 0x84CA
	TEXTURE11                                                  = 0x84CB
	TEXTURE12                                                  = 0x84CC
	TEXTURE13                                                  = 0x84CD
	TEXTURE14                                                  = 0x84CE
	TEXTURE15                                                  = 0x84CF
	TEXTURE16                                                  = 0x84D0
	TEXTURE17                                                  = 0x84D1
	TEXTURE18                                                  = 0x84D2
	TEXTURE19                                                  = 0x84D3
	TEXTURE2                                                   = 0x84C2
	TEXTURE20                                                  = 0x84D4
	TEXTURE21                                                  = 0x84D5
	TEXTURE22                                                  = 0x84D6
	TEXTURE23                                                  = 0x84D7
	TEXTURE24                                                  = 0x84D8
	TEXTURE25                                                  = 0x84D9
	TEXTURE26                                                  = 0x84DA
	TEXTURE27                                                  = 0x84DB
	TEXTURE28                                                  = 0x84DC
	TEXTURE29                                                  = 0x84DD
	TEXTURE3                                                   = 0x84C3
	TEXTURE30                                                  = 0x84DE
	TEXTURE31                                                  = 0x84DF
	TEXTURE4                                                   = 0x84C4
	TEXTURE5                                                   = 0x84C5
	TEXTURE6                                                   = 0x84C6
	TEXTURE7                                                   = 0x84C7
	TEXTURE8                                                   = 0x84C8
	TEXTURE9                                                   = 0x84C9
	TEXTURE_1D                                                 = 0x0DE0
	TEXTURE_1D_ARRAY                                           = 0x8C18
	TEXTURE_2D                                                 = 0x0DE1
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	TEXTURE_3D                                                 = 0x806F
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	TEXTURE_BASE_LEVEL                                         = 0x813C
	TEXTURE_BINDING_1D                                         = 0x8068
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	TEXTURE_BINDING_2D                                         = 0x8069
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	TEXTURE_BINDING_3D                                         = 0x806A
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	TEXTURE_BINDING_BUFFER_ARB                                 = 0x8C2C
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	TEXTURE_BINDING_CUBE_MAP_ARRAY_ARB                         = 0x900A
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	TEXTURE_BIT                                                = 0x00040000
	TEXTURE_BLUE_SIZE                                          = 0x805E
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	TEXTURE_BORDER                                             = 0x1005
	TEXTURE_BORDER_COLOR                                       = 0x1004
	TEXTURE_BUFFER                                             = 0x8C2A
	TEXTURE_BUFFER_ARB                                         = 0x8C2A
	TEXTURE_BUFFER_BINDING                                     = 0x8C2A
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	TEXTURE_BUFFER_DATA_STORE_BINDING_ARB                      = 0x8C2D
	TEXTURE_BUFFER_FORMAT_ARB                                  = 0x8C2E
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	TEXTURE_COMPARE_MODE                                       = 0x884C
	TEXTURE_COMPONENTS                                         = 0x1003
	TEXTURE_COMPRESSED                                         = 0x86A1
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	TEXTURE_COORD_ARRAY                                        = 0x8078
	TEXTURE_COORD_ARRAY_ADDRESS_NV                             = 0x8F25
	TEXTURE_COORD_ARRAY_BUFFER_BINDING                         = 0x889A
	TEXTURE_COORD_ARRAY_LENGTH_NV                              = 0x8F2F
	TEXTURE_COORD_ARRAY_POINTER                                = 0x8092
	TEXTURE_COORD_ARRAY_SIZE                                   = 0x8088
	TEXTURE_COORD_ARRAY_STRIDE                                 = 0x808A
	TEXTURE_COORD_ARRAY_TYPE                                   = 0x8089
	TEXTURE_CUBE_MAP                                           = 0x8513
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	TEXTURE_CUBE_MAP_ARRAY_ARB                                 = 0x9009
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	TEXTURE_DEPTH                                              = 0x8071
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	TEXTURE_ENV                                                = 0x2300
	TEXTURE_ENV_COLOR                                          = 0x2201
	TEXTURE_ENV_MODE                                           = 0x2200
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	TEXTURE_FILTER_CONTROL                                     = 0x8500
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	TEXTURE_GATHER                                             = 0x82A2
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	TEXTURE_GEN_MODE                                           = 0x2500
	TEXTURE_GEN_Q                                              = 0x0C63
	TEXTURE_GEN_R                                              = 0x0C62
	TEXTURE_GEN_S                                              = 0x0C60
	TEXTURE_GEN_T                                              = 0x0C61
	TEXTURE_GREEN_SIZE                                         = 0x805D
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	TEXTURE_HEIGHT                                             = 0x1001
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	TEXTURE_INTENSITY_SIZE                                     = 0x8061
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	TEXTURE_LOD_BIAS                                           = 0x8501
	TEXTURE_LUMINANCE_SIZE                                     = 0x8060
	TEXTURE_MAG_FILTER                                         = 0x2800
	TEXTURE_MATRIX                                             = 0x0BA8
	TEXTURE_MAX_ANISOTROPY                                     = 0x84FE
	TEXTURE_MAX_LEVEL                                          = 0x813D
	TEXTURE_MAX_LOD                                            = 0x813B
	TEXTURE_MIN_FILTER                                         = 0x2801
	TEXTURE_MIN_LOD                                            = 0x813A
	TEXTURE_PRIORITY                                           = 0x8066
	TEXTURE_RECTANGLE                                          = 0x84F5
	TEXTURE_REDUCTION_MODE_ARB                                 = 0x9366
	TEXTURE_REDUCTION_MODE_EXT                                 = 0x9366
	TEXTURE_RED_SIZE                                           = 0x805C
	TEXTURE_RED_TYPE                                           = 0x8C10
	TEXTURE_RESIDENT                                           = 0x8067
	TEXTURE_SAMPLES                                            = 0x9106
	TEXTURE_SHADOW                                             = 0x82A1
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	TEXTURE_SPARSE_ARB                                         = 0x91A6
	TEXTURE_SRGB_DECODE_EXT                                    = 0x8A48
	TEXTURE_STACK_DEPTH                                        = 0x0BA5
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	TEXTURE_TARGET                                             = 0x1006
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	TEXTURE_VIEW                                               = 0x82B5
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	TEXTURE_WIDTH                                              = 0x1000
	TEXTURE_WRAP_R                                             = 0x8072
	TEXTURE_WRAP_S                                             = 0x2802
	TEXTURE_WRAP_T                                             = 0x2803
	TIMEOUT_EXPIRED                                            = 0x911B
	TIMEOUT_IGNORED                                            = 0xFFFFFFFFFFFFFFFF
	TIMESTAMP                                                  = 0x8E28
	TIME_ELAPSED                                               = 0x88BF
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	TRANSFORM_BIT                                              = 0x00001000
	TRANSFORM_FEEDBACK                                         = 0x8E22
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = 0x8E24
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	TRANSFORM_FEEDBACK_BUFFER_INDEX                            = 0x934B
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = 0x8E23
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	TRANSFORM_FEEDBACK_BUFFER_STRIDE                           = 0x934C
	TRANSFORM_FEEDBACK_OVERFLOW                                = 0x82EC
	TRANSFORM_FEEDBACK_OVERFLOW_ARB                            = 0x82EC
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	TRANSFORM_FEEDBACK_STREAM_OVERFLOW                         = 0x82ED
	TRANSFORM_FEEDBACK_STREAM_OVERFLOW_ARB                     = 0x82ED
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	TRANSLATE_2D_NV                                            = 0x9090
	TRANSLATE_3D_NV                                            = 0x9091
	TRANSLATE_X_NV                                             = 0x908E
	TRANSLATE_Y_NV                                             = 0x908F
	TRANSPOSE_AFFINE_2D_NV                                     = 0x9096
	TRANSPOSE_AFFINE_3D_NV                                     = 0x9098
	TRANSPOSE_COLOR_MATRIX                                     = 0x84E6
	TRANSPOSE_MODELVIEW_MATRIX                                 = 0x84E3
	TRANSPOSE_PROGRAM_MATRIX_EXT                               = 0x8E2E
	TRANSPOSE_PROJECTION_MATRIX                                = 0x84E4
	TRANSPOSE_TEXTURE_MATRIX                                   = 0x84E5
	TRIANGLES                                                  = 0x0004
	TRIANGLES_ADJACENCY                                        = 0x000C
	TRIANGLES_ADJACENCY_ARB                                    = 0x000C
	TRIANGLE_FAN                                               = 0x0006
	TRIANGLE_STRIP                                             = 0x0005
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	TRIANGLE_STRIP_ADJACENCY_ARB                               = 0x000D
	TRIANGULAR_NV                                              = 0x90A5
	TRUE                                                       = 1
	TYPE                                                       = 0x92FA
	UNCORRELATED_NV                                            = 0x9282
	UNDEFINED_VERTEX                                           = 0x8260
	UNIFORM                                                    = 0x92E1
	UNIFORM_ADDRESS_COMMAND_NV                                 = 0x000A
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	UNIFORM_BLOCK                                              = 0x92E2
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                = 0x8A45
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	UNIFORM_BUFFER                                             = 0x8A11
	UNIFORM_BUFFER_ADDRESS_NV                                  = 0x936F
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	UNIFORM_BUFFER_LENGTH_NV                                   = 0x9370
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	UNIFORM_BUFFER_START                                       = 0x8A29
	UNIFORM_BUFFER_UNIFIED_NV                                  = 0x936E
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	UNIFORM_OFFSET                                             = 0x8A3B
	UNIFORM_SIZE                                               = 0x8A38
	UNIFORM_TYPE                                               = 0x8A37
	UNKNOWN_CONTEXT_RESET                                      = 0x8255
	UNKNOWN_CONTEXT_RESET_ARB                                  = 0x8255
	UNKNOWN_CONTEXT_RESET_KHR                                  = 0x8255
	UNPACK_ALIGNMENT                                           = 0x0CF5
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	UNPACK_LSB_FIRST                                           = 0x0CF1
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	UNPACK_SKIP_IMAGES                                         = 0x806D
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	UNSIGNALED                                                 = 0x9118
	UNSIGNED_BYTE                                              = 0x1401
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	UNSIGNED_INT                                               = 0x1405
	UNSIGNED_INT16_NV                                          = 0x8FF0
	UNSIGNED_INT16_VEC2_NV                                     = 0x8FF1
	UNSIGNED_INT16_VEC3_NV                                     = 0x8FF2
	UNSIGNED_INT16_VEC4_NV                                     = 0x8FF3
	UNSIGNED_INT64_AMD                                         = 0x8BC2
	UNSIGNED_INT64_ARB                                         = 0x140F
	UNSIGNED_INT64_NV                                          = 0x140F
	UNSIGNED_INT64_VEC2_ARB                                    = 0x8FF5
	UNSIGNED_INT64_VEC2_NV                                     = 0x8FF5
	UNSIGNED_INT64_VEC3_ARB                                    = 0x8FF6
	UNSIGNED_INT64_VEC3_NV                                     = 0x8FF6
	UNSIGNED_INT64_VEC4_ARB                                    = 0x8FF7
	UNSIGNED_INT64_VEC4_NV                                     = 0x8FF7
	UNSIGNED_INT8_NV                                           = 0x8FEC
	UNSIGNED_INT8_VEC2_NV                                      = 0x8FED
	UNSIGNED_INT8_VEC3_NV                                      = 0x8FEE
	UNSIGNED_INT8_VEC4_NV                                      = 0x8FEF
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	UNSIGNED_INT_24_8                                          = 0x84FA
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY_ARB                    = 0x900F
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	UNSIGNED_NORMALIZED                                        = 0x8C17
	UNSIGNED_SHORT                                             = 0x1403
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	UNSIGNED_SHORT_8_8_APPLE                                   = 0x85BA
	UNSIGNED_SHORT_8_8_REV_APPLE                               = 0x85BB
	UPPER_LEFT                                                 = 0x8CA2
	USE_MISSING_GLYPH_NV                                       = 0x90AA
	UTF16_NV                                                   = 0x909B
	UTF8_NV                                                    = 0x909A
	V2F                                                        = 0x2A20
	V3F                                                        = 0x2A21
	VALIDATE_STATUS                                            = 0x8B83
	VENDOR                                                     = 0x1F00
	VERSION                                                    = 0x1F02
	VERTEX_ARRAY                                               = 0x8074
	VERTEX_ARRAY_ADDRESS_NV                                    = 0x8F21
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	VERTEX_ARRAY_BUFFER_BINDING                                = 0x8896
	VERTEX_ARRAY_KHR                                           = 0x8074
	VERTEX_ARRAY_LENGTH_NV                                     = 0x8F2B
	VERTEX_ARRAY_OBJECT_EXT                                    = 0x9154
	VERTEX_ARRAY_POINTER                                       = 0x808E
	VERTEX_ARRAY_SIZE                                          = 0x807A
	VERTEX_ARRAY_STRIDE                                        = 0x807C
	VERTEX_ARRAY_TYPE                                          = 0x807B
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV                             = 0x8F20
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	VERTEX_ATTRIB_ARRAY_DIVISOR_ARB                            = 0x88FE
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	VERTEX_ATTRIB_ARRAY_LENGTH_NV                              = 0x8F2A
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV                             = 0x8F1E
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	VERTEX_BINDING_BUFFER                                      = 0x8F4F
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	VERTEX_PROGRAM_TWO_SIDE                                    = 0x8643
	VERTEX_SHADER                                              = 0x8B31
	VERTEX_SHADER_BIT                                          = 0x00000001
	VERTEX_SHADER_BIT_EXT                                      = 0x00000001
	VERTEX_SHADER_INVOCATIONS                                  = 0x82F0
	VERTEX_SHADER_INVOCATIONS_ARB                              = 0x82F0
	VERTEX_SUBROUTINE                                          = 0x92E8
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	VERTEX_TEXTURE                                             = 0x829B
	VERTICAL_LINE_TO_NV                                        = 0x08
	VERTICES_SUBMITTED                                         = 0x82EE
	VERTICES_SUBMITTED_ARB                                     = 0x82EE
	VIEWPORT                                                   = 0x0BA2
	VIEWPORT_BIT                                               = 0x00000800
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	VIEWPORT_COMMAND_NV                                        = 0x0010
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	VIEWPORT_POSITION_W_SCALE_NV                               = 0x937C
	VIEWPORT_POSITION_W_SCALE_X_COEFF_NV                       = 0x937D
	VIEWPORT_POSITION_W_SCALE_Y_COEFF_NV                       = 0x937E
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	VIEWPORT_SWIZZLE_NEGATIVE_W_NV                             = 0x9357
	VIEWPORT_SWIZZLE_NEGATIVE_X_NV                             = 0x9351
	VIEWPORT_SWIZZLE_NEGATIVE_Y_NV                             = 0x9353
	VIEWPORT_SWIZZLE_NEGATIVE_Z_NV                             = 0x9355
	VIEWPORT_SWIZZLE_POSITIVE_W_NV                             = 0x9356
	VIEWPORT_SWIZZLE_POSITIVE_X_NV                             = 0x9350
	VIEWPORT_SWIZZLE_POSITIVE_Y_NV                             = 0x9352
	VIEWPORT_SWIZZLE_POSITIVE_Z_NV                             = 0x9354
	VIEWPORT_SWIZZLE_W_NV                                      = 0x935B
	VIEWPORT_SWIZZLE_X_NV                                      = 0x9358
	VIEWPORT_SWIZZLE_Y_NV                                      = 0x9359
	VIEWPORT_SWIZZLE_Z_NV                                      = 0x935A
	VIEW_CLASS_128_BITS                                        = 0x82C4
	VIEW_CLASS_16_BITS                                         = 0x82CA
	VIEW_CLASS_24_BITS                                         = 0x82C9
	VIEW_CLASS_32_BITS                                         = 0x82C8
	VIEW_CLASS_48_BITS                                         = 0x82C7
	VIEW_CLASS_64_BITS                                         = 0x82C6
	VIEW_CLASS_8_BITS                                          = 0x82CB
	VIEW_CLASS_96_BITS                                         = 0x82C5
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	VIRTUAL_PAGE_SIZE_INDEX_ARB                                = 0x91A7
	VIRTUAL_PAGE_SIZE_X_ARB                                    = 0x9195
	VIRTUAL_PAGE_SIZE_Y_ARB                                    = 0x9196
	VIRTUAL_PAGE_SIZE_Z_ARB                                    = 0x9197
	VIVIDLIGHT_NV                                              = 0x92A6
	WAIT_FAILED                                                = 0x911D
	WARPS_PER_SM_NV                                            = 0x933A
	WARP_SIZE_NV                                               = 0x9339
	WEIGHTED_AVERAGE_ARB                                       = 0x9367
	WEIGHTED_AVERAGE_EXT                                       = 0x9367
	WEIGHT_ARRAY_BUFFER_BINDING                                = 0x889E
	WINDOW_RECTANGLE_EXT                                       = 0x8F12
	WINDOW_RECTANGLE_MODE_EXT                                  = 0x8F13
	WRITE_ONLY                                                 = 0x88B9
	XOR                                                        = 0x1506
	XOR_NV                                                     = 0x1506
	ZERO                                                       = 0
	ZERO_TO_ONE                                                = 0x935F
	ZOOM_X                                                     = 0x0D16
	ZOOM_Y                                                     = 0x0D17
)

var Accum func(op uint32, value float32)
var ActiveProgramEXT func(program uint32)

var ActiveShaderProgram func(pipeline uint32, program uint32)
var ActiveShaderProgramEXT func(pipeline uint32, program uint32)

var ActiveTexture func(texture uint32)

var AlphaFunc func(xfunc uint32, ref float32)
var ApplyFramebufferAttachmentCMAAINTEL func()

var AreTexturesResident func(n int32, textures *uint32, residences *bool) bool

var ArrayElement func(i int32)

var AttachShader func(program uint32, shader uint32)

var Begin func(mode uint32)

var BeginConditionalRender func(id uint32, mode uint32)
var BeginConditionalRenderNV func(id uint32, mode uint32)
var BeginPerfMonitorAMD func(monitor uint32)
var BeginPerfQueryINTEL func(queryHandle uint32)

var BeginQuery func(target uint32, id uint32)
var BeginQueryIndexed func(target uint32, index uint32, id uint32)

var BeginTransformFeedback func(primitiveMode uint32)

var BindAttribLocation func(program uint32, index uint32, name *uint8)

var BindBuffer func(target uint32, buffer uint32)

var BindBufferBase func(target uint32, index uint32, buffer uint32)

var BindBufferRange func(target uint32, index uint32, buffer uint32, offset int, size int)

var BindBuffersBase func(target uint32, first uint32, count int32, buffers *uint32)

var BindBuffersRange func(target uint32, first uint32, count int32, buffers *uint32, offsets *int, sizes *int)

var BindFragDataLocation func(program uint32, color uint32, name *uint8)

var BindFragDataLocationIndexed func(program uint32, colorNumber uint32, index uint32, name *uint8)

var BindFramebuffer func(target uint32, framebuffer uint32)

var BindImageTexture func(unit uint32, texture uint32, level int32, layered bool, layer int32, access uint32, format uint32)

var BindImageTextures func(first uint32, count int32, textures *uint32)
var BindMultiTextureEXT func(texunit uint32, target uint32, texture uint32)

var BindProgramPipeline func(pipeline uint32)
var BindProgramPipelineEXT func(pipeline uint32)

var BindRenderbuffer func(target uint32, renderbuffer uint32)

var BindSampler func(unit uint32, sampler uint32)

var BindSamplers func(first uint32, count int32, samplers *uint32)

var BindTexture func(target uint32, texture uint32)

var BindTextureUnit func(unit uint32, texture uint32)

var BindTextures func(first uint32, count int32, textures *uint32)

var BindTransformFeedback func(target uint32, id uint32)

var BindVertexArray func(array uint32)

var BindVertexBuffer func(bindingindex uint32, buffer uint32, offset int, stride int32)

var BindVertexBuffers func(first uint32, count int32, buffers *uint32, offsets *int, strides *int32)

var Bitmap func(width int32, height int32, xorig float32, yorig float32, xmove float32, ymove float32, bitmap *uint8)
var BlendBarrierKHR func()
var BlendBarrierNV func()

var BlendColor func(red float32, green float32, blue float32, alpha float32)

var BlendEquation func(mode uint32)

var BlendEquationSeparate func(modeRGB uint32, modeAlpha uint32)
var BlendEquationSeparatei func(buf uint32, modeRGB uint32, modeAlpha uint32)
var BlendEquationSeparateiARB func(buf uint32, modeRGB uint32, modeAlpha uint32)
var BlendEquationi func(buf uint32, mode uint32)
var BlendEquationiARB func(buf uint32, mode uint32)

var BlendFunc func(sfactor uint32, dfactor uint32)

var BlendFuncSeparate func(sfactorRGB uint32, dfactorRGB uint32, sfactorAlpha uint32, dfactorAlpha uint32)
var BlendFuncSeparatei func(buf uint32, srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32)
var BlendFuncSeparateiARB func(buf uint32, srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32)
var BlendFunci func(buf uint32, src uint32, dst uint32)
var BlendFunciARB func(buf uint32, src uint32, dst uint32)
var BlendParameteriNV func(pname uint32, value int32)

var BlitFramebuffer func(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32)

var BlitNamedFramebuffer func(readFramebuffer uint32, drawFramebuffer uint32, srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32)
var BufferAddressRangeNV func(pname uint32, index uint32, address uint64, length int)

var BufferData func(target uint32, size int, data unsafe.Pointer, usage uint32)
var BufferPageCommitmentARB func(target uint32, offset int, size int, commit bool)

var BufferStorage func(target uint32, size int, data unsafe.Pointer, flags uint32)

var BufferSubData func(target uint32, offset int, size int, data unsafe.Pointer)
var CallCommandListNV func(list uint32)

var CallList func(list uint32)

var CallLists func(n int32, xtype uint32, lists unsafe.Pointer)

var CheckFramebufferStatus func(target uint32) uint32

var CheckNamedFramebufferStatus func(framebuffer uint32, target uint32) uint32
var CheckNamedFramebufferStatusEXT func(framebuffer uint32, target uint32) uint32

var ClampColor func(target uint32, clamp uint32)

var Clear func(mask uint32)

var ClearAccum func(red float32, green float32, blue float32, alpha float32)

var ClearBufferData func(target uint32, internalformat uint32, format uint32, xtype uint32, data unsafe.Pointer)

var ClearBufferSubData func(target uint32, internalformat uint32, offset int, size int, format uint32, xtype uint32, data unsafe.Pointer)
var ClearBufferfi func(buffer uint32, drawbuffer int32, depth float32, stencil int32)
var ClearBufferfv func(buffer uint32, drawbuffer int32, value *float32)
var ClearBufferiv func(buffer uint32, drawbuffer int32, value *int32)
var ClearBufferuiv func(buffer uint32, drawbuffer int32, value *uint32)

var ClearColor func(red float32, green float32, blue float32, alpha float32)

var ClearDepth func(depth float64)

var ClearDepthf func(d float32)

var ClearIndex func(c float32)

var ClearNamedBufferData func(buffer uint32, internalformat uint32, format uint32, xtype uint32, data unsafe.Pointer)
var ClearNamedBufferDataEXT func(buffer uint32, internalformat uint32, format uint32, xtype uint32, data unsafe.Pointer)

var ClearNamedBufferSubData func(buffer uint32, internalformat uint32, offset int, size int, format uint32, xtype uint32, data unsafe.Pointer)
var ClearNamedBufferSubDataEXT func(buffer uint32, internalformat uint32, offset int, size int, format uint32, xtype uint32, data unsafe.Pointer)
var ClearNamedFramebufferfi func(framebuffer uint32, buffer uint32, drawbuffer int32, depth float32, stencil int32)
var ClearNamedFramebufferfv func(framebuffer uint32, buffer uint32, drawbuffer int32, value *float32)
var ClearNamedFramebufferiv func(framebuffer uint32, buffer uint32, drawbuffer int32, value *int32)
var ClearNamedFramebufferuiv func(framebuffer uint32, buffer uint32, drawbuffer int32, value *uint32)

var ClearStencil func(s int32)

var ClearTexImage func(texture uint32, level int32, format uint32, xtype uint32, data unsafe.Pointer)

var ClearTexSubImage func(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, data unsafe.Pointer)

var ClientActiveTexture func(texture uint32)
var ClientAttribDefaultEXT func(mask uint32)

var ClientWaitSync func(sync uintptr, flags uint32, timeout uint64) uint32

var ClipControl func(origin uint32, depth uint32)

var ClipPlane func(plane uint32, equation *float64)
var Color3b func(red int8, green int8, blue int8)
var Color3bv func(v *int8)
var Color3d func(red float64, green float64, blue float64)
var Color3dv func(v *float64)
var Color3f func(red float32, green float32, blue float32)
var Color3fv func(v *float32)
var Color3i func(red int32, green int32, blue int32)
var Color3iv func(v *int32)
var Color3s func(red int16, green int16, blue int16)
var Color3sv func(v *int16)
var Color3ub func(red uint8, green uint8, blue uint8)
var Color3ubv func(v *uint8)
var Color3ui func(red uint32, green uint32, blue uint32)
var Color3uiv func(v *uint32)
var Color3us func(red uint16, green uint16, blue uint16)
var Color3usv func(v *uint16)
var Color4b func(red int8, green int8, blue int8, alpha int8)
var Color4bv func(v *int8)
var Color4d func(red float64, green float64, blue float64, alpha float64)
var Color4dv func(v *float64)
var Color4f func(red float32, green float32, blue float32, alpha float32)
var Color4fv func(v *float32)
var Color4i func(red int32, green int32, blue int32, alpha int32)
var Color4iv func(v *int32)
var Color4s func(red int16, green int16, blue int16, alpha int16)
var Color4sv func(v *int16)
var Color4ub func(red uint8, green uint8, blue uint8, alpha uint8)
var Color4ubv func(v *uint8)
var Color4ui func(red uint32, green uint32, blue uint32, alpha uint32)
var Color4uiv func(v *uint32)
var Color4us func(red uint16, green uint16, blue uint16, alpha uint16)
var Color4usv func(v *uint16)
var ColorFormatNV func(size int32, xtype uint32, stride int32)
var ColorMask func(red bool, green bool, blue bool, alpha bool)
var ColorMaski func(index uint32, r bool, g bool, b bool, a bool)

var ColorMaterial func(face uint32, mode uint32)

var ColorPointer func(size int32, xtype uint32, stride int32, pointer unsafe.Pointer)
var CommandListSegmentsNV func(list uint32, segments uint32)
var CompileCommandListNV func(list uint32)

var CompileShader func(shader uint32)
var CompileShaderIncludeARB func(shader uint32, count int32, path **uint8, length *int32)
var CompressedMultiTexImage1DEXT func(texunit uint32, target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, bits unsafe.Pointer)
var CompressedMultiTexImage2DEXT func(texunit uint32, target uint32, level int32, internalformat uint32, width int32, height int32, border int32, imageSize int32, bits unsafe.Pointer)
var CompressedMultiTexImage3DEXT func(texunit uint32, target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, bits unsafe.Pointer)
var CompressedMultiTexSubImage1DEXT func(texunit uint32, target uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, bits unsafe.Pointer)
var CompressedMultiTexSubImage2DEXT func(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, bits unsafe.Pointer)
var CompressedMultiTexSubImage3DEXT func(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, bits unsafe.Pointer)

var CompressedTexImage1D func(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer)

var CompressedTexImage2D func(target uint32, level int32, internalformat uint32, width int32, height int32, border int32, imageSize int32, data unsafe.Pointer)

var CompressedTexImage3D func(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, data unsafe.Pointer)

var CompressedTexSubImage1D func(target uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer)

var CompressedTexSubImage2D func(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, data unsafe.Pointer)

var CompressedTexSubImage3D func(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer)
var CompressedTextureImage1DEXT func(texture uint32, target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, bits unsafe.Pointer)
var CompressedTextureImage2DEXT func(texture uint32, target uint32, level int32, internalformat uint32, width int32, height int32, border int32, imageSize int32, bits unsafe.Pointer)
var CompressedTextureImage3DEXT func(texture uint32, target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, bits unsafe.Pointer)

var CompressedTextureSubImage1D func(texture uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer)
var CompressedTextureSubImage1DEXT func(texture uint32, target uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, bits unsafe.Pointer)

var CompressedTextureSubImage2D func(texture uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, data unsafe.Pointer)
var CompressedTextureSubImage2DEXT func(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, bits unsafe.Pointer)

var CompressedTextureSubImage3D func(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer)
var CompressedTextureSubImage3DEXT func(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, bits unsafe.Pointer)
var ConservativeRasterParameterfNV func(pname uint32, value float32)
var ConservativeRasterParameteriNV func(pname uint32, param int32)

var CopyBufferSubData func(readTarget uint32, writeTarget uint32, readOffset int, writeOffset int, size int)

var CopyImageSubData func(srcName uint32, srcTarget uint32, srcLevel int32, srcX int32, srcY int32, srcZ int32, dstName uint32, dstTarget uint32, dstLevel int32, dstX int32, dstY int32, dstZ int32, srcWidth int32, srcHeight int32, srcDepth int32)
var CopyMultiTexImage1DEXT func(texunit uint32, target uint32, level int32, internalformat uint32, x int32, y int32, width int32, border int32)
var CopyMultiTexImage2DEXT func(texunit uint32, target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32)
var CopyMultiTexSubImage1DEXT func(texunit uint32, target uint32, level int32, xoffset int32, x int32, y int32, width int32)
var CopyMultiTexSubImage2DEXT func(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32)
var CopyMultiTexSubImage3DEXT func(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32)

var CopyNamedBufferSubData func(readBuffer uint32, writeBuffer uint32, readOffset int, writeOffset int, size int)
var CopyPathNV func(resultPath uint32, srcPath uint32)

var CopyPixels func(x int32, y int32, width int32, height int32, xtype uint32)

var CopyTexImage1D func(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, border int32)

var CopyTexImage2D func(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32)

var CopyTexSubImage1D func(target uint32, level int32, xoffset int32, x int32, y int32, width int32)

var CopyTexSubImage2D func(target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32)

var CopyTexSubImage3D func(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32)
var CopyTextureImage1DEXT func(texture uint32, target uint32, level int32, internalformat uint32, x int32, y int32, width int32, border int32)
var CopyTextureImage2DEXT func(texture uint32, target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32)

var CopyTextureSubImage1D func(texture uint32, level int32, xoffset int32, x int32, y int32, width int32)
var CopyTextureSubImage1DEXT func(texture uint32, target uint32, level int32, xoffset int32, x int32, y int32, width int32)

var CopyTextureSubImage2D func(texture uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32)
var CopyTextureSubImage2DEXT func(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32)

var CopyTextureSubImage3D func(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32)
var CopyTextureSubImage3DEXT func(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32)
var CoverFillPathInstancedNV func(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, coverMode uint32, transformType uint32, transformValues *float32)
var CoverFillPathNV func(path uint32, coverMode uint32)
var CoverStrokePathInstancedNV func(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, coverMode uint32, transformType uint32, transformValues *float32)
var CoverStrokePathNV func(path uint32, coverMode uint32)
var CoverageModulationNV func(components uint32)
var CoverageModulationTableNV func(n int32, v *float32)

var CreateBuffers func(n int32, buffers *uint32)
var CreateCommandListsNV func(n int32, lists *uint32)

var CreateFramebuffers func(n int32, framebuffers *uint32)
var CreatePerfQueryINTEL func(queryId uint32, queryHandle *uint32)

var CreateProgram func() uint32

var CreateProgramPipelines func(n int32, pipelines *uint32)

var CreateQueries func(target uint32, n int32, ids *uint32)

var CreateRenderbuffers func(n int32, renderbuffers *uint32)

var CreateSamplers func(n int32, samplers *uint32)

var CreateShader func(xtype uint32) uint32
var CreateShaderProgramEXT func(xtype uint32, xstring *uint8) uint32

var CreateShaderProgramv func(xtype uint32, count int32, strings **uint8) uint32
var CreateShaderProgramvEXT func(xtype uint32, count int32, strings **uint8) uint32
var CreateStatesNV func(n int32, states *uint32)

var CreateSyncFromCLeventARB func(context unsafe.Pointer, event unsafe.Pointer, flags uint32) uintptr

var CreateTextures func(target uint32, n int32, textures *uint32)

var CreateTransformFeedbacks func(n int32, ids *uint32)

var CreateVertexArrays func(n int32, arrays *uint32)

var CullFace func(mode uint32)

// var DebugMessageCallback func(callback DebugProc, userParam unsafe.Pointer)
// var DebugMessageCallbackARB func(callback DebugProc, userParam unsafe.Pointer)
// var DebugMessageCallbackKHR func(callback DebugProc, userParam unsafe.Pointer)

var DebugMessageControl func(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool)
var DebugMessageControlARB func(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool)
var DebugMessageControlKHR func(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool)

var DebugMessageInsert func(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8)
var DebugMessageInsertARB func(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8)
var DebugMessageInsertKHR func(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8)

var DeleteBuffers func(n int32, buffers *uint32)
var DeleteCommandListsNV func(n int32, lists *uint32)

var DeleteFramebuffers func(n int32, framebuffers *uint32)

var DeleteLists func(list uint32, xrange int32)
var DeleteNamedStringARB func(namelen int32, name *uint8)
var DeletePathsNV func(path uint32, xrange int32)
var DeletePerfMonitorsAMD func(n int32, monitors *uint32)
var DeletePerfQueryINTEL func(queryHandle uint32)

var DeleteProgram func(program uint32)

var DeleteProgramPipelines func(n int32, pipelines *uint32)
var DeleteProgramPipelinesEXT func(n int32, pipelines *uint32)

var DeleteQueries func(n int32, ids *uint32)

var DeleteRenderbuffers func(n int32, renderbuffers *uint32)

var DeleteSamplers func(count int32, samplers *uint32)

var DeleteShader func(shader uint32)
var DeleteStatesNV func(n int32, states *uint32)

var DeleteSync func(sync uintptr)

var DeleteTextures func(n int32, textures *uint32)

var DeleteTransformFeedbacks func(n int32, ids *uint32)

var DeleteVertexArrays func(n int32, arrays *uint32)

var DepthFunc func(xfunc uint32)

var DepthMask func(flag bool)

var DepthRange func(n float64, f float64)
var DepthRangeArrayv func(first uint32, count int32, v *float64)

var DepthRangeIndexed func(index uint32, n float64, f float64)

var DepthRangef func(n float32, f float32)

var DetachShader func(program uint32, shader uint32)
var Disable func(cap uint32)
var DisableClientState func(array uint32)
var DisableClientStateIndexedEXT func(array uint32, index uint32)
var DisableClientStateiEXT func(array uint32, index uint32)
var DisableIndexedEXT func(target uint32, index uint32)

var DisableVertexArrayAttrib func(vaobj uint32, index uint32)
var DisableVertexArrayAttribEXT func(vaobj uint32, index uint32)
var DisableVertexArrayEXT func(vaobj uint32, array uint32)

var DisableVertexAttribArray func(index uint32)
var Disablei func(target uint32, index uint32)

var DispatchCompute func(num_groups_x uint32, num_groups_y uint32, num_groups_z uint32)
var DispatchComputeGroupSizeARB func(num_groups_x uint32, num_groups_y uint32, num_groups_z uint32, group_size_x uint32, group_size_y uint32, group_size_z uint32)

var DispatchComputeIndirect func(indirect int)

var DrawArrays func(mode uint32, first int32, count int32)

var DrawArraysIndirect func(mode uint32, indirect unsafe.Pointer)

var DrawArraysInstanced func(mode uint32, first int32, count int32, instancecount int32)
var DrawArraysInstancedARB func(mode uint32, first int32, count int32, primcount int32)

var DrawArraysInstancedBaseInstance func(mode uint32, first int32, count int32, instancecount int32, baseinstance uint32)
var DrawArraysInstancedEXT func(mode uint32, start int32, count int32, primcount int32)

var DrawBuffer func(buf uint32)

var DrawBuffers func(n int32, bufs *uint32)
var DrawCommandsAddressNV func(primitiveMode uint32, indirects *uint64, sizes *int32, count uint32)
var DrawCommandsNV func(primitiveMode uint32, buffer uint32, indirects *int, sizes *int32, count uint32)
var DrawCommandsStatesAddressNV func(indirects *uint64, sizes *int32, states *uint32, fbos *uint32, count uint32)
var DrawCommandsStatesNV func(buffer uint32, indirects *int, sizes *int32, states *uint32, fbos *uint32, count uint32)

var DrawElements func(mode uint32, count int32, xtype uint32, indices unsafe.Pointer)

var DrawElementsBaseVertex func(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32)

var DrawElementsIndirect func(mode uint32, xtype uint32, indirect unsafe.Pointer)

var DrawElementsInstanced func(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32)
var DrawElementsInstancedARB func(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, primcount int32)

var DrawElementsInstancedBaseInstance func(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, baseinstance uint32)

var DrawElementsInstancedBaseVertex func(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32)

var DrawElementsInstancedBaseVertexBaseInstance func(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32, baseinstance uint32)
var DrawElementsInstancedEXT func(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, primcount int32)

var DrawPixels func(width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var DrawRangeElements func(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer)

var DrawRangeElementsBaseVertex func(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32)

var DrawTransformFeedback func(mode uint32, id uint32)

var DrawTransformFeedbackInstanced func(mode uint32, id uint32, instancecount int32)

var DrawTransformFeedbackStream func(mode uint32, id uint32, stream uint32)

var DrawTransformFeedbackStreamInstanced func(mode uint32, id uint32, stream uint32, instancecount int32)
var DrawVkImageNV func(vkImage uint64, sampler uint32, x0 float32, y0 float32, x1 float32, y1 float32, z float32, s0 float32, t0 float32, s1 float32, t1 float32)

var EGLImageTargetTexStorageEXT func(target uint32, image unsafe.Pointer, attrib_list *int32)

var EGLImageTargetTextureStorageEXT func(texture uint32, image unsafe.Pointer, attrib_list *int32)

var EdgeFlag func(flag bool)
var EdgeFlagFormatNV func(stride int32)

var EdgeFlagPointer func(stride int32, pointer unsafe.Pointer)
var EdgeFlagv func(flag *bool)

var Enable func(cap uint32)

var EnableClientState func(array uint32)
var EnableClientStateIndexedEXT func(array uint32, index uint32)
var EnableClientStateiEXT func(array uint32, index uint32)
var EnableIndexedEXT func(target uint32, index uint32)

var EnableVertexArrayAttrib func(vaobj uint32, index uint32)
var EnableVertexArrayAttribEXT func(vaobj uint32, index uint32)
var EnableVertexArrayEXT func(vaobj uint32, array uint32)

var EnableVertexAttribArray func(index uint32)
var Enablei func(target uint32, index uint32)
var End func()
var EndConditionalRender func()
var EndConditionalRenderNV func()
var EndList func()
var EndPerfMonitorAMD func(monitor uint32)
var EndPerfQueryINTEL func(queryHandle uint32)
var EndQuery func(target uint32)
var EndQueryIndexed func(target uint32, index uint32)
var EndTransformFeedback func()
var EvalCoord1d func(u float64)
var EvalCoord1dv func(u *float64)
var EvalCoord1f func(u float32)
var EvalCoord1fv func(u *float32)
var EvalCoord2d func(u float64, v float64)
var EvalCoord2dv func(u *float64)
var EvalCoord2f func(u float32, v float32)
var EvalCoord2fv func(u *float32)
var EvalMesh1 func(mode uint32, i1 int32, i2 int32)
var EvalMesh2 func(mode uint32, i1 int32, i2 int32, j1 int32, j2 int32)
var EvalPoint1 func(i int32)
var EvalPoint2 func(i int32, j int32)
var EvaluateDepthValuesARB func()

var FeedbackBuffer func(size int32, xtype uint32, buffer *float32)

var FenceSync func(condition uint32, flags uint32) uintptr

var Finish func()

var Flush func()

var FlushMappedBufferRange func(target uint32, offset int, length int)

var FlushMappedNamedBufferRange func(buffer uint32, offset int, length int)
var FlushMappedNamedBufferRangeEXT func(buffer uint32, offset int, length int)
var FogCoordFormatNV func(xtype uint32, stride int32)

var FogCoordPointer func(xtype uint32, stride int32, pointer unsafe.Pointer)
var FogCoordd func(coord float64)
var FogCoorddv func(coord *float64)
var FogCoordf func(coord float32)
var FogCoordfv func(coord *float32)
var Fogf func(pname uint32, param float32)
var Fogfv func(pname uint32, params *float32)
var Fogi func(pname uint32, param int32)
var Fogiv func(pname uint32, params *int32)
var FragmentCoverageColorNV func(color uint32)
var FramebufferDrawBufferEXT func(framebuffer uint32, mode uint32)
var FramebufferDrawBuffersEXT func(framebuffer uint32, n int32, bufs *uint32)
var FramebufferFetchBarrierEXT func()

var FramebufferParameteri func(target uint32, pname uint32, param int32)
var FramebufferReadBufferEXT func(framebuffer uint32, mode uint32)

var FramebufferRenderbuffer func(target uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32)
var FramebufferSampleLocationsfvARB func(target uint32, start uint32, count int32, v *float32)
var FramebufferSampleLocationsfvNV func(target uint32, start uint32, count int32, v *float32)

var FramebufferTexture func(target uint32, attachment uint32, texture uint32, level int32)
var FramebufferTexture1D func(target uint32, attachment uint32, textarget uint32, texture uint32, level int32)

var FramebufferTexture2D func(target uint32, attachment uint32, textarget uint32, texture uint32, level int32)
var FramebufferTexture3D func(target uint32, attachment uint32, textarget uint32, texture uint32, level int32, zoffset int32)
var FramebufferTextureARB func(target uint32, attachment uint32, texture uint32, level int32)
var FramebufferTextureFaceARB func(target uint32, attachment uint32, texture uint32, level int32, face uint32)

var FramebufferTextureLayer func(target uint32, attachment uint32, texture uint32, level int32, layer int32)
var FramebufferTextureLayerARB func(target uint32, attachment uint32, texture uint32, level int32, layer int32)
var FramebufferTextureMultiviewOVR func(target uint32, attachment uint32, texture uint32, level int32, baseViewIndex int32, numViews int32)

var FrontFace func(mode uint32)

var Frustum func(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64)

var GenBuffers func(n int32, buffers *uint32)

var GenFramebuffers func(n int32, framebuffers *uint32)

var GenLists func(xrange int32) uint32
var GenPathsNV func(xrange int32) uint32
var GenPerfMonitorsAMD func(n int32, monitors *uint32)

var GenProgramPipelines func(n int32, pipelines *uint32)
var GenProgramPipelinesEXT func(n int32, pipelines *uint32)

var GenQueries func(n int32, ids *uint32)

var GenRenderbuffers func(n int32, renderbuffers *uint32)

var GenSamplers func(count int32, samplers *uint32)

var GenTextures func(n int32, textures *uint32)

var GenTransformFeedbacks func(n int32, ids *uint32)

var GenVertexArrays func(n int32, arrays *uint32)

var GenerateMipmap func(target uint32)
var GenerateMultiTexMipmapEXT func(texunit uint32, target uint32)

var GenerateTextureMipmap func(texture uint32)
var GenerateTextureMipmapEXT func(texture uint32, target uint32)

var GetActiveAtomicCounterBufferiv func(program uint32, bufferIndex uint32, pname uint32, params *int32)

var GetActiveAttrib func(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8)

var GetActiveSubroutineName func(program uint32, shadertype uint32, index uint32, bufsize int32, length *int32, name *uint8)

var GetActiveSubroutineUniformName func(program uint32, shadertype uint32, index uint32, bufsize int32, length *int32, name *uint8)
var GetActiveSubroutineUniformiv func(program uint32, shadertype uint32, index uint32, pname uint32, values *int32)

var GetActiveUniform func(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8)

var GetActiveUniformBlockName func(program uint32, uniformBlockIndex uint32, bufSize int32, length *int32, uniformBlockName *uint8)

var GetActiveUniformBlockiv func(program uint32, uniformBlockIndex uint32, pname uint32, params *int32)

var GetActiveUniformName func(program uint32, uniformIndex uint32, bufSize int32, length *int32, uniformName *uint8)

var GetActiveUniformsiv func(program uint32, uniformCount int32, uniformIndices *uint32, pname uint32, params *int32)

var GetAttachedShaders func(program uint32, maxCount int32, count *int32, shaders *uint32)

var GetAttribLocation func(program uint32, name *uint8) int32
var GetBooleanIndexedvEXT func(target uint32, index uint32, data *bool)
var GetBooleani_v func(target uint32, index uint32, data *bool)
var GetBooleanv func(pname uint32, data *bool)

var GetBufferParameteri64v func(target uint32, pname uint32, params *int64)

var GetBufferParameteriv func(target uint32, pname uint32, params *int32)
var GetBufferParameterui64vNV func(target uint32, pname uint32, params *uint64)

var GetBufferPointerv func(target uint32, pname uint32, params *unsafe.Pointer)

var GetBufferSubData func(target uint32, offset int, size int, data unsafe.Pointer)

var GetClipPlane func(plane uint32, equation *float64)
var GetCommandHeaderNV func(tokenID uint32, size uint32) uint32
var GetCompressedMultiTexImageEXT func(texunit uint32, target uint32, lod int32, img unsafe.Pointer)

var GetCompressedTexImage func(target uint32, level int32, img unsafe.Pointer)

var GetCompressedTextureImage func(texture uint32, level int32, bufSize int32, pixels unsafe.Pointer)
var GetCompressedTextureImageEXT func(texture uint32, target uint32, lod int32, img unsafe.Pointer)

var GetCompressedTextureSubImage func(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, bufSize int32, pixels unsafe.Pointer)
var GetCoverageModulationTableNV func(bufsize int32, v *float32)

var GetDebugMessageLog func(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32
var GetDebugMessageLogARB func(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32
var GetDebugMessageLogKHR func(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32
var GetDoubleIndexedvEXT func(target uint32, index uint32, data *float64)
var GetDoublei_v func(target uint32, index uint32, data *float64)
var GetDoublei_vEXT func(pname uint32, index uint32, params *float64)
var GetDoublev func(pname uint32, data *float64)

var GetError func() uint32
var GetFirstPerfQueryIdINTEL func(queryId *uint32)
var GetFloatIndexedvEXT func(target uint32, index uint32, data *float32)
var GetFloati_v func(target uint32, index uint32, data *float32)
var GetFloati_vEXT func(pname uint32, index uint32, params *float32)
var GetFloatv func(pname uint32, data *float32)

var GetFragDataIndex func(program uint32, name *uint8) int32

var GetFragDataLocation func(program uint32, name *uint8) int32

var GetFramebufferAttachmentParameteriv func(target uint32, attachment uint32, pname uint32, params *int32)

var GetFramebufferParameteriv func(target uint32, pname uint32, params *int32)
var GetFramebufferParameterivEXT func(framebuffer uint32, pname uint32, params *int32)

var GetGraphicsResetStatus func() uint32
var GetGraphicsResetStatusARB func() uint32
var GetGraphicsResetStatusKHR func() uint32
var GetImageHandleARB func(texture uint32, level int32, layered bool, layer int32, format uint32) uint64
var GetImageHandleNV func(texture uint32, level int32, layered bool, layer int32, format uint32) uint64
var GetInteger64i_v func(target uint32, index uint32, data *int64)
var GetInteger64v func(pname uint32, data *int64)
var GetIntegerIndexedvEXT func(target uint32, index uint32, data *int32)
var GetIntegeri_v func(target uint32, index uint32, data *int32)
var GetIntegerui64i_vNV func(value uint32, index uint32, result *uint64)
var GetIntegerui64vNV func(value uint32, result *uint64)
var GetIntegerv func(pname uint32, data *int32)
var GetInternalformatSampleivNV func(target uint32, internalformat uint32, samples int32, pname uint32, bufSize int32, params *int32)
var GetInternalformati64v func(target uint32, internalformat uint32, pname uint32, bufSize int32, params *int64)

var GetInternalformativ func(target uint32, internalformat uint32, pname uint32, bufSize int32, params *int32)
var GetLightfv func(light uint32, pname uint32, params *float32)
var GetLightiv func(light uint32, pname uint32, params *int32)
var GetMapdv func(target uint32, query uint32, v *float64)
var GetMapfv func(target uint32, query uint32, v *float32)
var GetMapiv func(target uint32, query uint32, v *int32)
var GetMaterialfv func(face uint32, pname uint32, params *float32)
var GetMaterialiv func(face uint32, pname uint32, params *int32)
var GetMultiTexEnvfvEXT func(texunit uint32, target uint32, pname uint32, params *float32)
var GetMultiTexEnvivEXT func(texunit uint32, target uint32, pname uint32, params *int32)
var GetMultiTexGendvEXT func(texunit uint32, coord uint32, pname uint32, params *float64)
var GetMultiTexGenfvEXT func(texunit uint32, coord uint32, pname uint32, params *float32)
var GetMultiTexGenivEXT func(texunit uint32, coord uint32, pname uint32, params *int32)
var GetMultiTexImageEXT func(texunit uint32, target uint32, level int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var GetMultiTexLevelParameterfvEXT func(texunit uint32, target uint32, level int32, pname uint32, params *float32)
var GetMultiTexLevelParameterivEXT func(texunit uint32, target uint32, level int32, pname uint32, params *int32)
var GetMultiTexParameterIivEXT func(texunit uint32, target uint32, pname uint32, params *int32)
var GetMultiTexParameterIuivEXT func(texunit uint32, target uint32, pname uint32, params *uint32)
var GetMultiTexParameterfvEXT func(texunit uint32, target uint32, pname uint32, params *float32)
var GetMultiTexParameterivEXT func(texunit uint32, target uint32, pname uint32, params *int32)

var GetMultisamplefv func(pname uint32, index uint32, val *float32)

var GetNamedBufferParameteri64v func(buffer uint32, pname uint32, params *int64)

var GetNamedBufferParameteriv func(buffer uint32, pname uint32, params *int32)
var GetNamedBufferParameterivEXT func(buffer uint32, pname uint32, params *int32)
var GetNamedBufferParameterui64vNV func(buffer uint32, pname uint32, params *uint64)

var GetNamedBufferPointerv func(buffer uint32, pname uint32, params *unsafe.Pointer)
var GetNamedBufferPointervEXT func(buffer uint32, pname uint32, params *unsafe.Pointer)

var GetNamedBufferSubData func(buffer uint32, offset int, size int, data unsafe.Pointer)
var GetNamedBufferSubDataEXT func(buffer uint32, offset int, size int, data unsafe.Pointer)

var GetNamedFramebufferAttachmentParameteriv func(framebuffer uint32, attachment uint32, pname uint32, params *int32)
var GetNamedFramebufferAttachmentParameterivEXT func(framebuffer uint32, attachment uint32, pname uint32, params *int32)

var GetNamedFramebufferParameteriv func(framebuffer uint32, pname uint32, param *int32)
var GetNamedFramebufferParameterivEXT func(framebuffer uint32, pname uint32, params *int32)
var GetNamedProgramLocalParameterIivEXT func(program uint32, target uint32, index uint32, params *int32)
var GetNamedProgramLocalParameterIuivEXT func(program uint32, target uint32, index uint32, params *uint32)
var GetNamedProgramLocalParameterdvEXT func(program uint32, target uint32, index uint32, params *float64)
var GetNamedProgramLocalParameterfvEXT func(program uint32, target uint32, index uint32, params *float32)
var GetNamedProgramStringEXT func(program uint32, target uint32, pname uint32, xstring unsafe.Pointer)
var GetNamedProgramivEXT func(program uint32, target uint32, pname uint32, params *int32)

var GetNamedRenderbufferParameteriv func(renderbuffer uint32, pname uint32, params *int32)
var GetNamedRenderbufferParameterivEXT func(renderbuffer uint32, pname uint32, params *int32)
var GetNamedStringARB func(namelen int32, name *uint8, bufSize int32, stringlen *int32, xstring *uint8)
var GetNamedStringivARB func(namelen int32, name *uint8, pname uint32, params *int32)
var GetNextPerfQueryIdINTEL func(queryId uint32, nextQueryId *uint32)

var GetObjectLabel func(identifier uint32, name uint32, bufSize int32, length *int32, label *uint8)
var GetObjectLabelEXT func(xtype uint32, object uint32, bufSize int32, length *int32, label *uint8)
var GetObjectLabelKHR func(identifier uint32, name uint32, bufSize int32, length *int32, label *uint8)

var GetObjectPtrLabel func(ptr unsafe.Pointer, bufSize int32, length *int32, label *uint8)
var GetObjectPtrLabelKHR func(ptr unsafe.Pointer, bufSize int32, length *int32, label *uint8)
var GetPathCommandsNV func(path uint32, commands *uint8)
var GetPathCoordsNV func(path uint32, coords *float32)
var GetPathDashArrayNV func(path uint32, dashArray *float32)
var GetPathLengthNV func(path uint32, startSegment int32, numSegments int32) float32
var GetPathMetricRangeNV func(metricQueryMask uint32, firstPathName uint32, numPaths int32, stride int32, metrics *float32)
var GetPathMetricsNV func(metricQueryMask uint32, numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, stride int32, metrics *float32)
var GetPathParameterfvNV func(path uint32, pname uint32, value *float32)
var GetPathParameterivNV func(path uint32, pname uint32, value *int32)
var GetPathSpacingNV func(pathListMode uint32, numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, advanceScale float32, kerningScale float32, transformType uint32, returnedSpacing *float32)
var GetPerfCounterInfoINTEL func(queryId uint32, counterId uint32, counterNameLength uint32, counterName *uint8, counterDescLength uint32, counterDesc *uint8, counterOffset *uint32, counterDataSize *uint32, counterTypeEnum *uint32, counterDataTypeEnum *uint32, rawCounterMaxValue *uint64)
var GetPerfMonitorCounterDataAMD func(monitor uint32, pname uint32, dataSize int32, data *uint32, bytesWritten *int32)
var GetPerfMonitorCounterInfoAMD func(group uint32, counter uint32, pname uint32, data unsafe.Pointer)
var GetPerfMonitorCounterStringAMD func(group uint32, counter uint32, bufSize int32, length *int32, counterString *uint8)
var GetPerfMonitorCountersAMD func(group uint32, numCounters *int32, maxActiveCounters *int32, counterSize int32, counters *uint32)
var GetPerfMonitorGroupStringAMD func(group uint32, bufSize int32, length *int32, groupString *uint8)
var GetPerfMonitorGroupsAMD func(numGroups *int32, groupsSize int32, groups *uint32)
var GetPerfQueryDataINTEL func(queryHandle uint32, flags uint32, dataSize int32, data unsafe.Pointer, bytesWritten *uint32)
var GetPerfQueryIdByNameINTEL func(queryName *uint8, queryId *uint32)
var GetPerfQueryInfoINTEL func(queryId uint32, queryNameLength uint32, queryName *uint8, dataSize *uint32, noCounters *uint32, noInstances *uint32, capsMask *uint32)
var GetPixelMapfv func(xmap uint32, values *float32)
var GetPixelMapuiv func(xmap uint32, values *uint32)
var GetPixelMapusv func(xmap uint32, values *uint16)
var GetPointerIndexedvEXT func(target uint32, index uint32, data *unsafe.Pointer)
var GetPointeri_vEXT func(pname uint32, index uint32, params *unsafe.Pointer)

var GetPointerv func(pname uint32, params *unsafe.Pointer)
var GetPointervKHR func(pname uint32, params *unsafe.Pointer)

var GetPolygonStipple func(mask *uint8)

var GetProgramBinary func(program uint32, bufSize int32, length *int32, binaryFormat *uint32, binary unsafe.Pointer)

var GetProgramInfoLog func(program uint32, bufSize int32, length *int32, infoLog *uint8)
var GetProgramInterfaceiv func(program uint32, programInterface uint32, pname uint32, params *int32)

var GetProgramPipelineInfoLog func(pipeline uint32, bufSize int32, length *int32, infoLog *uint8)
var GetProgramPipelineInfoLogEXT func(pipeline uint32, bufSize int32, length *int32, infoLog *uint8)
var GetProgramPipelineiv func(pipeline uint32, pname uint32, params *int32)
var GetProgramPipelineivEXT func(pipeline uint32, pname uint32, params *int32)

var GetProgramResourceIndex func(program uint32, programInterface uint32, name *uint8) uint32

var GetProgramResourceLocation func(program uint32, programInterface uint32, name *uint8) int32

var GetProgramResourceLocationIndex func(program uint32, programInterface uint32, name *uint8) int32

var GetProgramResourceName func(program uint32, programInterface uint32, index uint32, bufSize int32, length *int32, name *uint8)
var GetProgramResourcefvNV func(program uint32, programInterface uint32, index uint32, propCount int32, props *uint32, bufSize int32, length *int32, params *float32)
var GetProgramResourceiv func(program uint32, programInterface uint32, index uint32, propCount int32, props *uint32, bufSize int32, length *int32, params *int32)
var GetProgramStageiv func(program uint32, shadertype uint32, pname uint32, values *int32)

var GetProgramiv func(program uint32, pname uint32, params *int32)
var GetQueryBufferObjecti64v func(id uint32, buffer uint32, pname uint32, offset int)
var GetQueryBufferObjectiv func(id uint32, buffer uint32, pname uint32, offset int)
var GetQueryBufferObjectui64v func(id uint32, buffer uint32, pname uint32, offset int)
var GetQueryBufferObjectuiv func(id uint32, buffer uint32, pname uint32, offset int)

var GetQueryIndexediv func(target uint32, index uint32, pname uint32, params *int32)
var GetQueryObjecti64v func(id uint32, pname uint32, params *int64)
var GetQueryObjectiv func(id uint32, pname uint32, params *int32)
var GetQueryObjectui64v func(id uint32, pname uint32, params *uint64)

var GetQueryObjectuiv func(id uint32, pname uint32, params *uint32)

var GetQueryiv func(target uint32, pname uint32, params *int32)

var GetRenderbufferParameteriv func(target uint32, pname uint32, params *int32)
var GetSamplerParameterIiv func(sampler uint32, pname uint32, params *int32)
var GetSamplerParameterIuiv func(sampler uint32, pname uint32, params *uint32)
var GetSamplerParameterfv func(sampler uint32, pname uint32, params *float32)
var GetSamplerParameteriv func(sampler uint32, pname uint32, params *int32)

var GetShaderInfoLog func(shader uint32, bufSize int32, length *int32, infoLog *uint8)

var GetShaderPrecisionFormat func(shadertype uint32, precisiontype uint32, xrange *int32, precision *int32)

var GetShaderSource func(shader uint32, bufSize int32, length *int32, source *uint8)

var GetShaderiv func(shader uint32, pname uint32, params *int32)
var GetStageIndexNV func(shadertype uint32) uint16

var GetString func(name uint32) *uint8
var GetStringi func(name uint32, index uint32) *uint8

var GetSubroutineIndex func(program uint32, shadertype uint32, name *uint8) uint32

var GetSubroutineUniformLocation func(program uint32, shadertype uint32, name *uint8) int32

var GetSynciv func(sync uintptr, pname uint32, bufSize int32, length *int32, values *int32)
var GetTexEnvfv func(target uint32, pname uint32, params *float32)
var GetTexEnviv func(target uint32, pname uint32, params *int32)
var GetTexGendv func(coord uint32, pname uint32, params *float64)
var GetTexGenfv func(coord uint32, pname uint32, params *float32)
var GetTexGeniv func(coord uint32, pname uint32, params *int32)

var GetTexImage func(target uint32, level int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var GetTexLevelParameterfv func(target uint32, level int32, pname uint32, params *float32)
var GetTexLevelParameteriv func(target uint32, level int32, pname uint32, params *int32)
var GetTexParameterIiv func(target uint32, pname uint32, params *int32)
var GetTexParameterIuiv func(target uint32, pname uint32, params *uint32)
var GetTexParameterfv func(target uint32, pname uint32, params *float32)
var GetTexParameteriv func(target uint32, pname uint32, params *int32)
var GetTextureHandleARB func(texture uint32) uint64
var GetTextureHandleNV func(texture uint32) uint64

var GetTextureImage func(texture uint32, level int32, format uint32, xtype uint32, bufSize int32, pixels unsafe.Pointer)
var GetTextureImageEXT func(texture uint32, target uint32, level int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var GetTextureLevelParameterfv func(texture uint32, level int32, pname uint32, params *float32)
var GetTextureLevelParameterfvEXT func(texture uint32, target uint32, level int32, pname uint32, params *float32)
var GetTextureLevelParameteriv func(texture uint32, level int32, pname uint32, params *int32)
var GetTextureLevelParameterivEXT func(texture uint32, target uint32, level int32, pname uint32, params *int32)
var GetTextureParameterIiv func(texture uint32, pname uint32, params *int32)
var GetTextureParameterIivEXT func(texture uint32, target uint32, pname uint32, params *int32)
var GetTextureParameterIuiv func(texture uint32, pname uint32, params *uint32)
var GetTextureParameterIuivEXT func(texture uint32, target uint32, pname uint32, params *uint32)
var GetTextureParameterfv func(texture uint32, pname uint32, params *float32)
var GetTextureParameterfvEXT func(texture uint32, target uint32, pname uint32, params *float32)
var GetTextureParameteriv func(texture uint32, pname uint32, params *int32)
var GetTextureParameterivEXT func(texture uint32, target uint32, pname uint32, params *int32)
var GetTextureSamplerHandleARB func(texture uint32, sampler uint32) uint64
var GetTextureSamplerHandleNV func(texture uint32, sampler uint32) uint64

var GetTextureSubImage func(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, bufSize int32, pixels unsafe.Pointer)

var GetTransformFeedbackVarying func(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8)
var GetTransformFeedbacki64_v func(xfb uint32, pname uint32, index uint32, param *int64)
var GetTransformFeedbacki_v func(xfb uint32, pname uint32, index uint32, param *int32)

var GetTransformFeedbackiv func(xfb uint32, pname uint32, param *int32)

var GetUniformBlockIndex func(program uint32, uniformBlockName *uint8) uint32

var GetUniformIndices func(program uint32, uniformCount int32, uniformNames **uint8, uniformIndices *uint32)

var GetUniformLocation func(program uint32, name *uint8) int32
var GetUniformSubroutineuiv func(shadertype uint32, location int32, params *uint32)
var GetUniformdv func(program uint32, location int32, params *float64)

var GetUniformfv func(program uint32, location int32, params *float32)
var GetUniformi64vARB func(program uint32, location int32, params *int64)
var GetUniformi64vNV func(program uint32, location int32, params *int64)

var GetUniformiv func(program uint32, location int32, params *int32)
var GetUniformui64vARB func(program uint32, location int32, params *uint64)
var GetUniformui64vNV func(program uint32, location int32, params *uint64)
var GetUniformuiv func(program uint32, location int32, params *uint32)
var GetVertexArrayIndexed64iv func(vaobj uint32, index uint32, pname uint32, param *int64)
var GetVertexArrayIndexediv func(vaobj uint32, index uint32, pname uint32, param *int32)
var GetVertexArrayIntegeri_vEXT func(vaobj uint32, index uint32, pname uint32, param *int32)
var GetVertexArrayIntegervEXT func(vaobj uint32, pname uint32, param *int32)
var GetVertexArrayPointeri_vEXT func(vaobj uint32, index uint32, pname uint32, param *unsafe.Pointer)
var GetVertexArrayPointervEXT func(vaobj uint32, pname uint32, param *unsafe.Pointer)

var GetVertexArrayiv func(vaobj uint32, pname uint32, param *int32)

var GetVertexAttribIiv func(index uint32, pname uint32, params *int32)

var GetVertexAttribIuiv func(index uint32, pname uint32, params *uint32)

var GetVertexAttribLdv func(index uint32, pname uint32, params *float64)
var GetVertexAttribLi64vNV func(index uint32, pname uint32, params *int64)
var GetVertexAttribLui64vARB func(index uint32, pname uint32, params *uint64)
var GetVertexAttribLui64vNV func(index uint32, pname uint32, params *uint64)

var GetVertexAttribPointerv func(index uint32, pname uint32, pointer *unsafe.Pointer)

var GetVertexAttribdv func(index uint32, pname uint32, params *float64)

var GetVertexAttribfv func(index uint32, pname uint32, params *float32)

var GetVertexAttribiv func(index uint32, pname uint32, params *int32)

var GetVkProcAddrNV func(name *uint8) unsafe.Pointer

var GetnCompressedTexImage func(target uint32, lod int32, bufSize int32, pixels unsafe.Pointer)
var GetnCompressedTexImageARB func(target uint32, lod int32, bufSize int32, img unsafe.Pointer)

var GetnTexImage func(target uint32, level int32, format uint32, xtype uint32, bufSize int32, pixels unsafe.Pointer)
var GetnTexImageARB func(target uint32, level int32, format uint32, xtype uint32, bufSize int32, img unsafe.Pointer)
var GetnUniformdv func(program uint32, location int32, bufSize int32, params *float64)
var GetnUniformdvARB func(program uint32, location int32, bufSize int32, params *float64)
var GetnUniformfv func(program uint32, location int32, bufSize int32, params *float32)
var GetnUniformfvARB func(program uint32, location int32, bufSize int32, params *float32)
var GetnUniformfvKHR func(program uint32, location int32, bufSize int32, params *float32)
var GetnUniformi64vARB func(program uint32, location int32, bufSize int32, params *int64)
var GetnUniformiv func(program uint32, location int32, bufSize int32, params *int32)
var GetnUniformivARB func(program uint32, location int32, bufSize int32, params *int32)
var GetnUniformivKHR func(program uint32, location int32, bufSize int32, params *int32)
var GetnUniformui64vARB func(program uint32, location int32, bufSize int32, params *uint64)
var GetnUniformuiv func(program uint32, location int32, bufSize int32, params *uint32)
var GetnUniformuivARB func(program uint32, location int32, bufSize int32, params *uint32)
var GetnUniformuivKHR func(program uint32, location int32, bufSize int32, params *uint32)

var Hint func(target uint32, mode uint32)
var IndexFormatNV func(xtype uint32, stride int32)

var IndexMask func(mask uint32)

var IndexPointer func(xtype uint32, stride int32, pointer unsafe.Pointer)
var Indexd func(c float64)
var Indexdv func(c *float64)
var Indexf func(c float32)
var Indexfv func(c *float32)
var Indexi func(c int32)
var Indexiv func(c *int32)
var Indexs func(c int16)
var Indexsv func(c *int16)
var Indexub func(c uint8)
var Indexubv func(c *uint8)

var InitNames func()
var InsertEventMarkerEXT func(length int32, marker *uint8)

var InterleavedArrays func(format uint32, stride int32, pointer unsafe.Pointer)
var InterpolatePathsNV func(resultPath uint32, pathA uint32, pathB uint32, weight float32)

var InvalidateBufferData func(buffer uint32)

var InvalidateBufferSubData func(buffer uint32, offset int, length int)

var InvalidateFramebuffer func(target uint32, numAttachments int32, attachments *uint32)

var InvalidateNamedFramebufferData func(framebuffer uint32, numAttachments int32, attachments *uint32)

var InvalidateNamedFramebufferSubData func(framebuffer uint32, numAttachments int32, attachments *uint32, x int32, y int32, width int32, height int32)

var InvalidateSubFramebuffer func(target uint32, numAttachments int32, attachments *uint32, x int32, y int32, width int32, height int32)

var InvalidateTexImage func(texture uint32, level int32)

var InvalidateTexSubImage func(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32)

var IsBuffer func(buffer uint32) bool
var IsBufferResidentNV func(target uint32) bool
var IsCommandListNV func(list uint32) bool
var IsEnabled func(cap uint32) bool
var IsEnabledIndexedEXT func(target uint32, index uint32) bool
var IsEnabledi func(target uint32, index uint32) bool

var IsFramebuffer func(framebuffer uint32) bool
var IsImageHandleResidentARB func(handle uint64) bool
var IsImageHandleResidentNV func(handle uint64) bool

var IsList func(list uint32) bool
var IsNamedBufferResidentNV func(buffer uint32) bool
var IsNamedStringARB func(namelen int32, name *uint8) bool
var IsPathNV func(path uint32) bool
var IsPointInFillPathNV func(path uint32, mask uint32, x float32, y float32) bool
var IsPointInStrokePathNV func(path uint32, x float32, y float32) bool

var IsProgram func(program uint32) bool

var IsProgramPipeline func(pipeline uint32) bool
var IsProgramPipelineEXT func(pipeline uint32) bool

var IsQuery func(id uint32) bool

var IsRenderbuffer func(renderbuffer uint32) bool

var IsSampler func(sampler uint32) bool

var IsShader func(shader uint32) bool
var IsStateNV func(state uint32) bool

var IsSync func(sync uintptr) bool

var IsTexture func(texture uint32) bool
var IsTextureHandleResidentARB func(handle uint64) bool
var IsTextureHandleResidentNV func(handle uint64) bool

var IsTransformFeedback func(id uint32) bool

var IsVertexArray func(array uint32) bool
var LabelObjectEXT func(xtype uint32, object uint32, length int32, label *uint8)
var LightModelf func(pname uint32, param float32)
var LightModelfv func(pname uint32, params *float32)
var LightModeli func(pname uint32, param int32)
var LightModeliv func(pname uint32, params *int32)
var Lightf func(light uint32, pname uint32, param float32)
var Lightfv func(light uint32, pname uint32, params *float32)
var Lighti func(light uint32, pname uint32, param int32)
var Lightiv func(light uint32, pname uint32, params *int32)

var LineStipple func(factor int32, pattern uint16)

var LineWidth func(width float32)

var LinkProgram func(program uint32)

var ListBase func(base uint32)
var ListDrawCommandsStatesClientNV func(list uint32, segment uint32, indirects *unsafe.Pointer, sizes *int32, states *uint32, fbos *uint32, count uint32)

var LoadIdentity func()
var LoadMatrixd func(m *float64)
var LoadMatrixf func(m *float32)

var LoadName func(name uint32)
var LoadTransposeMatrixd func(m *float64)
var LoadTransposeMatrixf func(m *float32)

var LogicOp func(opcode uint32)
var MakeBufferNonResidentNV func(target uint32)
var MakeBufferResidentNV func(target uint32, access uint32)
var MakeImageHandleNonResidentARB func(handle uint64)
var MakeImageHandleNonResidentNV func(handle uint64)
var MakeImageHandleResidentARB func(handle uint64, access uint32)
var MakeImageHandleResidentNV func(handle uint64, access uint32)
var MakeNamedBufferNonResidentNV func(buffer uint32)
var MakeNamedBufferResidentNV func(buffer uint32, access uint32)
var MakeTextureHandleNonResidentARB func(handle uint64)
var MakeTextureHandleNonResidentNV func(handle uint64)
var MakeTextureHandleResidentARB func(handle uint64)
var MakeTextureHandleResidentNV func(handle uint64)
var Map1d func(target uint32, u1 float64, u2 float64, stride int32, order int32, points *float64)
var Map1f func(target uint32, u1 float32, u2 float32, stride int32, order int32, points *float32)
var Map2d func(target uint32, u1 float64, u2 float64, ustride int32, uorder int32, v1 float64, v2 float64, vstride int32, vorder int32, points *float64)
var Map2f func(target uint32, u1 float32, u2 float32, ustride int32, uorder int32, v1 float32, v2 float32, vstride int32, vorder int32, points *float32)

var MapBuffer func(target uint32, access uint32) unsafe.Pointer

var MapBufferRange func(target uint32, offset int, length int, access uint32) unsafe.Pointer
var MapGrid1d func(un int32, u1 float64, u2 float64)
var MapGrid1f func(un int32, u1 float32, u2 float32)
var MapGrid2d func(un int32, u1 float64, u2 float64, vn int32, v1 float64, v2 float64)
var MapGrid2f func(un int32, u1 float32, u2 float32, vn int32, v1 float32, v2 float32)

var MapNamedBuffer func(buffer uint32, access uint32) unsafe.Pointer
var MapNamedBufferEXT func(buffer uint32, access uint32) unsafe.Pointer

var MapNamedBufferRange func(buffer uint32, offset int, length int, access uint32) unsafe.Pointer
var MapNamedBufferRangeEXT func(buffer uint32, offset int, length int, access uint32) unsafe.Pointer
var Materialf func(face uint32, pname uint32, param float32)
var Materialfv func(face uint32, pname uint32, params *float32)
var Materiali func(face uint32, pname uint32, param int32)
var Materialiv func(face uint32, pname uint32, params *int32)
var MatrixFrustumEXT func(mode uint32, left float64, right float64, bottom float64, top float64, zNear float64, zFar float64)
var MatrixLoad3x2fNV func(matrixMode uint32, m *float32)
var MatrixLoad3x3fNV func(matrixMode uint32, m *float32)
var MatrixLoadIdentityEXT func(mode uint32)
var MatrixLoadTranspose3x3fNV func(matrixMode uint32, m *float32)
var MatrixLoadTransposedEXT func(mode uint32, m *float64)
var MatrixLoadTransposefEXT func(mode uint32, m *float32)
var MatrixLoaddEXT func(mode uint32, m *float64)
var MatrixLoadfEXT func(mode uint32, m *float32)

var MatrixMode func(mode uint32)
var MatrixMult3x2fNV func(matrixMode uint32, m *float32)
var MatrixMult3x3fNV func(matrixMode uint32, m *float32)
var MatrixMultTranspose3x3fNV func(matrixMode uint32, m *float32)
var MatrixMultTransposedEXT func(mode uint32, m *float64)
var MatrixMultTransposefEXT func(mode uint32, m *float32)
var MatrixMultdEXT func(mode uint32, m *float64)
var MatrixMultfEXT func(mode uint32, m *float32)
var MatrixOrthoEXT func(mode uint32, left float64, right float64, bottom float64, top float64, zNear float64, zFar float64)
var MatrixPopEXT func(mode uint32)
var MatrixPushEXT func(mode uint32)
var MatrixRotatedEXT func(mode uint32, angle float64, x float64, y float64, z float64)
var MatrixRotatefEXT func(mode uint32, angle float32, x float32, y float32, z float32)
var MatrixScaledEXT func(mode uint32, x float64, y float64, z float64)
var MatrixScalefEXT func(mode uint32, x float32, y float32, z float32)
var MatrixTranslatedEXT func(mode uint32, x float64, y float64, z float64)
var MatrixTranslatefEXT func(mode uint32, x float32, y float32, z float32)
var MaxShaderCompilerThreadsARB func(count uint32)
var MaxShaderCompilerThreadsKHR func(count uint32)

var MemoryBarrier func(barriers uint32)
var MemoryBarrierByRegion func(barriers uint32)

var MinSampleShading func(value float32)
var MinSampleShadingARB func(value float32)
var MultMatrixd func(m *float64)
var MultMatrixf func(m *float32)
var MultTransposeMatrixd func(m *float64)
var MultTransposeMatrixf func(m *float32)

var MultiDrawArrays func(mode uint32, first *int32, count *int32, drawcount int32)

var MultiDrawArraysIndirect func(mode uint32, indirect unsafe.Pointer, drawcount int32, stride int32)
var MultiDrawArraysIndirectBindlessCountNV func(mode uint32, indirect unsafe.Pointer, drawCount int32, maxDrawCount int32, stride int32, vertexBufferCount int32)
var MultiDrawArraysIndirectBindlessNV func(mode uint32, indirect unsafe.Pointer, drawCount int32, stride int32, vertexBufferCount int32)
var MultiDrawArraysIndirectCount func(mode uint32, indirect unsafe.Pointer, drawcount int, maxdrawcount int32, stride int32)
var MultiDrawArraysIndirectCountARB func(mode uint32, indirect unsafe.Pointer, drawcount int, maxdrawcount int32, stride int32)

var MultiDrawElements func(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, drawcount int32)

var MultiDrawElementsBaseVertex func(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, drawcount int32, basevertex *int32)

var MultiDrawElementsIndirect func(mode uint32, xtype uint32, indirect unsafe.Pointer, drawcount int32, stride int32)
var MultiDrawElementsIndirectBindlessCountNV func(mode uint32, xtype uint32, indirect unsafe.Pointer, drawCount int32, maxDrawCount int32, stride int32, vertexBufferCount int32)
var MultiDrawElementsIndirectBindlessNV func(mode uint32, xtype uint32, indirect unsafe.Pointer, drawCount int32, stride int32, vertexBufferCount int32)
var MultiDrawElementsIndirectCount func(mode uint32, xtype uint32, indirect unsafe.Pointer, drawcount int, maxdrawcount int32, stride int32)
var MultiDrawElementsIndirectCountARB func(mode uint32, xtype uint32, indirect unsafe.Pointer, drawcount int, maxdrawcount int32, stride int32)
var MultiTexBufferEXT func(texunit uint32, target uint32, internalformat uint32, buffer uint32)
var MultiTexCoord1d func(target uint32, s float64)
var MultiTexCoord1dv func(target uint32, v *float64)
var MultiTexCoord1f func(target uint32, s float32)
var MultiTexCoord1fv func(target uint32, v *float32)
var MultiTexCoord1i func(target uint32, s int32)
var MultiTexCoord1iv func(target uint32, v *int32)
var MultiTexCoord1s func(target uint32, s int16)
var MultiTexCoord1sv func(target uint32, v *int16)
var MultiTexCoord2d func(target uint32, s float64, t float64)
var MultiTexCoord2dv func(target uint32, v *float64)
var MultiTexCoord2f func(target uint32, s float32, t float32)
var MultiTexCoord2fv func(target uint32, v *float32)
var MultiTexCoord2i func(target uint32, s int32, t int32)
var MultiTexCoord2iv func(target uint32, v *int32)
var MultiTexCoord2s func(target uint32, s int16, t int16)
var MultiTexCoord2sv func(target uint32, v *int16)
var MultiTexCoord3d func(target uint32, s float64, t float64, r float64)
var MultiTexCoord3dv func(target uint32, v *float64)
var MultiTexCoord3f func(target uint32, s float32, t float32, r float32)
var MultiTexCoord3fv func(target uint32, v *float32)
var MultiTexCoord3i func(target uint32, s int32, t int32, r int32)
var MultiTexCoord3iv func(target uint32, v *int32)
var MultiTexCoord3s func(target uint32, s int16, t int16, r int16)
var MultiTexCoord3sv func(target uint32, v *int16)
var MultiTexCoord4d func(target uint32, s float64, t float64, r float64, q float64)
var MultiTexCoord4dv func(target uint32, v *float64)
var MultiTexCoord4f func(target uint32, s float32, t float32, r float32, q float32)
var MultiTexCoord4fv func(target uint32, v *float32)
var MultiTexCoord4i func(target uint32, s int32, t int32, r int32, q int32)
var MultiTexCoord4iv func(target uint32, v *int32)
var MultiTexCoord4s func(target uint32, s int16, t int16, r int16, q int16)
var MultiTexCoord4sv func(target uint32, v *int16)
var MultiTexCoordPointerEXT func(texunit uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer)
var MultiTexEnvfEXT func(texunit uint32, target uint32, pname uint32, param float32)
var MultiTexEnvfvEXT func(texunit uint32, target uint32, pname uint32, params *float32)
var MultiTexEnviEXT func(texunit uint32, target uint32, pname uint32, param int32)
var MultiTexEnvivEXT func(texunit uint32, target uint32, pname uint32, params *int32)
var MultiTexGendEXT func(texunit uint32, coord uint32, pname uint32, param float64)
var MultiTexGendvEXT func(texunit uint32, coord uint32, pname uint32, params *float64)
var MultiTexGenfEXT func(texunit uint32, coord uint32, pname uint32, param float32)
var MultiTexGenfvEXT func(texunit uint32, coord uint32, pname uint32, params *float32)
var MultiTexGeniEXT func(texunit uint32, coord uint32, pname uint32, param int32)
var MultiTexGenivEXT func(texunit uint32, coord uint32, pname uint32, params *int32)
var MultiTexImage1DEXT func(texunit uint32, target uint32, level int32, internalformat int32, width int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var MultiTexImage2DEXT func(texunit uint32, target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var MultiTexImage3DEXT func(texunit uint32, target uint32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var MultiTexParameterIivEXT func(texunit uint32, target uint32, pname uint32, params *int32)
var MultiTexParameterIuivEXT func(texunit uint32, target uint32, pname uint32, params *uint32)
var MultiTexParameterfEXT func(texunit uint32, target uint32, pname uint32, param float32)
var MultiTexParameterfvEXT func(texunit uint32, target uint32, pname uint32, params *float32)
var MultiTexParameteriEXT func(texunit uint32, target uint32, pname uint32, param int32)
var MultiTexParameterivEXT func(texunit uint32, target uint32, pname uint32, params *int32)
var MultiTexRenderbufferEXT func(texunit uint32, target uint32, renderbuffer uint32)
var MultiTexSubImage1DEXT func(texunit uint32, target uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var MultiTexSubImage2DEXT func(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var MultiTexSubImage3DEXT func(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var NamedBufferData func(buffer uint32, size int, data unsafe.Pointer, usage uint32)
var NamedBufferDataEXT func(buffer uint32, size int, data unsafe.Pointer, usage uint32)
var NamedBufferPageCommitmentARB func(buffer uint32, offset int, size int, commit bool)
var NamedBufferPageCommitmentEXT func(buffer uint32, offset int, size int, commit bool)

var NamedBufferStorage func(buffer uint32, size int, data unsafe.Pointer, flags uint32)
var NamedBufferStorageEXT func(buffer uint32, size int, data unsafe.Pointer, flags uint32)

var NamedBufferSubData func(buffer uint32, offset int, size int, data unsafe.Pointer)
var NamedBufferSubDataEXT func(buffer uint32, offset int, size int, data unsafe.Pointer)
var NamedCopyBufferSubDataEXT func(readBuffer uint32, writeBuffer uint32, readOffset int, writeOffset int, size int)

var NamedFramebufferDrawBuffer func(framebuffer uint32, buf uint32)

var NamedFramebufferDrawBuffers func(framebuffer uint32, n int32, bufs *uint32)

var NamedFramebufferParameteri func(framebuffer uint32, pname uint32, param int32)
var NamedFramebufferParameteriEXT func(framebuffer uint32, pname uint32, param int32)

var NamedFramebufferReadBuffer func(framebuffer uint32, src uint32)

var NamedFramebufferRenderbuffer func(framebuffer uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32)
var NamedFramebufferRenderbufferEXT func(framebuffer uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32)
var NamedFramebufferSampleLocationsfvARB func(framebuffer uint32, start uint32, count int32, v *float32)
var NamedFramebufferSampleLocationsfvNV func(framebuffer uint32, start uint32, count int32, v *float32)
var NamedFramebufferTexture func(framebuffer uint32, attachment uint32, texture uint32, level int32)
var NamedFramebufferTexture1DEXT func(framebuffer uint32, attachment uint32, textarget uint32, texture uint32, level int32)
var NamedFramebufferTexture2DEXT func(framebuffer uint32, attachment uint32, textarget uint32, texture uint32, level int32)
var NamedFramebufferTexture3DEXT func(framebuffer uint32, attachment uint32, textarget uint32, texture uint32, level int32, zoffset int32)
var NamedFramebufferTextureEXT func(framebuffer uint32, attachment uint32, texture uint32, level int32)
var NamedFramebufferTextureFaceEXT func(framebuffer uint32, attachment uint32, texture uint32, level int32, face uint32)

var NamedFramebufferTextureLayer func(framebuffer uint32, attachment uint32, texture uint32, level int32, layer int32)
var NamedFramebufferTextureLayerEXT func(framebuffer uint32, attachment uint32, texture uint32, level int32, layer int32)
var NamedProgramLocalParameter4dEXT func(program uint32, target uint32, index uint32, x float64, y float64, z float64, w float64)
var NamedProgramLocalParameter4dvEXT func(program uint32, target uint32, index uint32, params *float64)
var NamedProgramLocalParameter4fEXT func(program uint32, target uint32, index uint32, x float32, y float32, z float32, w float32)
var NamedProgramLocalParameter4fvEXT func(program uint32, target uint32, index uint32, params *float32)
var NamedProgramLocalParameterI4iEXT func(program uint32, target uint32, index uint32, x int32, y int32, z int32, w int32)
var NamedProgramLocalParameterI4ivEXT func(program uint32, target uint32, index uint32, params *int32)
var NamedProgramLocalParameterI4uiEXT func(program uint32, target uint32, index uint32, x uint32, y uint32, z uint32, w uint32)
var NamedProgramLocalParameterI4uivEXT func(program uint32, target uint32, index uint32, params *uint32)
var NamedProgramLocalParameters4fvEXT func(program uint32, target uint32, index uint32, count int32, params *float32)
var NamedProgramLocalParametersI4ivEXT func(program uint32, target uint32, index uint32, count int32, params *int32)
var NamedProgramLocalParametersI4uivEXT func(program uint32, target uint32, index uint32, count int32, params *uint32)
var NamedProgramStringEXT func(program uint32, target uint32, format uint32, len int32, xstring unsafe.Pointer)

var NamedRenderbufferStorage func(renderbuffer uint32, internalformat uint32, width int32, height int32)
var NamedRenderbufferStorageEXT func(renderbuffer uint32, internalformat uint32, width int32, height int32)

var NamedRenderbufferStorageMultisample func(renderbuffer uint32, samples int32, internalformat uint32, width int32, height int32)
var NamedRenderbufferStorageMultisampleCoverageEXT func(renderbuffer uint32, coverageSamples int32, colorSamples int32, internalformat uint32, width int32, height int32)
var NamedRenderbufferStorageMultisampleEXT func(renderbuffer uint32, samples int32, internalformat uint32, width int32, height int32)
var NamedStringARB func(xtype uint32, namelen int32, name *uint8, stringlen int32, xstring *uint8)

var NewList func(list uint32, mode uint32)
var Normal3b func(nx int8, ny int8, nz int8)
var Normal3bv func(v *int8)
var Normal3d func(nx float64, ny float64, nz float64)
var Normal3dv func(v *float64)
var Normal3f func(nx float32, ny float32, nz float32)
var Normal3fv func(v *float32)
var Normal3i func(nx int32, ny int32, nz int32)
var Normal3iv func(v *int32)
var Normal3s func(nx int16, ny int16, nz int16)
var Normal3sv func(v *int16)
var NormalFormatNV func(xtype uint32, stride int32)

var NormalPointer func(xtype uint32, stride int32, pointer unsafe.Pointer)

var ObjectLabel func(identifier uint32, name uint32, length int32, label *uint8)
var ObjectLabelKHR func(identifier uint32, name uint32, length int32, label *uint8)

var ObjectPtrLabel func(ptr unsafe.Pointer, length int32, label *uint8)
var ObjectPtrLabelKHR func(ptr unsafe.Pointer, length int32, label *uint8)

var Ortho func(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64)

var PassThrough func(token float32)
var PatchParameterfv func(pname uint32, values *float32)

var PatchParameteri func(pname uint32, value int32)
var PathCommandsNV func(path uint32, numCommands int32, commands *uint8, numCoords int32, coordType uint32, coords unsafe.Pointer)
var PathCoordsNV func(path uint32, numCoords int32, coordType uint32, coords unsafe.Pointer)
var PathCoverDepthFuncNV func(xfunc uint32)
var PathDashArrayNV func(path uint32, dashCount int32, dashArray *float32)
var PathGlyphIndexArrayNV func(firstPathName uint32, fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, firstGlyphIndex uint32, numGlyphs int32, pathParameterTemplate uint32, emScale float32) uint32
var PathGlyphIndexRangeNV func(fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, pathParameterTemplate uint32, emScale float32, baseAndCount *uint32) uint32
var PathGlyphRangeNV func(firstPathName uint32, fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, firstGlyph uint32, numGlyphs int32, handleMissingGlyphs uint32, pathParameterTemplate uint32, emScale float32)
var PathGlyphsNV func(firstPathName uint32, fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, numGlyphs int32, xtype uint32, charcodes unsafe.Pointer, handleMissingGlyphs uint32, pathParameterTemplate uint32, emScale float32)
var PathMemoryGlyphIndexArrayNV func(firstPathName uint32, fontTarget uint32, fontSize int, fontData unsafe.Pointer, faceIndex int32, firstGlyphIndex uint32, numGlyphs int32, pathParameterTemplate uint32, emScale float32) uint32
var PathParameterfNV func(path uint32, pname uint32, value float32)
var PathParameterfvNV func(path uint32, pname uint32, value *float32)
var PathParameteriNV func(path uint32, pname uint32, value int32)
var PathParameterivNV func(path uint32, pname uint32, value *int32)
var PathStencilDepthOffsetNV func(factor float32, units float32)
var PathStencilFuncNV func(xfunc uint32, ref int32, mask uint32)
var PathStringNV func(path uint32, format uint32, length int32, pathString unsafe.Pointer)
var PathSubCommandsNV func(path uint32, commandStart int32, commandsToDelete int32, numCommands int32, commands *uint8, numCoords int32, coordType uint32, coords unsafe.Pointer)
var PathSubCoordsNV func(path uint32, coordStart int32, numCoords int32, coordType uint32, coords unsafe.Pointer)

var PauseTransformFeedback func()
var PixelMapfv func(xmap uint32, mapsize int32, values *float32)
var PixelMapuiv func(xmap uint32, mapsize int32, values *uint32)
var PixelMapusv func(xmap uint32, mapsize int32, values *uint16)
var PixelStoref func(pname uint32, param float32)

var PixelStorei func(pname uint32, param int32)
var PixelTransferf func(pname uint32, param float32)
var PixelTransferi func(pname uint32, param int32)

var PixelZoom func(xfactor float32, yfactor float32)
var PointAlongPathNV func(path uint32, startSegment int32, numSegments int32, distance float32, x *float32, y *float32, tangentX *float32, tangentY *float32) bool
var PointParameterf func(pname uint32, param float32)
var PointParameterfv func(pname uint32, params *float32)
var PointParameteri func(pname uint32, param int32)
var PointParameteriv func(pname uint32, params *int32)

var PointSize func(size float32)

var PolygonMode func(face uint32, mode uint32)

var PolygonOffset func(factor float32, units float32)
var PolygonOffsetClamp func(factor float32, units float32, clamp float32)
var PolygonOffsetClampEXT func(factor float32, units float32, clamp float32)

var PolygonStipple func(mask *uint8)
var PopAttrib func()
var PopClientAttrib func()

var PopDebugGroup func()
var PopDebugGroupKHR func()
var PopGroupMarkerEXT func()
var PopMatrix func()
var PopName func()
var PrimitiveBoundingBoxARB func(minX float32, minY float32, minZ float32, minW float32, maxX float32, maxY float32, maxZ float32, maxW float32)

var PrimitiveRestartIndex func(index uint32)

var PrioritizeTextures func(n int32, textures *uint32, priorities *float32)

var ProgramBinary func(program uint32, binaryFormat uint32, binary unsafe.Pointer, length int32)

var ProgramParameteri func(program uint32, pname uint32, value int32)
var ProgramParameteriARB func(program uint32, pname uint32, value int32)
var ProgramParameteriEXT func(program uint32, pname uint32, value int32)
var ProgramPathFragmentInputGenNV func(program uint32, location int32, genMode uint32, components int32, coeffs *float32)
var ProgramUniform1d func(program uint32, location int32, v0 float64)
var ProgramUniform1dEXT func(program uint32, location int32, x float64)
var ProgramUniform1dv func(program uint32, location int32, count int32, value *float64)
var ProgramUniform1dvEXT func(program uint32, location int32, count int32, value *float64)

var ProgramUniform1f func(program uint32, location int32, v0 float32)
var ProgramUniform1fEXT func(program uint32, location int32, v0 float32)

var ProgramUniform1fv func(program uint32, location int32, count int32, value *float32)
var ProgramUniform1fvEXT func(program uint32, location int32, count int32, value *float32)

var ProgramUniform1i func(program uint32, location int32, v0 int32)
var ProgramUniform1i64ARB func(program uint32, location int32, x int64)
var ProgramUniform1i64NV func(program uint32, location int32, x int64)
var ProgramUniform1i64vARB func(program uint32, location int32, count int32, value *int64)
var ProgramUniform1i64vNV func(program uint32, location int32, count int32, value *int64)
var ProgramUniform1iEXT func(program uint32, location int32, v0 int32)

var ProgramUniform1iv func(program uint32, location int32, count int32, value *int32)
var ProgramUniform1ivEXT func(program uint32, location int32, count int32, value *int32)

var ProgramUniform1ui func(program uint32, location int32, v0 uint32)
var ProgramUniform1ui64ARB func(program uint32, location int32, x uint64)
var ProgramUniform1ui64NV func(program uint32, location int32, x uint64)
var ProgramUniform1ui64vARB func(program uint32, location int32, count int32, value *uint64)
var ProgramUniform1ui64vNV func(program uint32, location int32, count int32, value *uint64)
var ProgramUniform1uiEXT func(program uint32, location int32, v0 uint32)

var ProgramUniform1uiv func(program uint32, location int32, count int32, value *uint32)
var ProgramUniform1uivEXT func(program uint32, location int32, count int32, value *uint32)
var ProgramUniform2d func(program uint32, location int32, v0 float64, v1 float64)
var ProgramUniform2dEXT func(program uint32, location int32, x float64, y float64)
var ProgramUniform2dv func(program uint32, location int32, count int32, value *float64)
var ProgramUniform2dvEXT func(program uint32, location int32, count int32, value *float64)

var ProgramUniform2f func(program uint32, location int32, v0 float32, v1 float32)
var ProgramUniform2fEXT func(program uint32, location int32, v0 float32, v1 float32)

var ProgramUniform2fv func(program uint32, location int32, count int32, value *float32)
var ProgramUniform2fvEXT func(program uint32, location int32, count int32, value *float32)

var ProgramUniform2i func(program uint32, location int32, v0 int32, v1 int32)
var ProgramUniform2i64ARB func(program uint32, location int32, x int64, y int64)
var ProgramUniform2i64NV func(program uint32, location int32, x int64, y int64)
var ProgramUniform2i64vARB func(program uint32, location int32, count int32, value *int64)
var ProgramUniform2i64vNV func(program uint32, location int32, count int32, value *int64)
var ProgramUniform2iEXT func(program uint32, location int32, v0 int32, v1 int32)

var ProgramUniform2iv func(program uint32, location int32, count int32, value *int32)
var ProgramUniform2ivEXT func(program uint32, location int32, count int32, value *int32)

var ProgramUniform2ui func(program uint32, location int32, v0 uint32, v1 uint32)
var ProgramUniform2ui64ARB func(program uint32, location int32, x uint64, y uint64)
var ProgramUniform2ui64NV func(program uint32, location int32, x uint64, y uint64)
var ProgramUniform2ui64vARB func(program uint32, location int32, count int32, value *uint64)
var ProgramUniform2ui64vNV func(program uint32, location int32, count int32, value *uint64)
var ProgramUniform2uiEXT func(program uint32, location int32, v0 uint32, v1 uint32)

var ProgramUniform2uiv func(program uint32, location int32, count int32, value *uint32)
var ProgramUniform2uivEXT func(program uint32, location int32, count int32, value *uint32)
var ProgramUniform3d func(program uint32, location int32, v0 float64, v1 float64, v2 float64)
var ProgramUniform3dEXT func(program uint32, location int32, x float64, y float64, z float64)
var ProgramUniform3dv func(program uint32, location int32, count int32, value *float64)
var ProgramUniform3dvEXT func(program uint32, location int32, count int32, value *float64)

var ProgramUniform3f func(program uint32, location int32, v0 float32, v1 float32, v2 float32)
var ProgramUniform3fEXT func(program uint32, location int32, v0 float32, v1 float32, v2 float32)

var ProgramUniform3fv func(program uint32, location int32, count int32, value *float32)
var ProgramUniform3fvEXT func(program uint32, location int32, count int32, value *float32)

var ProgramUniform3i func(program uint32, location int32, v0 int32, v1 int32, v2 int32)
var ProgramUniform3i64ARB func(program uint32, location int32, x int64, y int64, z int64)
var ProgramUniform3i64NV func(program uint32, location int32, x int64, y int64, z int64)
var ProgramUniform3i64vARB func(program uint32, location int32, count int32, value *int64)
var ProgramUniform3i64vNV func(program uint32, location int32, count int32, value *int64)
var ProgramUniform3iEXT func(program uint32, location int32, v0 int32, v1 int32, v2 int32)

var ProgramUniform3iv func(program uint32, location int32, count int32, value *int32)
var ProgramUniform3ivEXT func(program uint32, location int32, count int32, value *int32)

var ProgramUniform3ui func(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32)
var ProgramUniform3ui64ARB func(program uint32, location int32, x uint64, y uint64, z uint64)
var ProgramUniform3ui64NV func(program uint32, location int32, x uint64, y uint64, z uint64)
var ProgramUniform3ui64vARB func(program uint32, location int32, count int32, value *uint64)
var ProgramUniform3ui64vNV func(program uint32, location int32, count int32, value *uint64)
var ProgramUniform3uiEXT func(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32)

var ProgramUniform3uiv func(program uint32, location int32, count int32, value *uint32)
var ProgramUniform3uivEXT func(program uint32, location int32, count int32, value *uint32)
var ProgramUniform4d func(program uint32, location int32, v0 float64, v1 float64, v2 float64, v3 float64)
var ProgramUniform4dEXT func(program uint32, location int32, x float64, y float64, z float64, w float64)
var ProgramUniform4dv func(program uint32, location int32, count int32, value *float64)
var ProgramUniform4dvEXT func(program uint32, location int32, count int32, value *float64)

var ProgramUniform4f func(program uint32, location int32, v0 float32, v1 float32, v2 float32, v3 float32)
var ProgramUniform4fEXT func(program uint32, location int32, v0 float32, v1 float32, v2 float32, v3 float32)

var ProgramUniform4fv func(program uint32, location int32, count int32, value *float32)
var ProgramUniform4fvEXT func(program uint32, location int32, count int32, value *float32)

var ProgramUniform4i func(program uint32, location int32, v0 int32, v1 int32, v2 int32, v3 int32)
var ProgramUniform4i64ARB func(program uint32, location int32, x int64, y int64, z int64, w int64)
var ProgramUniform4i64NV func(program uint32, location int32, x int64, y int64, z int64, w int64)
var ProgramUniform4i64vARB func(program uint32, location int32, count int32, value *int64)
var ProgramUniform4i64vNV func(program uint32, location int32, count int32, value *int64)
var ProgramUniform4iEXT func(program uint32, location int32, v0 int32, v1 int32, v2 int32, v3 int32)

var ProgramUniform4iv func(program uint32, location int32, count int32, value *int32)
var ProgramUniform4ivEXT func(program uint32, location int32, count int32, value *int32)

var ProgramUniform4ui func(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32)
var ProgramUniform4ui64ARB func(program uint32, location int32, x uint64, y uint64, z uint64, w uint64)
var ProgramUniform4ui64NV func(program uint32, location int32, x uint64, y uint64, z uint64, w uint64)
var ProgramUniform4ui64vARB func(program uint32, location int32, count int32, value *uint64)
var ProgramUniform4ui64vNV func(program uint32, location int32, count int32, value *uint64)
var ProgramUniform4uiEXT func(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32)

var ProgramUniform4uiv func(program uint32, location int32, count int32, value *uint32)
var ProgramUniform4uivEXT func(program uint32, location int32, count int32, value *uint32)
var ProgramUniformHandleui64ARB func(program uint32, location int32, value uint64)
var ProgramUniformHandleui64NV func(program uint32, location int32, value uint64)
var ProgramUniformHandleui64vARB func(program uint32, location int32, count int32, values *uint64)
var ProgramUniformHandleui64vNV func(program uint32, location int32, count int32, values *uint64)
var ProgramUniformMatrix2dv func(program uint32, location int32, count int32, transpose bool, value *float64)
var ProgramUniformMatrix2dvEXT func(program uint32, location int32, count int32, transpose bool, value *float64)

var ProgramUniformMatrix2fv func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix2fvEXT func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix2x3dv func(program uint32, location int32, count int32, transpose bool, value *float64)
var ProgramUniformMatrix2x3dvEXT func(program uint32, location int32, count int32, transpose bool, value *float64)

var ProgramUniformMatrix2x3fv func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix2x3fvEXT func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix2x4dv func(program uint32, location int32, count int32, transpose bool, value *float64)
var ProgramUniformMatrix2x4dvEXT func(program uint32, location int32, count int32, transpose bool, value *float64)

var ProgramUniformMatrix2x4fv func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix2x4fvEXT func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix3dv func(program uint32, location int32, count int32, transpose bool, value *float64)
var ProgramUniformMatrix3dvEXT func(program uint32, location int32, count int32, transpose bool, value *float64)

var ProgramUniformMatrix3fv func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix3fvEXT func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix3x2dv func(program uint32, location int32, count int32, transpose bool, value *float64)
var ProgramUniformMatrix3x2dvEXT func(program uint32, location int32, count int32, transpose bool, value *float64)

var ProgramUniformMatrix3x2fv func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix3x2fvEXT func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix3x4dv func(program uint32, location int32, count int32, transpose bool, value *float64)
var ProgramUniformMatrix3x4dvEXT func(program uint32, location int32, count int32, transpose bool, value *float64)

var ProgramUniformMatrix3x4fv func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix3x4fvEXT func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix4dv func(program uint32, location int32, count int32, transpose bool, value *float64)
var ProgramUniformMatrix4dvEXT func(program uint32, location int32, count int32, transpose bool, value *float64)

var ProgramUniformMatrix4fv func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix4fvEXT func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix4x2dv func(program uint32, location int32, count int32, transpose bool, value *float64)
var ProgramUniformMatrix4x2dvEXT func(program uint32, location int32, count int32, transpose bool, value *float64)

var ProgramUniformMatrix4x2fv func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix4x2fvEXT func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix4x3dv func(program uint32, location int32, count int32, transpose bool, value *float64)
var ProgramUniformMatrix4x3dvEXT func(program uint32, location int32, count int32, transpose bool, value *float64)

var ProgramUniformMatrix4x3fv func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformMatrix4x3fvEXT func(program uint32, location int32, count int32, transpose bool, value *float32)
var ProgramUniformui64NV func(program uint32, location int32, value uint64)
var ProgramUniformui64vNV func(program uint32, location int32, count int32, value *uint64)

var ProvokingVertex func(mode uint32)

var PushAttrib func(mask uint32)

var PushClientAttrib func(mask uint32)
var PushClientAttribDefaultEXT func(mask uint32)

var PushDebugGroup func(source uint32, id uint32, length int32, message *uint8)
var PushDebugGroupKHR func(source uint32, id uint32, length int32, message *uint8)
var PushGroupMarkerEXT func(length int32, marker *uint8)

var PushMatrix func()

var PushName func(name uint32)

var QueryCounter func(id uint32, target uint32)
var RasterPos2d func(x float64, y float64)
var RasterPos2dv func(v *float64)
var RasterPos2f func(x float32, y float32)
var RasterPos2fv func(v *float32)
var RasterPos2i func(x int32, y int32)
var RasterPos2iv func(v *int32)
var RasterPos2s func(x int16, y int16)
var RasterPos2sv func(v *int16)
var RasterPos3d func(x float64, y float64, z float64)
var RasterPos3dv func(v *float64)
var RasterPos3f func(x float32, y float32, z float32)
var RasterPos3fv func(v *float32)
var RasterPos3i func(x int32, y int32, z int32)
var RasterPos3iv func(v *int32)
var RasterPos3s func(x int16, y int16, z int16)
var RasterPos3sv func(v *int16)
var RasterPos4d func(x float64, y float64, z float64, w float64)
var RasterPos4dv func(v *float64)
var RasterPos4f func(x float32, y float32, z float32, w float32)
var RasterPos4fv func(v *float32)
var RasterPos4i func(x int32, y int32, z int32, w int32)
var RasterPos4iv func(v *int32)
var RasterPos4s func(x int16, y int16, z int16, w int16)
var RasterPos4sv func(v *int16)
var RasterSamplesEXT func(samples uint32, fixedsamplelocations bool)

var ReadBuffer func(src uint32)

var ReadPixels func(x int32, y int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var ReadnPixels func(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer)
var ReadnPixelsARB func(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer)
var ReadnPixelsKHR func(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer)
var Rectd func(x1 float64, y1 float64, x2 float64, y2 float64)
var Rectdv func(v1 *float64, v2 *float64)
var Rectf func(x1 float32, y1 float32, x2 float32, y2 float32)
var Rectfv func(v1 *float32, v2 *float32)
var Recti func(x1 int32, y1 int32, x2 int32, y2 int32)
var Rectiv func(v1 *int32, v2 *int32)
var Rects func(x1 int16, y1 int16, x2 int16, y2 int16)
var Rectsv func(v1 *int16, v2 *int16)

var ReleaseShaderCompiler func()

var RenderMode func(mode uint32) int32

var RenderbufferStorage func(target uint32, internalformat uint32, width int32, height int32)

var RenderbufferStorageMultisample func(target uint32, samples int32, internalformat uint32, width int32, height int32)
var RenderbufferStorageMultisampleCoverageNV func(target uint32, coverageSamples int32, colorSamples int32, internalformat uint32, width int32, height int32)
var ResolveDepthValuesNV func()

var ResumeTransformFeedback func()
var Rotated func(angle float64, x float64, y float64, z float64)
var Rotatef func(angle float32, x float32, y float32, z float32)

var SampleCoverage func(value float32, invert bool)

var SampleMaski func(maskNumber uint32, mask uint32)
var SamplerParameterIiv func(sampler uint32, pname uint32, param *int32)
var SamplerParameterIuiv func(sampler uint32, pname uint32, param *uint32)
var SamplerParameterf func(sampler uint32, pname uint32, param float32)
var SamplerParameterfv func(sampler uint32, pname uint32, param *float32)
var SamplerParameteri func(sampler uint32, pname uint32, param int32)
var SamplerParameteriv func(sampler uint32, pname uint32, param *int32)
var Scaled func(x float64, y float64, z float64)
var Scalef func(x float32, y float32, z float32)

var Scissor func(x int32, y int32, width int32, height int32)
var ScissorArrayv func(first uint32, count int32, v *int32)

var ScissorIndexed func(index uint32, left int32, bottom int32, width int32, height int32)
var ScissorIndexedv func(index uint32, v *int32)
var SecondaryColor3b func(red int8, green int8, blue int8)
var SecondaryColor3bv func(v *int8)
var SecondaryColor3d func(red float64, green float64, blue float64)
var SecondaryColor3dv func(v *float64)
var SecondaryColor3f func(red float32, green float32, blue float32)
var SecondaryColor3fv func(v *float32)
var SecondaryColor3i func(red int32, green int32, blue int32)
var SecondaryColor3iv func(v *int32)
var SecondaryColor3s func(red int16, green int16, blue int16)
var SecondaryColor3sv func(v *int16)
var SecondaryColor3ub func(red uint8, green uint8, blue uint8)
var SecondaryColor3ubv func(v *uint8)
var SecondaryColor3ui func(red uint32, green uint32, blue uint32)
var SecondaryColor3uiv func(v *uint32)
var SecondaryColor3us func(red uint16, green uint16, blue uint16)
var SecondaryColor3usv func(v *uint16)
var SecondaryColorFormatNV func(size int32, xtype uint32, stride int32)

var SecondaryColorPointer func(size int32, xtype uint32, stride int32, pointer unsafe.Pointer)

var SelectBuffer func(size int32, buffer *uint32)
var SelectPerfMonitorCountersAMD func(monitor uint32, enable bool, group uint32, numCounters int32, counterList *uint32)

var ShadeModel func(mode uint32)

var ShaderBinary func(count int32, shaders *uint32, binaryformat uint32, binary unsafe.Pointer, length int32)

var ShaderSource func(shader uint32, count int32, xstring **uint8, length *int32)

var ShaderStorageBlockBinding func(program uint32, storageBlockIndex uint32, storageBlockBinding uint32)
var SignalVkFenceNV func(vkFence uint64)
var SignalVkSemaphoreNV func(vkSemaphore uint64)
var SpecializeShader func(shader uint32, pEntryPoint *uint8, numSpecializationConstants uint32, pConstantIndex *uint32, pConstantValue *uint32)
var SpecializeShaderARB func(shader uint32, pEntryPoint *uint8, numSpecializationConstants uint32, pConstantIndex *uint32, pConstantValue *uint32)
var StateCaptureNV func(state uint32, mode uint32)
var StencilFillPathInstancedNV func(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, fillMode uint32, mask uint32, transformType uint32, transformValues *float32)
var StencilFillPathNV func(path uint32, fillMode uint32, mask uint32)

var StencilFunc func(xfunc uint32, ref int32, mask uint32)

var StencilFuncSeparate func(face uint32, xfunc uint32, ref int32, mask uint32)

var StencilMask func(mask uint32)

var StencilMaskSeparate func(face uint32, mask uint32)

var StencilOp func(fail uint32, zfail uint32, zpass uint32)

var StencilOpSeparate func(face uint32, sfail uint32, dpfail uint32, dppass uint32)
var StencilStrokePathInstancedNV func(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, reference int32, mask uint32, transformType uint32, transformValues *float32)
var StencilStrokePathNV func(path uint32, reference int32, mask uint32)
var StencilThenCoverFillPathInstancedNV func(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, fillMode uint32, mask uint32, coverMode uint32, transformType uint32, transformValues *float32)
var StencilThenCoverFillPathNV func(path uint32, fillMode uint32, mask uint32, coverMode uint32)
var StencilThenCoverStrokePathInstancedNV func(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, reference int32, mask uint32, coverMode uint32, transformType uint32, transformValues *float32)
var StencilThenCoverStrokePathNV func(path uint32, reference int32, mask uint32, coverMode uint32)
var SubpixelPrecisionBiasNV func(xbits uint32, ybits uint32)

var TexBuffer func(target uint32, internalformat uint32, buffer uint32)
var TexBufferARB func(target uint32, internalformat uint32, buffer uint32)

var TexBufferRange func(target uint32, internalformat uint32, buffer uint32, offset int, size int)
var TexCoord1d func(s float64)
var TexCoord1dv func(v *float64)
var TexCoord1f func(s float32)
var TexCoord1fv func(v *float32)
var TexCoord1i func(s int32)
var TexCoord1iv func(v *int32)
var TexCoord1s func(s int16)
var TexCoord1sv func(v *int16)
var TexCoord2d func(s float64, t float64)
var TexCoord2dv func(v *float64)
var TexCoord2f func(s float32, t float32)
var TexCoord2fv func(v *float32)
var TexCoord2i func(s int32, t int32)
var TexCoord2iv func(v *int32)
var TexCoord2s func(s int16, t int16)
var TexCoord2sv func(v *int16)
var TexCoord3d func(s float64, t float64, r float64)
var TexCoord3dv func(v *float64)
var TexCoord3f func(s float32, t float32, r float32)
var TexCoord3fv func(v *float32)
var TexCoord3i func(s int32, t int32, r int32)
var TexCoord3iv func(v *int32)
var TexCoord3s func(s int16, t int16, r int16)
var TexCoord3sv func(v *int16)
var TexCoord4d func(s float64, t float64, r float64, q float64)
var TexCoord4dv func(v *float64)
var TexCoord4f func(s float32, t float32, r float32, q float32)
var TexCoord4fv func(v *float32)
var TexCoord4i func(s int32, t int32, r int32, q int32)
var TexCoord4iv func(v *int32)
var TexCoord4s func(s int16, t int16, r int16, q int16)
var TexCoord4sv func(v *int16)
var TexCoordFormatNV func(size int32, xtype uint32, stride int32)

var TexCoordPointer func(size int32, xtype uint32, stride int32, pointer unsafe.Pointer)
var TexEnvf func(target uint32, pname uint32, param float32)
var TexEnvfv func(target uint32, pname uint32, params *float32)
var TexEnvi func(target uint32, pname uint32, param int32)
var TexEnviv func(target uint32, pname uint32, params *int32)
var TexGend func(coord uint32, pname uint32, param float64)
var TexGendv func(coord uint32, pname uint32, params *float64)
var TexGenf func(coord uint32, pname uint32, param float32)
var TexGenfv func(coord uint32, pname uint32, params *float32)
var TexGeni func(coord uint32, pname uint32, param int32)
var TexGeniv func(coord uint32, pname uint32, params *int32)

var TexImage1D func(target uint32, level int32, internalformat int32, width int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var TexImage2D func(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var TexImage2DMultisample func(target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool)

var TexImage3D func(target uint32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var TexImage3DMultisample func(target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool)
var TexPageCommitmentARB func(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, commit bool)
var TexParameterIiv func(target uint32, pname uint32, params *int32)
var TexParameterIuiv func(target uint32, pname uint32, params *uint32)
var TexParameterf func(target uint32, pname uint32, param float32)
var TexParameterfv func(target uint32, pname uint32, params *float32)
var TexParameteri func(target uint32, pname uint32, param int32)
var TexParameteriv func(target uint32, pname uint32, params *int32)

var TexStorage1D func(target uint32, levels int32, internalformat uint32, width int32)

var TexStorage2D func(target uint32, levels int32, internalformat uint32, width int32, height int32)

var TexStorage2DMultisample func(target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool)

var TexStorage3D func(target uint32, levels int32, internalformat uint32, width int32, height int32, depth int32)

var TexStorage3DMultisample func(target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool)

var TexSubImage1D func(target uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var TexSubImage2D func(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var TexSubImage3D func(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var TextureBarrier func()
var TextureBarrierNV func()

var TextureBuffer func(texture uint32, internalformat uint32, buffer uint32)
var TextureBufferEXT func(texture uint32, target uint32, internalformat uint32, buffer uint32)

var TextureBufferRange func(texture uint32, internalformat uint32, buffer uint32, offset int, size int)
var TextureBufferRangeEXT func(texture uint32, target uint32, internalformat uint32, buffer uint32, offset int, size int)
var TextureImage1DEXT func(texture uint32, target uint32, level int32, internalformat int32, width int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var TextureImage2DEXT func(texture uint32, target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var TextureImage3DEXT func(texture uint32, target uint32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var TexturePageCommitmentEXT func(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, commit bool)
var TextureParameterIiv func(texture uint32, pname uint32, params *int32)
var TextureParameterIivEXT func(texture uint32, target uint32, pname uint32, params *int32)
var TextureParameterIuiv func(texture uint32, pname uint32, params *uint32)
var TextureParameterIuivEXT func(texture uint32, target uint32, pname uint32, params *uint32)
var TextureParameterf func(texture uint32, pname uint32, param float32)
var TextureParameterfEXT func(texture uint32, target uint32, pname uint32, param float32)
var TextureParameterfv func(texture uint32, pname uint32, param *float32)
var TextureParameterfvEXT func(texture uint32, target uint32, pname uint32, params *float32)
var TextureParameteri func(texture uint32, pname uint32, param int32)
var TextureParameteriEXT func(texture uint32, target uint32, pname uint32, param int32)
var TextureParameteriv func(texture uint32, pname uint32, param *int32)
var TextureParameterivEXT func(texture uint32, target uint32, pname uint32, params *int32)
var TextureRenderbufferEXT func(texture uint32, target uint32, renderbuffer uint32)

var TextureStorage1D func(texture uint32, levels int32, internalformat uint32, width int32)
var TextureStorage1DEXT func(texture uint32, target uint32, levels int32, internalformat uint32, width int32)

var TextureStorage2D func(texture uint32, levels int32, internalformat uint32, width int32, height int32)
var TextureStorage2DEXT func(texture uint32, target uint32, levels int32, internalformat uint32, width int32, height int32)

var TextureStorage2DMultisample func(texture uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool)
var TextureStorage2DMultisampleEXT func(texture uint32, target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool)

var TextureStorage3D func(texture uint32, levels int32, internalformat uint32, width int32, height int32, depth int32)
var TextureStorage3DEXT func(texture uint32, target uint32, levels int32, internalformat uint32, width int32, height int32, depth int32)

var TextureStorage3DMultisample func(texture uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool)
var TextureStorage3DMultisampleEXT func(texture uint32, target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool)

var TextureSubImage1D func(texture uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var TextureSubImage1DEXT func(texture uint32, target uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var TextureSubImage2D func(texture uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var TextureSubImage2DEXT func(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var TextureSubImage3D func(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var TextureSubImage3DEXT func(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer)

var TextureView func(texture uint32, target uint32, origtexture uint32, internalformat uint32, minlevel uint32, numlevels uint32, minlayer uint32, numlayers uint32)

var TransformFeedbackBufferBase func(xfb uint32, index uint32, buffer uint32)

var TransformFeedbackBufferRange func(xfb uint32, index uint32, buffer uint32, offset int, size int)

var TransformFeedbackVaryings func(program uint32, count int32, varyings **uint8, bufferMode uint32)
var TransformPathNV func(resultPath uint32, srcPath uint32, transformType uint32, transformValues *float32)
var Translated func(x float64, y float64, z float64)
var Translatef func(x float32, y float32, z float32)
var Uniform1d func(location int32, x float64)
var Uniform1dv func(location int32, count int32, value *float64)

var Uniform1f func(location int32, v0 float32)

var Uniform1fv func(location int32, count int32, value *float32)

var Uniform1i func(location int32, v0 int32)
var Uniform1i64ARB func(location int32, x int64)
var Uniform1i64NV func(location int32, x int64)
var Uniform1i64vARB func(location int32, count int32, value *int64)
var Uniform1i64vNV func(location int32, count int32, value *int64)

var Uniform1iv func(location int32, count int32, value *int32)

var Uniform1ui func(location int32, v0 uint32)
var Uniform1ui64ARB func(location int32, x uint64)
var Uniform1ui64NV func(location int32, x uint64)
var Uniform1ui64vARB func(location int32, count int32, value *uint64)
var Uniform1ui64vNV func(location int32, count int32, value *uint64)

var Uniform1uiv func(location int32, count int32, value *uint32)
var Uniform2d func(location int32, x float64, y float64)
var Uniform2dv func(location int32, count int32, value *float64)

var Uniform2f func(location int32, v0 float32, v1 float32)

var Uniform2fv func(location int32, count int32, value *float32)

var Uniform2i func(location int32, v0 int32, v1 int32)
var Uniform2i64ARB func(location int32, x int64, y int64)
var Uniform2i64NV func(location int32, x int64, y int64)
var Uniform2i64vARB func(location int32, count int32, value *int64)
var Uniform2i64vNV func(location int32, count int32, value *int64)

var Uniform2iv func(location int32, count int32, value *int32)

var Uniform2ui func(location int32, v0 uint32, v1 uint32)
var Uniform2ui64ARB func(location int32, x uint64, y uint64)
var Uniform2ui64NV func(location int32, x uint64, y uint64)
var Uniform2ui64vARB func(location int32, count int32, value *uint64)
var Uniform2ui64vNV func(location int32, count int32, value *uint64)

var Uniform2uiv func(location int32, count int32, value *uint32)
var Uniform3d func(location int32, x float64, y float64, z float64)
var Uniform3dv func(location int32, count int32, value *float64)

var Uniform3f func(location int32, v0 float32, v1 float32, v2 float32)

var Uniform3fv func(location int32, count int32, value *float32)

var Uniform3i func(location int32, v0 int32, v1 int32, v2 int32)
var Uniform3i64ARB func(location int32, x int64, y int64, z int64)
var Uniform3i64NV func(location int32, x int64, y int64, z int64)
var Uniform3i64vARB func(location int32, count int32, value *int64)
var Uniform3i64vNV func(location int32, count int32, value *int64)

var Uniform3iv func(location int32, count int32, value *int32)

var Uniform3ui func(location int32, v0 uint32, v1 uint32, v2 uint32)
var Uniform3ui64ARB func(location int32, x uint64, y uint64, z uint64)
var Uniform3ui64NV func(location int32, x uint64, y uint64, z uint64)
var Uniform3ui64vARB func(location int32, count int32, value *uint64)
var Uniform3ui64vNV func(location int32, count int32, value *uint64)

var Uniform3uiv func(location int32, count int32, value *uint32)
var Uniform4d func(location int32, x float64, y float64, z float64, w float64)
var Uniform4dv func(location int32, count int32, value *float64)

var Uniform4f func(location int32, v0 float32, v1 float32, v2 float32, v3 float32)

var Uniform4fv func(location int32, count int32, value *float32)

var Uniform4i func(location int32, v0 int32, v1 int32, v2 int32, v3 int32)
var Uniform4i64ARB func(location int32, x int64, y int64, z int64, w int64)
var Uniform4i64NV func(location int32, x int64, y int64, z int64, w int64)
var Uniform4i64vARB func(location int32, count int32, value *int64)
var Uniform4i64vNV func(location int32, count int32, value *int64)

var Uniform4iv func(location int32, count int32, value *int32)

var Uniform4ui func(location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32)
var Uniform4ui64ARB func(location int32, x uint64, y uint64, z uint64, w uint64)
var Uniform4ui64NV func(location int32, x uint64, y uint64, z uint64, w uint64)
var Uniform4ui64vARB func(location int32, count int32, value *uint64)
var Uniform4ui64vNV func(location int32, count int32, value *uint64)

var Uniform4uiv func(location int32, count int32, value *uint32)

var UniformBlockBinding func(program uint32, uniformBlockIndex uint32, uniformBlockBinding uint32)
var UniformHandleui64ARB func(location int32, value uint64)
var UniformHandleui64NV func(location int32, value uint64)
var UniformHandleui64vARB func(location int32, count int32, value *uint64)
var UniformHandleui64vNV func(location int32, count int32, value *uint64)
var UniformMatrix2dv func(location int32, count int32, transpose bool, value *float64)

var UniformMatrix2fv func(location int32, count int32, transpose bool, value *float32)
var UniformMatrix2x3dv func(location int32, count int32, transpose bool, value *float64)

var UniformMatrix2x3fv func(location int32, count int32, transpose bool, value *float32)
var UniformMatrix2x4dv func(location int32, count int32, transpose bool, value *float64)

var UniformMatrix2x4fv func(location int32, count int32, transpose bool, value *float32)
var UniformMatrix3dv func(location int32, count int32, transpose bool, value *float64)

var UniformMatrix3fv func(location int32, count int32, transpose bool, value *float32)
var UniformMatrix3x2dv func(location int32, count int32, transpose bool, value *float64)

var UniformMatrix3x2fv func(location int32, count int32, transpose bool, value *float32)
var UniformMatrix3x4dv func(location int32, count int32, transpose bool, value *float64)

var UniformMatrix3x4fv func(location int32, count int32, transpose bool, value *float32)
var UniformMatrix4dv func(location int32, count int32, transpose bool, value *float64)

var UniformMatrix4fv func(location int32, count int32, transpose bool, value *float32)
var UniformMatrix4x2dv func(location int32, count int32, transpose bool, value *float64)

var UniformMatrix4x2fv func(location int32, count int32, transpose bool, value *float32)
var UniformMatrix4x3dv func(location int32, count int32, transpose bool, value *float64)

var UniformMatrix4x3fv func(location int32, count int32, transpose bool, value *float32)
var UniformSubroutinesuiv func(shadertype uint32, count int32, indices *uint32)
var Uniformui64NV func(location int32, value uint64)
var Uniformui64vNV func(location int32, count int32, value *uint64)

var UnmapBuffer func(target uint32) bool

var UnmapNamedBuffer func(buffer uint32) bool
var UnmapNamedBufferEXT func(buffer uint32) bool

var UseProgram func(program uint32)

var UseProgramStages func(pipeline uint32, stages uint32, program uint32)
var UseProgramStagesEXT func(pipeline uint32, stages uint32, program uint32)
var UseShaderProgramEXT func(xtype uint32, program uint32)

var ValidateProgram func(program uint32)

var ValidateProgramPipeline func(pipeline uint32)
var ValidateProgramPipelineEXT func(pipeline uint32)
var Vertex2d func(x float64, y float64)
var Vertex2dv func(v *float64)
var Vertex2f func(x float32, y float32)
var Vertex2fv func(v *float32)
var Vertex2i func(x int32, y int32)
var Vertex2iv func(v *int32)
var Vertex2s func(x int16, y int16)
var Vertex2sv func(v *int16)
var Vertex3d func(x float64, y float64, z float64)
var Vertex3dv func(v *float64)
var Vertex3f func(x float32, y float32, z float32)
var Vertex3fv func(v *float32)
var Vertex3i func(x int32, y int32, z int32)
var Vertex3iv func(v *int32)
var Vertex3s func(x int16, y int16, z int16)
var Vertex3sv func(v *int16)
var Vertex4d func(x float64, y float64, z float64, w float64)
var Vertex4dv func(v *float64)
var Vertex4f func(x float32, y float32, z float32, w float32)
var Vertex4fv func(v *float32)
var Vertex4i func(x int32, y int32, z int32, w int32)
var Vertex4iv func(v *int32)
var Vertex4s func(x int16, y int16, z int16, w int16)
var Vertex4sv func(v *int16)
var VertexArrayAttribBinding func(vaobj uint32, attribindex uint32, bindingindex uint32)

var VertexArrayAttribFormat func(vaobj uint32, attribindex uint32, size int32, xtype uint32, normalized bool, relativeoffset uint32)
var VertexArrayAttribIFormat func(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32)
var VertexArrayAttribLFormat func(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32)
var VertexArrayBindVertexBufferEXT func(vaobj uint32, bindingindex uint32, buffer uint32, offset int, stride int32)

var VertexArrayBindingDivisor func(vaobj uint32, bindingindex uint32, divisor uint32)
var VertexArrayColorOffsetEXT func(vaobj uint32, buffer uint32, size int32, xtype uint32, stride int32, offset int)
var VertexArrayEdgeFlagOffsetEXT func(vaobj uint32, buffer uint32, stride int32, offset int)

var VertexArrayElementBuffer func(vaobj uint32, buffer uint32)
var VertexArrayFogCoordOffsetEXT func(vaobj uint32, buffer uint32, xtype uint32, stride int32, offset int)
var VertexArrayIndexOffsetEXT func(vaobj uint32, buffer uint32, xtype uint32, stride int32, offset int)
var VertexArrayMultiTexCoordOffsetEXT func(vaobj uint32, buffer uint32, texunit uint32, size int32, xtype uint32, stride int32, offset int)
var VertexArrayNormalOffsetEXT func(vaobj uint32, buffer uint32, xtype uint32, stride int32, offset int)
var VertexArraySecondaryColorOffsetEXT func(vaobj uint32, buffer uint32, size int32, xtype uint32, stride int32, offset int)
var VertexArrayTexCoordOffsetEXT func(vaobj uint32, buffer uint32, size int32, xtype uint32, stride int32, offset int)
var VertexArrayVertexAttribBindingEXT func(vaobj uint32, attribindex uint32, bindingindex uint32)
var VertexArrayVertexAttribDivisorEXT func(vaobj uint32, index uint32, divisor uint32)
var VertexArrayVertexAttribFormatEXT func(vaobj uint32, attribindex uint32, size int32, xtype uint32, normalized bool, relativeoffset uint32)
var VertexArrayVertexAttribIFormatEXT func(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32)
var VertexArrayVertexAttribIOffsetEXT func(vaobj uint32, buffer uint32, index uint32, size int32, xtype uint32, stride int32, offset int)
var VertexArrayVertexAttribLFormatEXT func(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32)
var VertexArrayVertexAttribLOffsetEXT func(vaobj uint32, buffer uint32, index uint32, size int32, xtype uint32, stride int32, offset int)
var VertexArrayVertexAttribOffsetEXT func(vaobj uint32, buffer uint32, index uint32, size int32, xtype uint32, normalized bool, stride int32, offset int)
var VertexArrayVertexBindingDivisorEXT func(vaobj uint32, bindingindex uint32, divisor uint32)

var VertexArrayVertexBuffer func(vaobj uint32, bindingindex uint32, buffer uint32, offset int, stride int32)

var VertexArrayVertexBuffers func(vaobj uint32, first uint32, count int32, buffers *uint32, offsets *int, strides *int32)
var VertexArrayVertexOffsetEXT func(vaobj uint32, buffer uint32, size int32, xtype uint32, stride int32, offset int)
var VertexAttrib1d func(index uint32, x float64)
var VertexAttrib1dv func(index uint32, v *float64)
var VertexAttrib1f func(index uint32, x float32)
var VertexAttrib1fv func(index uint32, v *float32)
var VertexAttrib1s func(index uint32, x int16)
var VertexAttrib1sv func(index uint32, v *int16)
var VertexAttrib2d func(index uint32, x float64, y float64)
var VertexAttrib2dv func(index uint32, v *float64)
var VertexAttrib2f func(index uint32, x float32, y float32)
var VertexAttrib2fv func(index uint32, v *float32)
var VertexAttrib2s func(index uint32, x int16, y int16)
var VertexAttrib2sv func(index uint32, v *int16)
var VertexAttrib3d func(index uint32, x float64, y float64, z float64)
var VertexAttrib3dv func(index uint32, v *float64)
var VertexAttrib3f func(index uint32, x float32, y float32, z float32)
var VertexAttrib3fv func(index uint32, v *float32)
var VertexAttrib3s func(index uint32, x int16, y int16, z int16)
var VertexAttrib3sv func(index uint32, v *int16)
var VertexAttrib4Nbv func(index uint32, v *int8)
var VertexAttrib4Niv func(index uint32, v *int32)
var VertexAttrib4Nsv func(index uint32, v *int16)
var VertexAttrib4Nub func(index uint32, x uint8, y uint8, z uint8, w uint8)
var VertexAttrib4Nubv func(index uint32, v *uint8)
var VertexAttrib4Nuiv func(index uint32, v *uint32)
var VertexAttrib4Nusv func(index uint32, v *uint16)
var VertexAttrib4bv func(index uint32, v *int8)
var VertexAttrib4d func(index uint32, x float64, y float64, z float64, w float64)
var VertexAttrib4dv func(index uint32, v *float64)
var VertexAttrib4f func(index uint32, x float32, y float32, z float32, w float32)
var VertexAttrib4fv func(index uint32, v *float32)
var VertexAttrib4iv func(index uint32, v *int32)
var VertexAttrib4s func(index uint32, x int16, y int16, z int16, w int16)
var VertexAttrib4sv func(index uint32, v *int16)
var VertexAttrib4ubv func(index uint32, v *uint8)
var VertexAttrib4uiv func(index uint32, v *uint32)
var VertexAttrib4usv func(index uint32, v *uint16)

var VertexAttribBinding func(attribindex uint32, bindingindex uint32)

var VertexAttribDivisor func(index uint32, divisor uint32)
var VertexAttribDivisorARB func(index uint32, divisor uint32)

var VertexAttribFormat func(attribindex uint32, size int32, xtype uint32, normalized bool, relativeoffset uint32)
var VertexAttribFormatNV func(index uint32, size int32, xtype uint32, normalized bool, stride int32)
var VertexAttribI1i func(index uint32, x int32)
var VertexAttribI1iv func(index uint32, v *int32)
var VertexAttribI1ui func(index uint32, x uint32)
var VertexAttribI1uiv func(index uint32, v *uint32)
var VertexAttribI2i func(index uint32, x int32, y int32)
var VertexAttribI2iv func(index uint32, v *int32)
var VertexAttribI2ui func(index uint32, x uint32, y uint32)
var VertexAttribI2uiv func(index uint32, v *uint32)
var VertexAttribI3i func(index uint32, x int32, y int32, z int32)
var VertexAttribI3iv func(index uint32, v *int32)
var VertexAttribI3ui func(index uint32, x uint32, y uint32, z uint32)
var VertexAttribI3uiv func(index uint32, v *uint32)
var VertexAttribI4bv func(index uint32, v *int8)
var VertexAttribI4i func(index uint32, x int32, y int32, z int32, w int32)
var VertexAttribI4iv func(index uint32, v *int32)
var VertexAttribI4sv func(index uint32, v *int16)
var VertexAttribI4ubv func(index uint32, v *uint8)
var VertexAttribI4ui func(index uint32, x uint32, y uint32, z uint32, w uint32)
var VertexAttribI4uiv func(index uint32, v *uint32)
var VertexAttribI4usv func(index uint32, v *uint16)
var VertexAttribIFormat func(attribindex uint32, size int32, xtype uint32, relativeoffset uint32)
var VertexAttribIFormatNV func(index uint32, size int32, xtype uint32, stride int32)
var VertexAttribIPointer func(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer)
var VertexAttribL1d func(index uint32, x float64)
var VertexAttribL1dv func(index uint32, v *float64)
var VertexAttribL1i64NV func(index uint32, x int64)
var VertexAttribL1i64vNV func(index uint32, v *int64)
var VertexAttribL1ui64ARB func(index uint32, x uint64)
var VertexAttribL1ui64NV func(index uint32, x uint64)
var VertexAttribL1ui64vARB func(index uint32, v *uint64)
var VertexAttribL1ui64vNV func(index uint32, v *uint64)
var VertexAttribL2d func(index uint32, x float64, y float64)
var VertexAttribL2dv func(index uint32, v *float64)
var VertexAttribL2i64NV func(index uint32, x int64, y int64)
var VertexAttribL2i64vNV func(index uint32, v *int64)
var VertexAttribL2ui64NV func(index uint32, x uint64, y uint64)
var VertexAttribL2ui64vNV func(index uint32, v *uint64)
var VertexAttribL3d func(index uint32, x float64, y float64, z float64)
var VertexAttribL3dv func(index uint32, v *float64)
var VertexAttribL3i64NV func(index uint32, x int64, y int64, z int64)
var VertexAttribL3i64vNV func(index uint32, v *int64)
var VertexAttribL3ui64NV func(index uint32, x uint64, y uint64, z uint64)
var VertexAttribL3ui64vNV func(index uint32, v *uint64)
var VertexAttribL4d func(index uint32, x float64, y float64, z float64, w float64)
var VertexAttribL4dv func(index uint32, v *float64)
var VertexAttribL4i64NV func(index uint32, x int64, y int64, z int64, w int64)
var VertexAttribL4i64vNV func(index uint32, v *int64)
var VertexAttribL4ui64NV func(index uint32, x uint64, y uint64, z uint64, w uint64)
var VertexAttribL4ui64vNV func(index uint32, v *uint64)
var VertexAttribLFormat func(attribindex uint32, size int32, xtype uint32, relativeoffset uint32)
var VertexAttribLFormatNV func(index uint32, size int32, xtype uint32, stride int32)
var VertexAttribLPointer func(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer)
var VertexAttribP1ui func(index uint32, xtype uint32, normalized bool, value uint32)
var VertexAttribP1uiv func(index uint32, xtype uint32, normalized bool, value *uint32)
var VertexAttribP2ui func(index uint32, xtype uint32, normalized bool, value uint32)
var VertexAttribP2uiv func(index uint32, xtype uint32, normalized bool, value *uint32)
var VertexAttribP3ui func(index uint32, xtype uint32, normalized bool, value uint32)
var VertexAttribP3uiv func(index uint32, xtype uint32, normalized bool, value *uint32)
var VertexAttribP4ui func(index uint32, xtype uint32, normalized bool, value uint32)
var VertexAttribP4uiv func(index uint32, xtype uint32, normalized bool, value *uint32)

var VertexAttribPointer func(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer)

var VertexBindingDivisor func(bindingindex uint32, divisor uint32)
var VertexFormatNV func(size int32, xtype uint32, stride int32)

var VertexPointer func(size int32, xtype uint32, stride int32, pointer unsafe.Pointer)

var Viewport func(x int32, y int32, width int32, height int32)
var ViewportArrayv func(first uint32, count int32, v *float32)
var ViewportIndexedf func(index uint32, x float32, y float32, w float32, h float32)
var ViewportIndexedfv func(index uint32, v *float32)
var ViewportPositionWScaleNV func(index uint32, xcoeff float32, ycoeff float32)
var ViewportSwizzleNV func(index uint32, swizzlex uint32, swizzley uint32, swizzlez uint32, swizzlew uint32)

var WaitSync func(sync uintptr, flags uint32, timeout uint64)
var WaitVkSemaphoreNV func(vkSemaphore uint64)
var WeightPathsNV func(resultPath uint32, numPaths int32, paths *uint32, weights *float32)
var WindowPos2d func(x float64, y float64)
var WindowPos2dv func(v *float64)
var WindowPos2f func(x float32, y float32)
var WindowPos2fv func(v *float32)
var WindowPos2i func(x int32, y int32)
var WindowPos2iv func(v *int32)
var WindowPos2s func(x int16, y int16)
var WindowPos2sv func(v *int16)
var WindowPos3d func(x float64, y float64, z float64)
var WindowPos3dv func(v *float64)
var WindowPos3f func(x float32, y float32, z float32)
var WindowPos3fv func(v *float32)
var WindowPos3i func(x int32, y int32, z int32)
var WindowPos3iv func(v *int32)
var WindowPos3s func(x int16, y int16, z int16)
var WindowPos3sv func(v *int16)
var WindowRectanglesEXT func(mode uint32, count int32, box *int32)
