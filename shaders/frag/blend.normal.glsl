#version 330

uniform sampler2D tex0;
uniform sampler2D tex1;
uniform sampler2D tex2;

in vec2 fragTexCoord;
out vec4 outputColor;

void main() {
    vec4 mask = texture(tex0, fragTexCoord);
    vec4 sample1 = texture(tex1, fragTexCoord);
    vec4 sample2 = texture(tex2, fragTexCoord);
    outputColor = mask*sample1 + (vec4(1,1,1,1) - mask) * sample2;
}
