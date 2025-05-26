#version 330

in vec2 fragTexCoord;
in vec4 fragColor;

uniform sampler2D texture0;
uniform vec4 colPrimary;
uniform vec4 colSecondary;
uniform vec4 newPrimary;
uniform vec4 newSecondary;

out vec4 finalColor;

void main() {
    vec4 texColor = texture(texture0, fragTexCoord);

    const float tolerance = 0.1;

    vec3 diffPrimary = abs(texColor.rgb - colPrimary.rgb);
    vec3 diffSecondary = abs(texColor.rgb - colSecondary.rgb);

    if (all(lessThan(diffPrimary, vec3(tolerance)))) {
        finalColor = vec4(newPrimary.rgb, texColor.a);
    }
    else if (all(lessThan(diffSecondary, vec3(tolerance)))) {
        finalColor = vec4(newSecondary.rgb, texColor.a);
    }
    else {
        finalColor = texColor;
    }
}
