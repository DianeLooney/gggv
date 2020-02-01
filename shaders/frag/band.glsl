uniform float rBands = 2.0;
uniform float gBands = 2.0;
uniform float bBands = 2.0;

void main() {
  vec4 color = texture(tex0, fragTexCoord);

  if (rBands == 0) {
    color.r = 0;
  } else {
    color.r = floor(0.5 + color.r * (rBands + 1)) / rBands;
  }
  if (gBands == 0) {
    color.g = 0;
  } else {
    color.g = floor(0.5 + color.g * (gBands + 1)) / gBands;
  }
  if (bBands == 0) {
    color.b = 0;
  } else {
    color.b = floor(0.5 + color.b * (bBands + 1)) / bBands;
  }
  color.a = 1;
  outputColor = color;
}
