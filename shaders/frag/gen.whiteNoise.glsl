#version 330

uniform vec4 color;

in vec2 fragTexCoord;
out vec4 outputColor;

void main() {
  outputColor = vec4(fract(sin(dot(fragTexCoord*time,vec2(12.9898,78.233))) * 43758.5453));
}
