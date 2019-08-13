# Common Uniforms

The following uniforms are provided by default to every shader.

```
uniform mat4 projection; // uniform projection matrix
uniform mat4 camera; // camera position + facing matrix
uniform mat4 model; // currently disabled

uniform sampler2D prevFrame; // what was rendered in the previous frame
uniform sampler2D prevPass: // what was rendered up until this current layer
uniform sampler2D tex; // the source video

uniform float time; // amount of time since daemon startup (in seconds)

uniform int fps; // rough fps estimate, updates ~ once per second, slightly inaccurate until it can be done properly
uniform float renderTime; // amount of time previous frame took to render (in seconds)
```
