uniform float slats = 20;
uniform float minSize = 0.05;

void main() {
  vec2 coord = fragTexCoord;
  coord.y = floor(coord.y * slats) / slats;
  vec4 sample = texture(tex0, coord);
  float ht = length(sample.rgb) / 1.5;
  float dy = fract(fragTexCoord.y * slats);
  if ((dy < 0.5 && dy > 1 - ht) || (dy > 0.5 && dy < ht) ||
      (dy < 0.5 && dy > 0.5 - minSize/2) || (dy > 0.5 && dy < 0.5 + minSize/2)) {
    outputColor = sample;
  } else {
    outputColor = vec4(0,0,0,1);
  }
}
