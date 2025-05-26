package gradient

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/skwb/realengo-conflict/core/config"
)

type MovingGradient struct {
	color1         rl.Color
	color2         rl.Color
	rect           rl.Rectangle
	time           float32
	amplitude      float32
	frequency      float32
	shader         rl.Shader
	topColorLoc    int32
	bottomColorLoc int32
	offsetLoc      int32
}

func NewMovingGradient(player_color rl.Color, enemy_color rl.Color) *MovingGradient {
	var cfg, _ = config.LoadConfig()
	var gradient = &MovingGradient{
		color1:    player_color,
		color2:    enemy_color,
		time:      0.0,
		amplitude: 2.5,
		frequency: 0.5,
		rect:      rl.NewRectangle(0, 0, float32(cfg.Window.ViewportWidth), float32(cfg.Window.ViewportHeight)),
		shader:    rl.LoadShader("", "./assets/shaders/gradient.fs"),
	}

	gradient.topColorLoc = rl.GetShaderLocation(gradient.shader, "topColor")
	gradient.bottomColorLoc = rl.GetShaderLocation(gradient.shader, "bottomColor")
	gradient.offsetLoc = rl.GetShaderLocation(gradient.shader, "offset")
	rl.SetShaderValue(gradient.shader, gradient.topColorLoc, []float32{
		float32(enemy_color.R) / 255.0,
		float32(enemy_color.G) / 255.0,
		float32(enemy_color.B) / 255.0,
		1.0,
	}, rl.ShaderUniformVec4)

	rl.SetShaderValue(gradient.shader, gradient.bottomColorLoc, []float32{
		float32(player_color.R) / 255.0,
		float32(player_color.G) / 255.0,
		float32(player_color.B) / 255.0,
		1.0,
	}, rl.ShaderUniformVec4)
	return gradient
}

func (m *MovingGradient) Draw() {
	m.time += rl.GetFrameTime()

	var offset float32 = float32(math.Sin(float64(m.time)*(math.Pi*2)*float64(m.frequency))) * m.amplitude / 50.0
	rl.SetShaderValue(m.shader, m.offsetLoc, []float32{offset}, rl.ShaderUniformFloat)
	rl.BeginShaderMode(m.shader)

	rl.ClearBackground(rl.Black)
	rl.DrawRectangleRec(m.rect, rl.White)

	rl.EndShaderMode()
}

func (m *MovingGradient) ChangeColors(player_color rl.Color, enemy_color rl.Color) {
	rl.SetShaderValue(m.shader, m.topColorLoc, []float32{
		float32(enemy_color.R) / 255.0,
		float32(enemy_color.G) / 255.0,
		float32(enemy_color.B) / 255.0,
		1.0,
	}, rl.ShaderUniformVec4)

	rl.SetShaderValue(m.shader, m.bottomColorLoc, []float32{
		float32(player_color.R) / 255.0,
		float32(player_color.G) / 255.0,
		float32(player_color.B) / 255.0,
		1.0,
	}, rl.ShaderUniformVec4)
}
