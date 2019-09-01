#version 330

uniform float time;
uniform sampler2D tex0;

in vec2 fragTexCoord;
out vec4 outputColor;

float PHI = 1.61803398874989484820459 * 00000.1; // Golden Ratio   
float PI  = 3.14159265358979323846264 * 00000.1; // PI
float SQ2 = 1.41421356237309504880169 * 10000.0; // Square Root of Two
float gold_noise(in vec2 coordinate, in float seed){
    return fract(tan(distance(coordinate*(seed+PHI), vec2(PHI, PI)))*SQ2);
}

vec4 randither(vec4 input, vec2 co, float seed){
    vec3 thresh = vec3(gold_noise(co+input.rg,seed),gold_noise(co+input.br,seed),gold_noise(co-input.rb,seed));
    vec4 ret = input;
    if(input.r < thresh.r){
        ret.r = 0;
    }
    else{
        ret.r = 1;
    }
    if(input.g < thresh.g){
        ret.g = 0;
    }
    else{
        ret.g = 1;
    }
    if(input.b < thresh.b){
        ret.b = 0;
    }
    else{
        ret.b = 1;
    }

    return ret;
}

void main(void)
{
    outputColor=texture(tex0,fragTexCoord);
    outputColor = randither(outputColor, fragTexCoord, time);//randither(outputColor);
}
