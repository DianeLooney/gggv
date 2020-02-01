// trans pride
uniform vec3 color0 = vec3(0, 0, 0);
uniform vec3 color1 =
    vec3(85.0 / 255.0, 205.0 / 255.0, 252.0 / 255.0);  // blue (85, 205, 252)
uniform vec3 color2 =
    vec3(247.0 / 255.0, 168.0 / 255.0, 184.0 / 255.0);  // pink (247, 168, 184)
uniform vec3 color3 = vec3(1, 1, 1);
uniform float color0Wt = 1;
uniform float color1Wt = 1;
uniform float color2Wt = 1;
uniform float color3Wt = 2;

/*
// nb pride
uniform vec3 color0 = vec3(0, 0, 0);
uniform vec3 color1 = vec3(155.0 / 255.0, 89.0 / 255.0, 208.0 / 255.0); //
purple (155,89,208) uniform vec3 color2 = vec3(255.0 / 255.0, 244.0 /
255.0, 51.0 / 255.0); // yellow (255,244,51) uniform vec3 color3 = vec3(1,1,1);
uniform float color0Wt = 1;
uniform float color1Wt = 1;
uniform float color2Wt = 1;
uniform float color3Wt = 0;
*/
/*
// ace pride
uniform vec3 color0 = vec3(0, 0, 0);
uniform vec3 color1 = vec3(129.0 / 255.0, 0.0 / 255.0, 129.0 / 255.0); // purple
(129,0,129) uniform vec3 color2 = vec3(164.0 / 255.0, 164.0 / 255.0, 164.0 /
255.0); // gray (164,0,164) uniform vec3 color3 = vec3(1,1,1); uniform float
color0Wt = 1; uniform float color1Wt = 2; uniform float color2Wt = 1; uniform
float color3Wt = 3;
*/
/*
// bi pride
uniform vec3 color0 = vec3(0, 0, 0);
uniform vec3 color1 = vec3(  0.0 / 255.0,  56.0 / 255.0, 168.0 / 255.0); //
purple (0, 56, 168) uniform vec3 color2 = vec3(115.0 / 255.0,  79.0 / 255.0,
150.0 / 255.0); // yellow (115, 79, 150) uniform vec3 color3 = vec3(215.0 /
255.0,   2.0 / 255.0, 112.0 / 255.0); uniform float color0Wt = 1; uniform float
color1Wt = 2; uniform float color2Wt = 2; uniform float color3Wt = 3;
*/

void main() {
  vec4 color = texture(tex0, fragTexCoord);
  float totalWt = color0Wt + color1Wt + color2Wt + color3Wt;
  float brightness = (color.r + color.g + color.b) / 3;

  brightness *= totalWt;

  if (brightness < color0Wt) {
    outputColor = vec4(color0, 1);
  } else if (brightness < color0Wt + color1Wt) {
    outputColor = vec4(color1, 1);
  } else if (brightness < color0Wt + color1Wt + color2Wt) {
    outputColor = vec4(color2, 1);
  } else if (brightness < color0Wt + color1Wt + color2Wt + color3Wt) {
    outputColor = vec4(color3, 1);
  }
}
