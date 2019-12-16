#version 330

uniform sampler2D tex0;
uniform sampler2D tex1;
uniform sampler2D tex2;

in float sample;
in vec2 fragTexCoord;

out vec4 outputColor;

const float PI = 3.1415926535897932384626433832795;

void main() {
    if (sample < 1) {
        outputColor = texture(tex0, fragTexCoord);
    } else if (sample < 2) {
        outputColor = texture(tex1, fragTexCoord);
    } else {
        outputColor = texture(tex2, fragTexCoord);
    }
}
