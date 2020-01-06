void main() {
  geomTexCoord = vertTexCoord;
  geomScreenCoord = geomTexCoord * vec2(windowWidth + 1, windowHeight + 1);
  gl_Position = projection * camera * vec4(vert, 1);
}
