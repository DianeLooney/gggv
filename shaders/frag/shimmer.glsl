uniform float rate = 0.02;
uniform float threshold = 1.4;

void main() {
  vec4 col = texture(tex0, fragTexCoord);
  outputColor = texture(lastFrame, fragTexCoord);
  if (col.r < outputColor.r) outputColor.r += rate;
  if (col.g < outputColor.g) outputColor.g += rate;
  if (col.b < outputColor.b) outputColor.b += rate;
  if (col.r > outputColor.r) outputColor.r -= rate;
  if (col.g > outputColor.g) outputColor.g -= rate;
  if (col.b > outputColor.b) outputColor.b -= rate;
  outputColor.a = 1;

  if (outputColor.r < 0
      || outputColor.g < 0
      || outputColor.b < 0
      || outputColor.r > 1
      || outputColor.g > 1
      || outputColor.b > 1
      || abs(col.r - outputColor.r) > threshold 
      || abs(col.g - outputColor.g) > threshold
      || abs(col.b - outputColor.b) > threshold) outputColor = col;
}
