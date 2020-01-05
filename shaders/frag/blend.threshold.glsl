uniform float threshold = 0.1;

void main() {
  outputColor = texture(tex0, fragTexCoord);
  if (length(outputColor.rgb) > threshold) return;
  outputColor = texture(tex1, fragTexCoord);
  if (length(outputColor.rgb) > threshold) return;
  outputColor = texture(tex2, fragTexCoord);
  if (length(outputColor.rgb) > threshold) return;
  outputColor = texture(tex3, fragTexCoord);
  if (length(outputColor.rgb) > threshold) return;
  outputColor = texture(tex4, fragTexCoord);
  if (length(outputColor.rgb) > threshold) return;
  outputColor = texture(tex5, fragTexCoord);
  if (length(outputColor.rgb) > threshold) return;
  outputColor = texture(tex6, fragTexCoord);
  if (length(outputColor.rgb) > threshold) return;
  outputColor = texture(tex7, fragTexCoord);
  if (length(outputColor.rgb) > threshold) return;
  outputColor = texture(tex8, fragTexCoord);
}
