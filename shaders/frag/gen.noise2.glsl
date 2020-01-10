uniform vec2 pixelate1 = vec2(44, 30);
uniform vec2 pixelate2 = vec2(220, 150);

void main() {
  outputColor = vec4(
    noise(floor(fragTexCoord * pixelate1), time) *
    noise(floor(fragTexCoord * pixelate2), time));
}
