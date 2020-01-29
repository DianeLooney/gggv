#version 330

//default uniforms from GGGV

//futher sources are set as tex1, tex2, tex3
uniform sampler2D tex0;
uniform sampler2D tex1;
uniform sampler2D tex2;
uniform sampler2D tex3;
uniform sampler2D tex4;
uniform sampler2D tex5;
uniform sampler2D tex6;
uniform sampler2D tex7;
uniform sampler2D tex8;
uniform sampler2D tex9;

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
in vec2 screenCoord;
in float particleN;

//output pixel color
out vec4 outputColor;

#define PI 3.1415926535897932384626433832795

vec2 screenToPolar(vec2 screenCoords) {
  screenCoords -= vec2(0.5, 0.5);
  screenCoords *= vec2(windowWidth, windowHeight);
  return vec2(
    length(screenCoords),
    atan(screenCoords.y, screenCoords.x)
  );
}

vec2 polarToScreen(vec2 polarCoords) {
  return vec2(0.5, 0.5) + vec2(
    polarCoords.x * cos(polarCoords.y),
    polarCoords.x * sin(polarCoords.y)
  ) / vec2(windowWidth, windowHeight);
}

vec2 cToP(vec2 coord) {
  return vec2(
    length(coord),
    atan(coord.y, coord.x)
  );
}

vec2 pToC(vec2 coord) {
  return vec2(
    cos(coord.y) * coord.x,
    sin(coord.y) * coord.x
  );
}

float noise1(vec2 st) {
    return fract(sin(dot(st.xy,vec2(12.9898,78.233)) * 43758.5453123));
}

vec2 noise2(vec2 st) {
    return fract(sin(st.xy * vec2(12.9898,78.233)) * 43758.5453123);
}
/*
vec2 noise2(vec2 coord, float seed) {
  return fract(1264.2135 * sin((coord + seed) * 17237.28348));
}*/

vec2 ftc = (fragTexCoord - 0.5) * vec2(windowWidth / windowHeight, 1);
