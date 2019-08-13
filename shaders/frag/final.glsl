#version 330

uniform sampler2D tex;
uniform sampler2D prevPass;
uniform sampler2D prevFrame;
uniform float time;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    outputColor = 10*(texture(prevPass, fragTexCoord) - 0.95 * texture(prevFrame, fragTexCoord));
}
