vec4 pride(vec4 c, vec3 c1, vec3 c2, vec3 c3, float w1, float w2, float w3) {
              c1 = hsv2rgb(c1);
              c2 = hsv2rgb(c2);
              c3 = hsv2rgb(c3);
              float x = rgb2hsv(c.rgb).z * (w1 + w2 + w3);
              if (x<w1) return vec4(c1, 1);
              if (x<w1 + w2) return vec4(c2, 1);
              return vec4(c3, 1);
            }
uniform float u83;
vec4 store(vec4 c, layout(rgba8) image2D storage) {
              imageStore(storage, iftc, c);
              return c;
            }
vec2 scale(vec2 xy, float r) { return xy * vec2(r); }
uniform float u80;
layout(rgba8) uniform image2D x;
vec4 osc(vec2 xy) { return vec4(.5 * (1 + cos(xy)), 0, 1); }
uniform float u75;
vec2 pulsate(vec2 xy, float r, float ex) {
             vec2 p = cToP(xy);
             p.x -= 3 * pow(abs( cos(p.x - time*r)), ex);
             return pToC(p);
            }
uniform float u81;
vec4 from(vec2 xy, layout(rgba8) image2D storage) {
              vec2 _xy = fract(xy);
              return imageLoad(
                storage,
                ivec2(_xy*vec2(windowWidth, windowHeight))
              );
            }
vec4 setOutputColor(vec4 c) { outputColor = c; return c; }
vec2 rotate(vec2 xy, float theta) { return pToC(cToP(xy) + vec2(0, 2*PI*theta)); }
uniform float u73;
uniform vec3 u78;
uniform vec3 u77;
uniform float u74;
uniform vec3 u79;
uniform float u82;
uniform float u76;
void main() {
store(pride(osc(scale(rotate(pulsate(ftc, u73, u74), u75), u76)), u77, u78, u79, u80, u81, u82), x);
setOutputColor(from(scale(ftc, u83), x));
};