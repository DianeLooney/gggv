#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
uniform float depth;

in vec3 vert;
in vec2 vertTexCoord;

out vec2 geomTexCoord;

void main() {
    geomTexCoord = vertTexCoord;
    gl_Position = projection * camera * vec4(vert, 1);
}
