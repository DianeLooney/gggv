void main() {
  float weight0 = sin(time + 0)+1.3;
  float weight1 = sin(time + 1)+1.3;
  float weight2 = sin(time + 2)+1.3;
  float weightT = weight0 + weight1 + weight2;

  vec4 centerPos = (
      (weight1 * gl_in[0].gl_Position) +
      (weight2 * gl_in[1].gl_Position) +
      (weight0 * gl_in[2].gl_Position)) / weightT;

  vec2 centerTex = (geomTexCoord[0] + geomTexCoord[1] + geomTexCoord[2]) / 3;

  gl_Position = gl_in[0].gl_Position;
  fragTexCoord = geomTexCoord[1];
  EmitVertex();
  gl_Position = gl_in[1].gl_Position; 
  fragTexCoord = centerTex;
  EmitVertex();
  gl_Position = centerPos;
  fragTexCoord = geomTexCoord[0];
  EmitVertex();
  EndPrimitive();


  gl_Position = gl_in[1].gl_Position; 
  fragTexCoord = geomTexCoord[2];
  EmitVertex();
  gl_Position = gl_in[2].gl_Position; 
  fragTexCoord = centerTex;
  EmitVertex();
  gl_Position = centerPos;
  fragTexCoord = geomTexCoord[1];
  EmitVertex();
  EndPrimitive();

  gl_Position = gl_in[2].gl_Position; 
  fragTexCoord = geomTexCoord[0];
  EmitVertex();
  gl_Position = gl_in[0].gl_Position;
  fragTexCoord = centerTex;
  EmitVertex();
  gl_Position = centerPos;
  fragTexCoord = geomTexCoord[2];
  EmitVertex();
  EndPrimitive();
}  
