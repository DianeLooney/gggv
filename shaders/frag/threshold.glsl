uniform float fadeRate = 0.003;
uniform float threshold = 0.4;

uniform float highlightR = 255;
uniform float highlightG = 255;
uniform float highlightB = 255;

void main() {
  outputColor = texture(tex0, fragTexCoord);
  float total = (outputColor.r + outputColor.g + outputColor.b) / 3;

  if (total > threshold) {
    outputColor = vec4(1, 1, 1, 1);
  } else {
    outputColor = vec4(0, 0, 0, 1);
  }
}
