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

const float PI = 3.1415926535897932384626433832795;

void main() {
    vec4 sample = texture(tex0, fragTexCoord);
    float rhigh = 0.5 + 0.3 * sin(time / 10);
    float rlow  = 0.5 - 0.3 * sin(time / 10);
    float ghigh = 0.5 + 0.3 * sin(time / 10+2*PI/3);
    float glow  = 0.5 - 0.3 * sin(time / 10+2*PI/3);
    float bhigh = 0.5 + 0.3 * sin(time / 10+4*PI/3);
    float blow  = 0.5 - 0.3 * sin(time / 10+4*PI/3);

    if ((sample.r > rlow && sample.r < rhigh) ||
        (sample.g > glow && sample.g < ghigh) || 
        (sample.b > blow && sample.b < bhigh)) {
        outputColor = vec4(1,1,1,1);
    } else {
        outputColor = vec4(0,0,0,1);
    }
}
