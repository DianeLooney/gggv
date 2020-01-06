uniform float flakes = 4;
uniform float radius = 300;

void main() {
  vec2 coords = fragTexCoord;
  coords = screenToPolar(coords);
  coords.y += time * 0.1;
  coords.x -= 1 * time;
  coords.x -= 1080 * time;
  coords.x = mod(coords.x, 1080);
  if (coords.x > radius) {
    coords.x = 2 * radius - coords.x;
  }
  if (coords.x < 0) {
    coords.x = -coords.x;
  }
  coords.y = mod(coords.y, PI / flakes);
  coords.x += (coords.y / 50) *  cos(coords.y * 8);
  if (coords.y > PI / (2*flakes)) coords.y = PI / flakes - coords.y;
  coords = polarToScreen(coords);
  outputColor = texture(tex0, coords);
}
