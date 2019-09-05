#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in int flipOutput;
in vec3 vert;
in vec2 vertTexCoord;

out vec2 geomTexCoord;

void main() {
    geomTexCoord = vertTexCoord;
    if (flipOutput != 0) {
        geomTexCoord.y = 1 - geomTexCoord.y;   
    }
    gl_Position = projection * camera * vec4(vert, 1);
}
