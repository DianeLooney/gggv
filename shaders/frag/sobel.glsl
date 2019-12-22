#version 330

//default uniforms from GGGV

//futher sources are set as tex1, tex2, tex3
uniform sampler2D tex0;

//result of the last render of this shader
uniform sampler2D lastFrame;

//only set when tex* is a FFsource
uniform float tex0Width;
uniform float tex0Height;

//seconds since application start. 
//monotonically increasing, will never be 0 
uniform float time;
//time the last frame took to render
uniform float renderTime;
//number of frames in the last second
uniform float fps;
//output window size
uniform float windowHeight;
uniform float windowWidth;
//cursor position
uniform float cursorX;
uniform float cursorY;

//texCoord
in vec2 fragTexCoord;
//output pixel color
out vec4 outputColor;

//zero breaks everything so we dont let the user enter 0
uniform float strength;
float internalStrength = strength == 0 ? 0.0001 : strength;

mat3 x_kernel;
mat3 y_kernel;


void main() {
  x_kernel[0] = vec3(-1, 0, 1);
  x_kernel[1] = vec3(-2, 0, 2);
  x_kernel[2] = vec3(-1, 0, 1);

  y_kernel[0] = vec3( 1, 2, 1);
  y_kernel[1] = vec3( 0, 0, 0);
  y_kernel[2] = vec3(-1,-2,-1);

  float x_step = 1/windowWidth;
  float y_step = 1/windowHeight;

  vec4 x_aggregator = vec4(0);
  x_aggregator += x_kernel[0][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));
  x_aggregator += x_kernel[0][1] * texture(tex0, vec2((x_step *  0 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));
  x_aggregator += x_kernel[0][2] * texture(tex0, vec2((x_step *  1 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));

  x_aggregator += x_kernel[1][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));
  x_aggregator += x_kernel[1][1] * texture(tex0, vec2((x_step *  0 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));
  x_aggregator += x_kernel[1][2] * texture(tex0, vec2((x_step *  1 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));

  x_aggregator += x_kernel[2][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));
  x_aggregator += x_kernel[2][1] * texture(tex0, vec2((x_step *  0 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));
  x_aggregator += x_kernel[2][2] * texture(tex0, vec2((x_step *  1 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));

  vec4 y_aggregator = vec4(0);
  y_aggregator += y_kernel[0][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));
  y_aggregator += y_kernel[0][1] * texture(tex0, vec2((x_step *  0 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));
  y_aggregator += y_kernel[0][2] * texture(tex0, vec2((x_step *  1 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y)));

  y_aggregator += y_kernel[1][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));
  y_aggregator += y_kernel[1][1] * texture(tex0, vec2((x_step *  0 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));
  y_aggregator += y_kernel[1][2] * texture(tex0, vec2((x_step *  1 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y)));

  y_aggregator += y_kernel[2][0] * texture(tex0, vec2((x_step * -1 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));
  y_aggregator += y_kernel[2][1] * texture(tex0, vec2((x_step *  0 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));
  y_aggregator += y_kernel[2][2] * texture(tex0, vec2((x_step *  1 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y)));


  outputColor = sqrt(x_aggregator*x_aggregator+y_aggregator*y_aggregator);

  outputColor.a = 1;
}

