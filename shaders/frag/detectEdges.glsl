uniform float dxSearchDistance = 1.0;
uniform float dySearchDistance = 1.0;
uniform float dxSearchSamples = 2.0;
uniform float dySearchSamples = 2.0;
uniform float threshold = 0.13;
uniform float rCull = 1.0;
uniform float gCull = 1.0;
uniform float bCull = 1.0;

void main() {
  float dx = dxSearchDistance / windowWidth;
  float dy = dySearchDistance / windowHeight;

  float minr = 1.0;
  float ming = 1.0;
  float minb = 1.0;
  float maxr = 0.0;
  float maxg = 0.0;
  float maxb = 0.0;

  float r, g, b;
  vec4 sample;
  float x, y;
  float stride = 3;
  for(x = fragTexCoord.x - dxSearchSamples * dx; x <= fragTexCoord.x + dxSearchSamples * dx; x += dx) {
    for(y = fragTexCoord.y - dySearchSamples * dy; y <= fragTexCoord.y + dySearchSamples * dy; y += dy) {
      sample = texture(tex0, vec2(x, y));
      r = sample.r;
      g = sample.g;
      b = sample.b;
      if (r < minr) minr = r;
      if (g < ming) ming = g;
      if (b < minb) minb = b;
      if (r > maxr) maxr = r;
      if (g > maxg) maxg = g;
      if (b > maxb) maxb = b;
    }
  }
  sample = texture(tex0, fragTexCoord);
  if ((maxr - minr) + (maxg - ming) + (maxb - minb) > threshold) {
    outputColor = vec4(1,1,1,1);
  } else {
    if (maxr - minr < rCull) sample.r = 0;
    if (maxg - ming < gCull) sample.g = 0;
    if (maxb - minb < bCull) sample.b = 0;
    outputColor = sample;
  }
}
