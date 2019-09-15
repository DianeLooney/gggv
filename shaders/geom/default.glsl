#version 150 core

layout (triangles) in;
layout (triangle_strip, max_vertices = 3) out;

in vec2 geomTexCoord[];
out vec2 fragTexCoord;

void main() {    
    gl_Position = gl_in[0].gl_Position;
    fragTexCoord = geomTexCoord[0];
    EmitVertex();
    gl_Position = gl_in[1].gl_Position; 
    fragTexCoord = geomTexCoord[1];
    EmitVertex();
    gl_Position = gl_in[2].gl_Position; 
    fragTexCoord = geomTexCoord[2];
    EmitVertex();
    EndPrimitive();
}  
