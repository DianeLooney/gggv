
#version 330

uniform sampler2D tex0;
uniform sampler2D tex1;
uniform sampler2D tex2;
uniform float time;

uniform float windowWidth;
uniform float cursorX;

uniform float ampl;

in vec2 fragTexCoord;

out vec4 outputColor;

const float PI = 3.1415926535897932384626433832795;

void main() {
    outputColor = texture(tex0, fragTexCoord) + texture(tex1, fragTexCoord) + texture(tex2, fragTexCoord);
}
