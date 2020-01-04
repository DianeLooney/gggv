#version 330

uniform sampler2D tex0;
uniform sampler2D tex1;

in vec2 fragTexCoord;
out vec4 outputColor;

void main() {
  outputColor = texture(tex0, fragTexCoord) * texture(tex1, fragTexCoord);
}
