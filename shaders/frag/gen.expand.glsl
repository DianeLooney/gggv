
uniform float frequency = 3;
uniform float spread = 0.05;
uniform float FPS = 60;
uniform float lineThickness = 0.01;
uniform vec3 decay = vec3(0, 0, 0);

float randX() {
  return noise(time*frequency) * windowWidth / windowHeight;
}
float randY() {
  return noise(time*frequency + 1);
}

vec2 distToLine(vec2 center, vec2 direction) {
  vec2 adjFTC = fragTexCoord * vec2(windowWidth / windowHeight, 1);
  return (adjFTC - center) - direction * dot(adjFTC - center, direction);
}

void main() {
  vec2 center = vec2(randX(), randY());
  float rnd = noise(center) * PI;
  vec2 direction = vec2(cos(rnd), sin(rnd));
  vec2 d = distToLine(center, direction);

  float splitSize = fract(time * frequency) * spread;
  if(length(d) < splitSize) {
    outputColor = vec4(fragTexCoord, 1, 1);
    return;
  }
  if(length(d) < lineThickness + splitSize ) {
    outputColor = vec4(0,0,0,0);
    return;
  }

  vec2 unitd = normalize(d);
  outputColor = texture(lastFrame, fragTexCoord - (unitd * spread * frequency / FPS) * vec2(windowHeight / windowWidth, 1)) - vec4(decay, 0);
}
