uniform float dxPixels = 1.0;
uniform float dyPixels = 1.0;
uniform float xSamples = 1;
uniform float ySamples = 1;
uniform float rThreshold = 0.45;
uniform float gThreshold = 0.4;
uniform float bThreshold = 0.4;

void main() {
  float dx = (dxPixels+0.1) / windowWidth;
  float dy = (dyPixels+0.1) / windowHeight;

  float x, y, i, j;
  vec4 color = vec4(0,0,0,1);
  for(i = -xSamples; i <= xSamples; i ++) {
    for(j = -ySamples; j <= ySamples; j++) {
      color += texture(tex0, vec2(fragTexCoord.x + i * dx * dxPixels, fragTexCoord.y + j * dy * dyPixels));
    }
  }
  color /= (1.0 + 2.0 * xSamples) * (1.0 + 2.0 * ySamples);

  color.r = color.r > rThreshold ? 1 : 0;
  color.g = color.g > gThreshold ? 1 : 0;
  color.b = color.b > bThreshold ? 1 : 0;
  color.a = 1;
  outputColor = color;
}
