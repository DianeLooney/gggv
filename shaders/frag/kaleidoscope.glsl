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

uniform float flakes = 4;
uniform float radius = 300;

void main() {
    vec2 coords = fragTexCoord;
    coords = screenToPolar(coords);
    coords.y += time * 0.1;
	coords.x -= 1 * time;
	coords.x -= 1080 * time;
	coords.x = mod(coords.x, 1080);
	if (coords.x > radius) {
		coords.x = 2 * radius - coords.x;
	}
	if (coords.x < 0) {
		coords.x = -coords.x;
	}
    coords.y = mod(coords.y, PI / flakes);
	coords.x += (coords.y / 50) *  cos(coords.y * 8);
    if (coords.y > PI / (2*flakes)) coords.y = PI / flakes - coords.y;
    coords = polarToScreen(coords);
    outputColor = texture(tex0, coords);
}
