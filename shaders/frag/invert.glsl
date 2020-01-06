void main() {
  outputColor = texture(tex0, fragTexCoord);
  outputColor = vec4(1) - outputColor;
  outputColor.a = 1;
}
