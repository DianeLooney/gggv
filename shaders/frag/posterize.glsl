#version 330

uniform float gamma = 1;
uniform float bins = 5;

uniform sampler2D tex0;
in vec2 fragTexCoord;
out vec4 outputColor;

void main() {
  outputColor = texture(tex0, fragTexCoord);
  outputColor = pow(outputColor, vec4(gamma));
  outputColor *= vec4(bins);
  outputColor = floor(outputColor);
  outputColor /= vec4(bins);
  outputColor = pow(outputColor, vec4(1.0/gamma));
  outputColor.a = 1;
}
