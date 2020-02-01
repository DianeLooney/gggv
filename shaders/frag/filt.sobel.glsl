// zero breaks everything so we dont let the user enter 0
uniform float strength;
float internalStrength = strength == 0 ? 0.0001 : strength;

mat3 x_kernel;
mat3 y_kernel;

void main() {
  x_kernel[0] = vec3(-1, 0, 1);
  x_kernel[1] = vec3(-2, 0, 2);
  x_kernel[2] = vec3(-1, 0, 1);

  y_kernel[0] = vec3(1, 2, 1);
  y_kernel[1] = vec3(0, 0, 0);
  y_kernel[2] = vec3(-1, -2, -1);

  float x_step = 1 / windowWidth;
  float y_step = 1 / windowHeight;

  vec4 x_aggregator = vec4(0);
  x_aggregator +=
      x_kernel[0][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),
                                          (y_step * -1 + fragTexCoord.y)));
  x_aggregator +=
      x_kernel[0][1] * texture(tex0, vec2((x_step * 0 + fragTexCoord.x),
                                          (y_step * -1 + fragTexCoord.y)));
  x_aggregator +=
      x_kernel[0][2] * texture(tex0, vec2((x_step * 1 + fragTexCoord.x),
                                          (y_step * -1 + fragTexCoord.y)));

  x_aggregator +=
      x_kernel[1][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),
                                          (y_step * 0 + fragTexCoord.y)));
  x_aggregator +=
      x_kernel[1][1] * texture(tex0, vec2((x_step * 0 + fragTexCoord.x),
                                          (y_step * 0 + fragTexCoord.y)));
  x_aggregator +=
      x_kernel[1][2] * texture(tex0, vec2((x_step * 1 + fragTexCoord.x),
                                          (y_step * 0 + fragTexCoord.y)));

  x_aggregator +=
      x_kernel[2][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),
                                          (y_step * 1 + fragTexCoord.y)));
  x_aggregator +=
      x_kernel[2][1] * texture(tex0, vec2((x_step * 0 + fragTexCoord.x),
                                          (y_step * 1 + fragTexCoord.y)));
  x_aggregator +=
      x_kernel[2][2] * texture(tex0, vec2((x_step * 1 + fragTexCoord.x),
                                          (y_step * 1 + fragTexCoord.y)));

  vec4 y_aggregator = vec4(0);
  y_aggregator +=
      y_kernel[0][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),
                                          (y_step * -1 + fragTexCoord.y)));
  y_aggregator +=
      y_kernel[0][1] * texture(tex0, vec2((x_step * 0 + fragTexCoord.x),
                                          (y_step * -1 + fragTexCoord.y)));
  y_aggregator +=
      y_kernel[0][2] * texture(tex0, vec2((x_step * 1 + fragTexCoord.x),
                                          (y_step * -1 + fragTexCoord.y)));

  y_aggregator +=
      y_kernel[1][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),
                                          (y_step * 0 + fragTexCoord.y)));
  y_aggregator +=
      y_kernel[1][1] * texture(tex0, vec2((x_step * 0 + fragTexCoord.x),
                                          (y_step * 0 + fragTexCoord.y)));
  y_aggregator +=
      y_kernel[1][2] * texture(tex0, vec2((x_step * 1 + fragTexCoord.x),
                                          (y_step * 0 + fragTexCoord.y)));

  y_aggregator +=
      y_kernel[2][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),
                                          (y_step * 1 + fragTexCoord.y)));
  y_aggregator +=
      y_kernel[2][1] * texture(tex0, vec2((x_step * 0 + fragTexCoord.x),
                                          (y_step * 1 + fragTexCoord.y)));
  y_aggregator +=
      y_kernel[2][2] * texture(tex0, vec2((x_step * 1 + fragTexCoord.x),
                                          (y_step * 1 + fragTexCoord.y)));

  outputColor = sqrt(x_aggregator * x_aggregator + y_aggregator * y_aggregator);

  outputColor.a = 1;
}
