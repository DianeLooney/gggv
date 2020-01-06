uniform float gamma = 1;
uniform float bins = 5;

void main() {
  outputColor = texture(tex0, fragTexCoord);
  outputColor = pow(outputColor, vec4(gamma));
  outputColor *= vec4(bins);
  outputColor = floor(outputColor);
  outputColor /= vec4(bins);
  outputColor = pow(outputColor, vec4(1.0/gamma));
  outputColor.a = 1;
}
