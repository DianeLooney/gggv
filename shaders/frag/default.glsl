
#version 330

uniform sampler2D tex0;
uniform sampler2D tex1;
uniform sampler2D tex2;
uniform float time;

in vec2 fragTexCoord;

out vec4 outputColor;

const float PI = 3.1415926535897932384626433832795;

void main() {
	vec4 color0 = texture(tex0, fragTexCoord);
    vec4 color1 = texture(tex1, fragTexCoord);
    vec4 color2 = texture(tex2, fragTexCoord);
    outputColor = vec4(
        0.25 + 1.2 * (sin(time) * color0.r + sin(time + 2.0 * PI / 3.0) * color1.r + sin(time + 4.0 * PI / 3.0) * color2.r),
        0.25 + 1.2 * (sin(time) * color1.g + sin(time + 2.0 * PI / 3.0) * color2.g + sin(time + 4.0 * PI / 3.0) * color0.g),
        0.25 + 1.2 * (sin(time) * color2.b + sin(time + 2.0 * PI / 3.0) * color0.b + sin(time + 4.0 * PI / 3.0) * color1.b),
        1
    );
}
