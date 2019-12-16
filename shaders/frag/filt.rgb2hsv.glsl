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

//neato code from the lolengine ppl
//http://lolengine.net/blog/2013/07/27/rgb-to-hsv-in-glsl#
vec3 rgb2hsv(vec3 c)
{
  vec4 K = vec4(0.0, -1.0 / 3.0, 2.0 / 3.0, -1.0);
  vec4 p = mix(vec4(c.bg, K.wz), vec4(c.gb, K.xy), step(c.b, c.g));
  vec4 q = mix(vec4(p.xyw, c.r), vec4(c.r, p.yzx), step(p.x, c.r));

  float d = q.x - min(q.w, q.y);
  float e = 1.0e-10;
  return vec3(abs(q.z + (q.w - q.y) / (6.0 * d + e)), d / (q.x + e), q.x);
}

void main() {
  vec4 inp = texture(tex0, fragTexCoord);
  vec3 hsv = rgb2hsv(inp.rgb);
  outputColor = vec4(hsv.r,hsv.g,hsv.b,inp.a);
}
