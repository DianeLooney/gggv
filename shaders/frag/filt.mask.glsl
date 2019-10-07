#version 330

//default uniforms from GGGV

//futher sources are set as tex1, tex2, tex3
uniform sampler2D tex0;
uniform sampler2D tex1;
uniform sampler2D tex2;

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
    vec4 mask = texture(tex0, fragTexCoord);
    vec4 sample1 = texture(tex1, fragTexCoord);
    vec4 sample2 = texture(tex2, fragTexCoord);
    outputColor = mask*sample1 + (vec4(1,1,1,1) - mask) * sample2;
}
