#version 330

uniform float brightness;
uniform sampler2D tex0;

in vec2 fragTexCoord;
out vec4 outputColor;

void main() {
    outputColor = texture(tex0, fragTexCoord) * brightness;
}
