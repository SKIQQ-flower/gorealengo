#version 330 core

in vec2 fragTexCoord;
out vec4 fragColor;

uniform vec4 topColor;
uniform vec4 bottomColor;
uniform float offset;

void main() {
    float center = 0.5 + offset;
    float range = 0.11;

    float t = 0.0;
    
    if (fragTexCoord.y < center - range) {
        fragColor = topColor;
        return;
    } else if (fragTexCoord.y > center + range) {
        fragColor = bottomColor;
        return;
    } else {
        t = (fragTexCoord.y - (center - range)) / (2.0 * range);
        fragColor = mix(topColor, bottomColor, t);
    }
}
