#version 330

uniform sampler2D tex0;
uniform float time;


uniform float frequency;
uniform float amplitude;
uniform float speed;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
  vec2 c = fragTexCoord;
  c.x = c.x + sin((fragTexCoord.y*frequency)+(time*speed))*amplitude;
  outputColor = texture(tex0, c);
}
