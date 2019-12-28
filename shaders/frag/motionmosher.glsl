#version 330
uniform sampler2D tex0;
uniform sampler2D lastFrame;
uniform float tex0Width;
uniform float tex0Height;
uniform float time;
uniform float renderTime;
uniform float fps;
uniform float windowHeight;
uniform float windowWidth;
uniform float cursorX;
uniform float cursorY;
in vec2 fragTexCoord;
out vec4 outputColor;

float triggerThreshold    = 0.01;
float similarityThreshold = 0.02;
float x_step = 1/windowWidth;
float y_step = 1/windowHeight;

void main() {
  vec4 last = texture(lastFrame, fragTexCoord);
  vec4 curr = texture(tex0,      fragTexCoord);
  if(distance(last,curr) > triggerThreshold){
    for(int i=-1;i<=1;i++){
      for(int j=-1;j<=1;j++){
        vec2 newCoord = fragTexCoord;
        newCoord.x += i*x_step;
        newCoord.y += j*y_step;
        vec4 old_guy = texture(lastFrame, newCoord);
        if(distance(old_guy,curr) < similarityThreshold){
          outputColor = old_guy;
          return;
        }
        else{
          outputColor = last;
        }
      }
    }
  }
  else{
    outputColor = curr;
  }
  outputColor.a = 1;
}

