#version 330

uniform sampler2D tex0;
uniform sampler2D lastFrame;
uniform float windowHeight;
uniform float windowWidth;
uniform float time;
in vec2 fragTexCoord;
out vec4 outputColor;


uniform float sobelForce = 0.012;
uniform float newFrameStrength = 0.1;
uniform int millisBetween = 4;

mat3 x_kernel;
mat3 y_kernel;
float x_step = 1/windowWidth;
float y_step = 1/windowHeight;

void main() {
  x_kernel[0] = vec3(-1, 0, 1);
  x_kernel[1] = vec3(-2, 0, 2);
  x_kernel[2] = vec3(-1, 0, 1);

  y_kernel[0] = vec3( 1, 2, 1);
  y_kernel[1] = vec3( 0, 0, 0);
  y_kernel[2] = vec3(-1,-2,-1);

  vec4 x_aggregator = vec4(0);
  x_aggregator += x_kernel[0][0] * texture(lastFrame, vec2((x_step * -1 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));
  x_aggregator += x_kernel[0][1] * texture(lastFrame, vec2((x_step *  0 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));
  x_aggregator += x_kernel[0][2] * texture(lastFrame, vec2((x_step *  1 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));

  x_aggregator += x_kernel[1][0] * texture(lastFrame, vec2((x_step * -1 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));
  x_aggregator += x_kernel[1][1] * texture(lastFrame, vec2((x_step *  0 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));
  x_aggregator += x_kernel[1][2] * texture(lastFrame, vec2((x_step *  1 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));

  x_aggregator += x_kernel[2][0] * texture(lastFrame, vec2((x_step * -1 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));
  x_aggregator += x_kernel[2][1] * texture(lastFrame, vec2((x_step *  0 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));
  x_aggregator += x_kernel[2][2] * texture(lastFrame, vec2((x_step *  1 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));

  vec4 y_aggregator = vec4(0);
  y_aggregator += y_kernel[0][0] * texture(lastFrame, vec2((x_step * -1 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));
  y_aggregator += y_kernel[0][1] * texture(lastFrame, vec2((x_step *  0 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));
  y_aggregator += y_kernel[0][2] * texture(lastFrame, vec2((x_step *  1 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));

  y_aggregator += y_kernel[1][0] * texture(lastFrame, vec2((x_step * -1 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));
  y_aggregator += y_kernel[1][1] * texture(lastFrame, vec2((x_step *  0 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));
  y_aggregator += y_kernel[1][2] * texture(lastFrame, vec2((x_step *  1 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));

  y_aggregator += y_kernel[2][0] * texture(lastFrame, vec2((x_step * -1 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));
  y_aggregator += y_kernel[2][1] * texture(lastFrame, vec2((x_step *  0 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));
  y_aggregator += y_kernel[2][2] * texture(lastFrame, vec2((x_step *  1 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));


  vec2 coef;
  coef.x = (x_aggregator.r + x_aggregator.g + x_aggregator.b)/3;
  coef.y = (y_aggregator.r + y_aggregator.g + y_aggregator.b)/3;
  vec2 progressedCoord = fragTexCoord+sobelForce*coef;

  vec4 progFrame = texture(lastFrame, progressedCoord);
  
  if(int(time*1000) % millisBetween == 0){
    outputColor = texture(tex0, fragTexCoord)*newFrameStrength + (1-newFrameStrength)*progFrame;
  }
  else{
    outputColor = progFrame;
  }
  outputColor.a = 1;
}

