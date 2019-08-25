
#version 330

uniform sampler2D tex0;
uniform sampler2D tex1;
uniform sampler2D tex2;
uniform float time;

uniform float windowWidth;
uniform float cursorX;

uniform float ampl;

uniform float randR;
uniform float randG;
uniform float randB;

in vec2 fragTexCoord;

out vec4 outputColor;

const float PI = 3.1415926535897932384626433832795;

void main() {
    outputColor = texture(tex0, fragTexCoord) + vec4(randR, randG, randB, 1.0);
}
