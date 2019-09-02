# OSC
This document will provide a reference for all OSC calls
Anything in angle brackets (< and >) is to be replaced with the relevant values.

#### Add Video
`/source.ffvideo/create ,ss <name for video> <path to file>`

#### Add Static Shader Program
`/program/create ,sss <name for the program> <vert shader path> <frag shader path>`

#### Add Watched Shader Program
`/program/watch ,sss <name for the program> <vert shader path> <frag shader path>`

#### Add Shader Node
`/source.shader/create ,s <name you loaded the program under>`
Will default to use the program with the same name as itself

#### Set Shader Program
`/source.shader/set/program ,sis <name of shader source> <program name>`

#### Set a texture for a Shader
`/source.shader/set/input ,sis <name of shader source> <index> <name of source node>`

#### Set Wrap Mode for a 
s is the X wrap behavior
`/source/set/wrap.s ,ss <source name> <wrap mode>`
t is the Y wrap behavior
`/source/set/wrap.t ,ss <source name> <wrap mode>`

Wrap mode is a string from the following:
- CLAMP_TO_EDGE
- CLAMP_TO_BORDER
- MIRRORED_REPEAT
- REPEAT
- MIRROR_CLAMP_TO_EDGE

#### Set Min and Mag
`/source/set/magfilter ,ss <source name> <filter type>`
`/source/set/maxfilter ,ss <source name> <filter type>`

Filter type is one of the following:
- NEAREST
- LINEAR


#### Set a Uniform float
`/source.shader/set/uniform1f ,ssf <shader name> <uniform name> <value to send>`


