#version 330
uniform sampler2D tex0;
uniform sampler2D lastFrame;
in vec2 fragTexCoord;
out vec4 outputColor;

uniform float newColorCoeff = 0.05;
uniform float distanceThreshold = 0.01;

//meant to be similar to frei0r's "baltan" effect. 
//which itself was a clone of "StreamingTV" from an older library
void main() {
  vec4 newCol = texture(tex0,      fragTexCoord);
  vec4 oldCol = texture(lastFrame, fragTexCoord);
  outputColor = (newCol * newColorCoeff) + ((1-newColorCoeff) * oldCol);
  if(distance(oldCol, newCol) < distanceThreshold){
    outputColor = newCol;
  }
  outputColor.a = 1;
}

