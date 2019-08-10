
#version 330

uniform sampler2D tex;
uniform float time;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
	vec4 color = texture(tex, fragTexCoord);
    outputColor = vec4(color.r, color.g, color.b, 1);
}
