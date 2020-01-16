vec2 res = vec2(windowWidth / windowHeight, 1);
vec2 chunk(vec2 v, float n) {
  return floor((v-0.5) *res * n)/n+0.5;
}

uniform vec3 color = vec3(0, 1, 0);

void main() {
  outputColor = vec4(noise(chunk((screenToPolar(fragTexCoord)+vec2(time, 0)) * vec2(0.001 * smoothstep(-1, 1, 1+sin(time)), 0.01) * vec2(1, tan(time*10)), 16 + 8 * sin(time*14))));
  vec4 last = texture(lastFrame, fragTexCoord);
  float bright = length(outputColor.r);
  if (bright < 0.03) {
    outputColor = vec4(color, 1);
    return;
  }
  
  if (bright < 0.97) {
    outputColor = last;
    return;
  }
  outputColor = vec4(0,0,0,1);
}

