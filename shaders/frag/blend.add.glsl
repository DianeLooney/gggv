void main() {
  outputColor = texture(tex0, fragTexCoord) + texture(tex1, fragTexCoord) +  texture(tex2, fragTexCoord);
}
