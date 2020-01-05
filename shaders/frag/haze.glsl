uniform float threshold = 0;
uniform float decay = 0.003;

void main() {
  outputColor = texture(tex0, fragTexCoord);
  outputColor.a = 1;
  vec4 decayed = texture(lastFrame, fragTexCoord);
  decayed.a = 1;
  if (outputColor.r > decayed.r) outputColor.r = decayed.r + decay;
  else outputColor.r = decayed.r - decay;
  if (outputColor.g > decayed.g) outputColor.g = decayed.g + decay;
  else outputColor.g = decayed.g - decay;
  if (outputColor.b > decayed.b) outputColor.b = decayed.b + decay;
  else outputColor.b = decayed.b - decay;
}
