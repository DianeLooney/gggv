uniform float brightness;

void main() {
  outputColor = texture(tex0, fragTexCoord) * brightness;
}
