float PHI = 1.61803398874989484820459 * 00000.1; // Golden Ratio
float SQ2 = 1.41421356237309504880169 * 10000.0; // Square Root of Two
float gold_noise(in vec2 coordinate, in float seed){
  return fract(tan(distance(coordinate*(seed+PHI), vec2(PHI, PI)))*SQ2);
}

vec4 randither(vec4 inColor, vec2 co, float seed){
  vec3 thresh = vec3(
    gold_noise(co+inColor.rg,seed),
    gold_noise(co+inColor.br,seed),
    gold_noise(co-inColor.rb,seed));
  vec4 ret = inColor;
  if(inColor.r < thresh.r){
    ret.r = 0;
  }
  else{
    ret.r = 1;
  }
  if(inColor.g < thresh.g){
    ret.g = 0;
  }
  else{
    ret.g = 1;
  }
  if(inColor.b < thresh.b){
    ret.b = 0;
  }
  else{
    ret.b = 1;
  }

  return ret;
}

void main(void)
{
  outputColor = texture(tex0,fragTexCoord);
  outputColor = randither(outputColor, fragTexCoord, time);//randither(outputColor);
}
