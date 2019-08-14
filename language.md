Language handling is contained within `pkg/gvdl/gvdl.go`.

```
add source {name} {path}
# Adds a source with the given name and filepath.
# add source default0 "sample0.mp4"


add layer {name} {program-name} {depth} {source-name0} {source-name1} {source-name2}
# Adds a layer, higher depths show up on top (I think)
# add layer default default 1.0 default0 default1 default2


add program {name} {vertex-shader-path} {fragment-shader-path}
# Adds a program with the shaders. Currently only the name program named "default" is supported.
# add program default "shaders/vert/default.glsl" "shaders/frag/rave.glsl"


set uniform {type} {name} {value} {layers}
# Sets the uniforms to use for the layers. Layers is a list of space-separated layer names, no grouping symbols
# set uniform float amplitude "1.4" "default" "swap" "window"


reload programs
# Reloads all shaders
```

On daemon startup, the following commands are run:
```
add program default "shaders/vert/default.glsl" "shaders/frag/default.glsl"
add source default0 "sample0.mp4"
add source default1 "sample1.mp4"
add source default2 "sample2.mp4"
add layer default default 0 default0 default1 default2
```

Quotation marks are understood, and are required for most filepaths and numeric values with decimals.

Additional commands can be run at runtime using netcat. By default the daemon listens at `localhost:4200`. This can be changed by starting the daemon with `-net=":8765"`.

On OSX this looks like:
```
dianes-air:gvd dianelooney$ nc localhost 4200
add program default "shaders/vert/default.glsl" "shaders/frag/rave.glsl"
add source default "sample2.mp4"
add layer default default 1.0 default0 default1 default2
```
