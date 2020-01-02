#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

uniform float windowHeight;
uniform float windowWidth;

in vec3 vert;
in vec2 vertTexCoord;

out vec2 geomTexCoord;
out vec2 geomScreenCoord;

void main() {
  geomTexCoord = vertTexCoord;
  geomScreenCoord = geomTexCoord * vec2(windowWidth + 1, windowHeight + 1);
  gl_Position = projection * camera * vec4(vert, 1);
}
