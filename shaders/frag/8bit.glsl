#version 330

//default uniforms from GGGV

//futher sources are set as tex1, tex2, tex3
uniform sampler2D tex0;

uniform float hres = 320;
uniform float vres = 200;
uniform vec3 color0 = vec3(187, 187, 187) / 255.0;
uniform vec3 color1 = vec3(0, 136, 255) / 255.0;
uniform vec3 color2 = vec3(170, 255, 102) / 255.0;
uniform vec3 color3 = vec3(119, 119, 119) / 255.0;
uniform vec3 color4 = vec3(51, 51, 51) / 255.0;
uniform vec3 color5 = vec3(255, 119, 119) / 255.0;
uniform vec3 color6 = vec3(102, 68, 0) / 255.0;
uniform vec3 color7 = vec3(221, 136, 85) / 255.0;
uniform vec3 color8 = vec3(238, 238, 119) / 255.0;
uniform vec3 color9 = vec3(0, 0, 170) / 255.0;
uniform vec3 colorA = vec3(0, 204, 85) / 255.0;
uniform vec3 colorB = vec3(204, 68, 204) / 255.0;
uniform vec3 colorC = vec3(170, 255, 238) / 255.0;
uniform vec3 colorD = vec3(136, 0, 0) / 255.0;
uniform vec3 colorE = vec3(255, 255, 255) / 255.0;
uniform vec3 colorF = vec3(0, 0, 0) / 255.0;

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

void main() {
	vec2 ftc = vec2(
		floor(fragTexCoord.x * hres + 0.5) / hres,
		floor(fragTexCoord.y * vres + 0.5) / vres);
	vec3 actual = texture(tex0, ftc).rgb;
	float diff0 = distance(actual, color0);
	float diff1 = distance(actual, color1);
	float diff2 = distance(actual, color2);
	float diff3 = distance(actual, color3);
	float diff4 = distance(actual, color4);
	float diff5 = distance(actual, color5);
	float diff6 = distance(actual, color6);
	float diff7 = distance(actual, color7);
	float diff8 = distance(actual, color8);
	float diff9 = distance(actual, color9);
	float diffA = distance(actual, colorA);
	float diffB = distance(actual, colorB);
	float diffC = distance(actual, colorC);
	float diffD = distance(actual, colorD);
	float diffE = distance(actual, colorE);
	float diffF = distance(actual, colorF);
	
	outputColor = vec4(0, 0, 0, 1);
	float diff = diff0; outputColor = vec4(color0, 1);
	if (diff1 < diff) { diff = diff1; outputColor = vec4(color1, 1); }
	if (diff2 < diff) { diff = diff2; outputColor = vec4(color2, 1); }
	if (diff3 < diff) { diff = diff3; outputColor = vec4(color3, 1); }
	if (diff4 < diff) { diff = diff4; outputColor = vec4(color4, 1); }
	if (diff5 < diff) { diff = diff5; outputColor = vec4(color5, 1); }
	if (diff6 < diff) { diff = diff6; outputColor = vec4(color6, 1); }
	if (diff7 < diff) { diff = diff7; outputColor = vec4(color7, 1); }
	if (diff8 < diff) { diff = diff8; outputColor = vec4(color8, 1); }
	if (diff9 < diff) { diff = diff9; outputColor = vec4(color9, 1); }
	if (diffA < diff) { diff = diffA; outputColor = vec4(colorA, 1); }
	if (diffB < diff) { diff = diffB; outputColor = vec4(colorB, 1); }
	if (diffC < diff) { diff = diffC; outputColor = vec4(colorC, 1); }
	if (diffD < diff) { diff = diffD; outputColor = vec4(colorD, 1); }
	if (diffE < diff) { diff = diffE; outputColor = vec4(colorE, 1); }
	if (diffF < diff) { diff = diffF; outputColor = vec4(colorF, 1); }
}
