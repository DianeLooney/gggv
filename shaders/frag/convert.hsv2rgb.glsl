// neato code from the lolengine ppl
// http://lolengine.net/blog/2013/07/27/rgb-to-hsv-in-glsl#
vec3 hsv2rgb(vec3 c) {
  vec4 K = vec4(1.0, 2.0 / 3.0, 1.0 / 3.0, 3.0);
  vec3 p = abs(fract(c.xxx + K.xyz) * 6.0 - K.www);
  return c.z * mix(K.xxx, clamp(p - K.xxx, 0.0, 1.0), c.y);
}

void main() {
  vec4 inp = texture(tex0, fragTexCoord);
  vec3 rgb = hsv2rgb(inp.rgb);
  outputColor = vec4(rgb.r, rgb.g, rgb.b, inp.a);
}
