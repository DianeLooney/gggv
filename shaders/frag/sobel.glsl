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

vec4 kernel[3][3];
vec4 dot(vec4 m1[3][3], mat3 m2) {
  return (
    m1[0][0] * m2[0][0] + 
    m1[0][1] * m2[0][1] + 
    m1[0][2] * m2[0][2] + 
    m1[1][0] * m2[1][0] + 
    m1[1][1] * m2[1][1] + 
    m1[1][2] * m2[1][2] + 
    m1[2][0] * m2[2][0] + 
    m1[2][1] * m2[2][1] + 
    m1[2][2] * m2[2][2]
  );
}

void requireKernel() {
  vec2 dx = vec2(1.0 / windowWidth, 0);
  vec2 dy = vec2(0, 1.0 / windowHeight);
  kernel[0][0] = texture(tex0, fragTexCoord + -1*dx + -1*dy);
  kernel[0][1] = texture(tex0, fragTexCoord + -1*dx +  0*dy);
  kernel[0][2] = texture(tex0, fragTexCoord + -1*dx +  1*dy);
  kernel[1][0] = texture(tex0, fragTexCoord +  0*dx + -1*dy);
  kernel[1][1] = texture(tex0, fragTexCoord +  0*dx +  0*dy);
  kernel[1][2] = texture(tex0, fragTexCoord +  0*dx +  1*dy);
  kernel[2][0] = texture(tex0, fragTexCoord +  1*dx + -1*dy);
  kernel[2][1] = texture(tex0, fragTexCoord +  1*dx +  0*dy);
  kernel[2][2] = texture(tex0, fragTexCoord +  1*dx +  1*dy);
}

void main() {
  requireKernel();
  mat3 xWeights = mat3(
    -1, 0, 1,
    -2, 0, 2,
    -1, 0, 1
  );
  mat3 yWeights = mat3(
     1,  2,  1,
     0,  0,  0,
    -1, -2, -1
  );
  vec4 xAgg = dot(kernel, xWeights);
  vec4 yAgg = dot(kernel, yWeights);
  outputColor = sqrt(xAgg * xAgg + yAgg * yAgg);
  outputColor.a = 1;
}
