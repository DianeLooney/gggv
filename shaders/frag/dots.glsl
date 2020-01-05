uniform float dist = 70;
uniform float diam = 15;

void main() {
  vec2 sampleCoords = vec2(
    floor(screenCoord.x / dist) * dist,
    floor(screenCoord.y / dist) * dist
  );
  vec2 centerCoords = vec2(
    floor(screenCoord.x / dist) * dist + dist/2,
    floor(screenCoord.y / dist) * dist + dist/2
  );
  outputColor = texture(tex0, sampleCoords / vec2(windowWidth, windowHeight));
  float dotSize = (diam / 3) * (outputColor.r + outputColor.g + outputColor.b);
  if (length(centerCoords - centerCoords) > dotSize) {
    outputColor = vec4(0,0,0,1);
  }
}
