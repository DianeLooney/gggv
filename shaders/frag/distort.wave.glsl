uniform float frequency;
uniform float amplitude;
uniform float speed;

void main() {
  vec2 c = fragTexCoord;
  c.x = c.x + sin((fragTexCoord.y * frequency) + (time * speed)) * amplitude;
  outputColor = texture(tex0, c);
}
