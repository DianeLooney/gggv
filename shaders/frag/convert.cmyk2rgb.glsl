
void main() {
  outputColor = texture(tex0, fragTexCoord);
  float r = outputColor.r;
  float g = outputColor.g;
  float b = outputColor.b;
  float k = 1 - max(max(outputColor.r, outputColor.g), outputColor.b);
  outputColor = vec4(
    (1 - r - k) / (1 - k),
    (1 - g - k) / (1 - k),
    (1 - b - k) / (1 - k),
    1
  );
}

