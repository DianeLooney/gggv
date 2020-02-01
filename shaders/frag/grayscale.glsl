void main() {
  outputColor = texture(tex0, fragTexCoord);
  float val = (outputColor.r + outputColor.g + outputColor.b) / 3;
  outputColor = vec4(val, val, val, 1);
}
