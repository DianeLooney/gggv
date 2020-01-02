#version 150 core

layout (triangles) in;
layout (triangle_strip, max_vertices = 3) out;

uniform float windowHeight;
uniform float windowWidth;

in vec2 geomTexCoord[];
in vec2 geomScreenCoord[];
out vec2 fragTexCoord;
out vec2 screenCoord;

void main() {
  gl_Position = gl_in[0].gl_Position;
  fragTexCoord = geomTexCoord[0];
  screenCoord = geomScreenCoord[0];
  EmitVertex();

  gl_Position = gl_in[1].gl_Position; 
  fragTexCoord = geomTexCoord[1];
  screenCoord = geomScreenCoord[1];
  EmitVertex();

  gl_Position = gl_in[2].gl_Position; 
  fragTexCoord = geomTexCoord[2];
  screenCoord = geomScreenCoord[2];
  EmitVertex();

  EndPrimitive();
}
