Language handling is contained within `pkg/gvdl/gvdl.go`.

```
add source {name} {path} # Adds a source with the given name and filepath.
add layer {name} {source-name} {program-name} {depth} # Adds a layer, higher depths show up on top (I think)
add program {name} {vertex-shader-path} {fragment-shader-path} # Adds a program with the shaders. Currently only the name program named "default" is supported.
reload programs
```

On daemon startup, the following commands are run:
```
add program default "shaders/vert/default.glsl" "shaders/frag/default.glsl"
add source default "sample.mp4"
add layer default default default 0
```

Quotation marks are understood, and are required for most filepaths and numeric values with decimals.

Additional commands can be run at runtime using netcat. By default the daemon listens at `localhost:4200`. This can be changed by starting the daemon with `-net=":8765"`.

On OSX this looks like:
```
dianes-air:gvd dianelooney$ nc localhost 4200
add program default "shaders/vert/default.glsl" "shaders/frag/rave.glsl"
add source default "sample2.mp4"
add layer default default default 1.0
```
