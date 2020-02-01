void main() {
  geomTexCoord = vertTexCoord;
  gl_Position = projection * camera * vec4(vert, 1);
}
