void main() {
  vec4 resampleCoords = texture(tex1, fragTexCoord);
  outputColor = texture(tex0, resampleCoords.rg) * resampleCoords.b;
  outputColor.a = 1;
}
