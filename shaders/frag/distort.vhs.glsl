#version 330

uniform sampler2D tex0;
uniform float time;
uniform float windowHeight;
uniform float windowWidth;

in vec2 fragTexCoord;
in vec2 screenCoord;
out vec4 outputColor;

uniform float bandSize = 150;
uniform float bumpCount = 80;
uniform float interval = 10;

void main() {
  float factor = screenCoord.y / bandSize; 
  float distX = sin(screenCoord.y / (sin(factor) * bumpCount));
  vec2 sampleCoord = fragTexCoord;
  sampleCoord.x += distX / bandSize;
  float stride = 2 * atan(interval);
  //sampleCoord.x += mod(atan(sampleCoord.x - 0.5 + mod(time * 2, 2 * interval) - interval)/stride+1, 2)-1;
  sampleCoord.x += pow(cos((time + sampleCoord.x * 6.28) / 2), 25) / 10;
  // sampleCoord.x -= 0.5;
  sampleCoord.y = (screenCoord.y + bandSize * cos(factor + time / 10)) / windowHeight;
  outputColor = texture(tex0, sampleCoord);
}
