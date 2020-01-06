void main() {
    vec4 sample = texture(tex0, fragTexCoord);
    float rhigh = 0.5 + 0.3 * sin(time / 10);
    float rlow  = 0.5 - 0.3 * sin(time / 10);
    float ghigh = 0.5 + 0.3 * sin(time / 10+2*PI/3);
    float glow  = 0.5 - 0.3 * sin(time / 10+2*PI/3);
    float bhigh = 0.5 + 0.3 * sin(time / 10+4*PI/3);
    float blow  = 0.5 - 0.3 * sin(time / 10+4*PI/3);

    if ((sample.r > rlow && sample.r < rhigh) ||
        (sample.g > glow && sample.g < ghigh) || 
        (sample.b > blow && sample.b < bhigh)) {
        outputColor = vec4(1,1,1,1);
    } else {
        outputColor = vec4(0,0,0,1);
    }
}
