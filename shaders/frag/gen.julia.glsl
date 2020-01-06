vec2 func(vec2 o, vec2 c) {
    return c + vec2(
        o.x * o.x - o.y * o.y,
        2 * o.x * o.y
    );
}

void main() {
  float r = 20;
  float cx = 0.7 * sin(time);
  float cy = 0.7 * cos(time);
  outputColor = texture(tex0, fragTexCoord);
  vec2 v = 4 * (fragTexCoord - vec2(0.5, 0.5));
  vec2 c = vec2(cx, cy);
  float baseIter = 20;
  float iterMult = 1;
  float iter = baseIter * iterMult;
  for (float i = 0; i < iter; i++) {
    v = func(v, c);
    if (length(v) > r) {
      float x = ( i / iter) * iterMult;
      outputColor = vec4(x / 49, x / 7, x, 1);
      return;
    }
  }
  outputColor = vec4(0,0,0,1);
}
