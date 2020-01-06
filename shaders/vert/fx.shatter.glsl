void main() {
  geomTexCoord = vertTexCoord;
  gl_Position = projection * camera * vec4(vert, 1);
  float r = 2 + sin(2*gl_Position.x + 5*gl_Position.y ) / 2;   
  float dy = sin(r * time + (2+gl_Position.x) * (3+gl_Position.y)) / 15;
  float dx = sin(r * time + (4+gl_Position.x) * (2+gl_Position.y)) / 15;

  gl_Position.x += sin(gl_Position.y * 24) / 10;
  gl_Position.y += sin(gl_Position.x * 12) / 10;
  gl_Position.x += dx;
  gl_Position.y += dy;
}
