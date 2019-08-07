
#version 330

uniform sampler2D tex;
uniform float time;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    float color = 0.0;
	color += sin( fragTexCoord.x * cos( time / 15.0 ) * 80.0 ) + cos( fragTexCoord.y * cos( time / 15.0 ) * 10.0 );
	color += sin( fragTexCoord.y * sin( time / 10.0 ) * 40.0 ) + cos( fragTexCoord.x * sin( time / 25.0 ) * 40.0 );
	color += sin( fragTexCoord.x * sin( time / 5.0 ) * 10.0 ) + sin( fragTexCoord.y * sin( time / 35.0 ) * 80.0 );
	color *= sin( time / 10.0 ) * 0.5;

    outputColor = 0.25 * (2 + sin(time*10.0)) * texture(tex, fragTexCoord) + 0.25 * (2 - sin(time / 3.0)) *vec4( vec3( color, color * 0.5, sin( color + time / 3.0 ) * 0.75 ), 1.0 );
}
