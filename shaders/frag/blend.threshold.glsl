#version 330

uniform sampler2D tex0;
uniform sampler2D tex1;
uniform sampler2D tex2;
uniform sampler2D tex3;
uniform sampler2D tex4;
uniform sampler2D tex5;
uniform sampler2D tex6;
uniform sampler2D tex7;
uniform sampler2D tex8;

in vec2 fragTexCoord;
out vec4 outputColor;

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
