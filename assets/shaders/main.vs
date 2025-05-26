#version 330 core

// Atributos de entrada
layout(location = 0) in vec3 vertexPosition;
layout(location = 1) in vec2 vertexTexCoord;

// Saída para o fragment shader
out vec2 fragTexCoord;

// Uniformes
uniform mat4 projection;
uniform mat4 modelview;

void main() {
    UV = vertexTexCoord; // Passa as coordenadas de textura
    gl_Position = projection * modelview * vec4(vertexPosition, 1.0);
}