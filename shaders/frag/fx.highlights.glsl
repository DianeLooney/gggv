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

uniform float fadeRate = 0.01;
uniform float threshold = 0.8;

uniform float highlightR = 255;
uniform float highlightG = 255;
uniform float highlightB = 255;

void main() {

    outputColor = texture(lastFrame, fragTexCoord) - vec4(fadeRate, fadeRate, fadeRate, 0);
    vec4 col = texture(tex0, fragTexCoord);
    if (col.r > threshold && col.g > threshold && col.b > threshold) {
        outputColor = vec4(highlightR/255.0,highlightG/255.0,highlightB/255.0,1);
    }

    if (outputColor.r < 0) outputColor.r = 0;
    if (outputColor.g < 0) outputColor.g = 0;
    if (outputColor.b < 0) outputColor.b = 0;
    if (outputColor.a > 1) outputColor.a = 1;
    if (outputColor.r > 1) outputColor.r = 1;
    if (outputColor.g > 1) outputColor.g = 1;
    if (outputColor.b > 1) outputColor.b = 1;
    if (outputColor.a > 1) outputColor.a = 1;
}
