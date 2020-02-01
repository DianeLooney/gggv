void main() {
  gl_Position = gl_in[0].gl_Position;
  fragTexCoord = geomTexCoord[0];
  screenCoord = geomScreenCoord[0];
  particleN = geomParticleN[0];
  EmitVertex();

  gl_Position = gl_in[1].gl_Position;
  fragTexCoord = geomTexCoord[1];
  screenCoord = geomScreenCoord[1];
  particleN = geomParticleN[1];
  EmitVertex();

  gl_Position = gl_in[2].gl_Position;
  fragTexCoord = geomTexCoord[2];
  screenCoord = geomScreenCoord[2];
  particleN = geomParticleN[2];
  EmitVertex();

  EndPrimitive();
}
