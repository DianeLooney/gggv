uniform float decay = 0.003;
uniform float threshold = 1.2;
uniform vec3 highlightColor = vec3(1, 0, 1);
void main() {
  if (length(texture(tex0, fragTexCoord).rgb) > threshold) {
    outputColor = vec4(highlightColor, 1);
    return;
  }
  outputColor = clamp(texture(lastFrame, fragTexCoord) - vec4(decay, decay, decay, 0), 0, 1);
}
