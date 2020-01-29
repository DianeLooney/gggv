#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

uniform float time;
uniform float windowHeight;
uniform float windowWidth;

in vec3 vert;
in vec2 vertTexCoord;
in float vertParticleN;

out vec2 geomTexCoord;
out vec2 geomScreenCoord;
out float geomParticleN;
