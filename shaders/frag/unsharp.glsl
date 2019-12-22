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

float _0x0 = internalStrength * 1;
float _0x1 = internalStrength * 4;
float _0x2 = internalStrength * 6;
float _0x3 = internalStrength * 4;
float _0x4 = internalStrength * 1;
float _1x0 = internalStrength * 4;
float _1x1 = internalStrength * 16;
float _1x2 = internalStrength * 24;
float _1x3 = internalStrength * 16;
float _1x4 = internalStrength * 4;
float _2x0 = internalStrength * 6;
float _2x1 = internalStrength * 24;
float _2x3 = internalStrength * 24;
float _2x4 = internalStrength * 6;
float _3x0 = internalStrength * 4;
float _3x1 = internalStrength * 16;
float _3x2 = internalStrength * 24;
float _3x3 = internalStrength * 16;
float _3x4 = internalStrength * 4;
float _4x0 = internalStrength * 1;
float _4x1 = internalStrength * 4;
float _4x2 = internalStrength * 6;
float _4x3 = internalStrength * 4;
float _4x4 = internalStrength * 1;
float _2x2 = -256 + (-1 * ( _0x0 + _0x1 + _0x2 + _0x3 + _0x4 + _1x0 + _1x1 + _1x2 + _1x3 + _1x4 + _2x0 + _2x1 + _2x3 + _2x4 + _3x0 + _3x1 + _3x2 + _3x3 + _3x4 + _4x0 + _4x1 + _4x2 + _4x3 + _4x4));

float coefficient_numerator   = -1;
float coefficient_denomenator = 256;


void main() {
  float x_step = 1/windowWidth;
  float y_step = 1/windowHeight;
  vec4 aggregator = vec4(0);
  aggregator += texture(tex0, vec2((x_step * -2 + fragTexCoord.x),(y_step*-2 + fragTexCoord.y))) * _0x0;
  aggregator += texture(tex0, vec2((x_step * -1 + fragTexCoord.x),(y_step*-2 + fragTexCoord.y))) * _0x1;
  aggregator += texture(tex0, vec2((x_step *  0 + fragTexCoord.x),(y_step*-2 + fragTexCoord.y))) * _0x2;
  aggregator += texture(tex0, vec2((x_step *  1 + fragTexCoord.x),(y_step*-2 + fragTexCoord.y))) * _0x3;
  aggregator += texture(tex0, vec2((x_step *  2 + fragTexCoord.x),(y_step*-2 + fragTexCoord.y))) * _0x4;

  aggregator += texture(tex0, vec2((x_step * -2 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y))) * _1x0;
  aggregator += texture(tex0, vec2((x_step * -1 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y))) * _1x1;
  aggregator += texture(tex0, vec2((x_step *  0 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y))) * _1x2;
  aggregator += texture(tex0, vec2((x_step *  1 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y))) * _1x3;
  aggregator += texture(tex0, vec2((x_step *  2 + fragTexCoord.x),(y_step*-1 + fragTexCoord.y))) * _1x4;

  aggregator += texture(tex0, vec2((x_step * -2 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y))) * _2x0;
  aggregator += texture(tex0, vec2((x_step * -1 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y))) * _2x1;
  aggregator += texture(tex0, vec2((x_step *  0 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y))) * _2x2;
  aggregator += texture(tex0, vec2((x_step *  1 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y))) * _2x3;
  aggregator += texture(tex0, vec2((x_step *  2 + fragTexCoord.x),(y_step* 0 + fragTexCoord.y))) * _2x4;

  aggregator += texture(tex0, vec2((x_step * -2 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y))) * _3x0;
  aggregator += texture(tex0, vec2((x_step * -1 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y))) * _3x1;
  aggregator += texture(tex0, vec2((x_step *  0 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y))) * _3x2;
  aggregator += texture(tex0, vec2((x_step *  1 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y))) * _3x3;
  aggregator += texture(tex0, vec2((x_step *  2 + fragTexCoord.x),(y_step* 1 + fragTexCoord.y))) * _3x4;

  aggregator += texture(tex0, vec2((x_step * -2 + fragTexCoord.x),(y_step* 2 + fragTexCoord.y))) * _4x0;
  aggregator += texture(tex0, vec2((x_step * -1 + fragTexCoord.x),(y_step* 2 + fragTexCoord.y))) * _4x1;
  aggregator += texture(tex0, vec2((x_step *  0 + fragTexCoord.x),(y_step* 2 + fragTexCoord.y))) * _4x2;
  aggregator += texture(tex0, vec2((x_step *  1 + fragTexCoord.x),(y_step* 2 + fragTexCoord.y))) * _4x3;
  aggregator += texture(tex0, vec2((x_step *  2 + fragTexCoord.x),(y_step* 2 + fragTexCoord.y))) * _4x4;

  aggregator *= coefficient_numerator;
  aggregator /= coefficient_denomenator;

  outputColor = aggregator;
  outputColor.a = 1;
}

