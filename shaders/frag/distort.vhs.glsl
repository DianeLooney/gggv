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
uniform float interval = 4;
uniform float invertDuration = 0.23;
uniform float invertFrequency = 4;

void main() {
  float factor = screenCoord.y / bandSize; 
  float distX = sin(screenCoord.y / (sin(factor) * bumpCount));
  vec2 sampleCoord = fragTexCoord;
  sampleCoord.x += distX / bandSize;
  float stride = 2 * atan(interval);
  float invert = floor(mod(time, invertFrequency + invertDuration) / invertFrequency);
  sampleCoord.x += mod(atan(sampleCoord.x - 0.5 + mod(time * 2, 2 * interval) - interval)/stride+1, 2)-1;
  sampleCoord.x -= 0.5;
  sampleCoord.x -= 0.3 * invert * sin(sampleCoord.y + time);
  sampleCoord.y = (screenCoord.y + bandSize * cos(factor)) / windowHeight;
  if (invert == 1) sampleCoord.y *= -1;
  outputColor = texture(tex0, sampleCoord);
}
