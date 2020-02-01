#version 150 core

layout(triangles) in;
layout(triangle_strip, max_vertices = 3) out;

uniform float time;
uniform float windowHeight;
uniform float windowWidth;

in vec2 geomTexCoord[];
in vec2 geomScreenCoord[];
in float geomParticleN[];

out vec2 fragTexCoord;
out vec2 screenCoord;
out float particleN;
