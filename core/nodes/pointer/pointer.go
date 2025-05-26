package pointer

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/skwb/realengo-conflict/core/globals"
	"github.com/skwb/realengo-conflict/core/utils"
)

type Pointer struct {
	regular_sprite  rl.Texture2D
	pressed_sprite  rl.Texture2D
	lastPressedTime float64
}

func NewPointer(texture_regular rl.Texture2D, texture_pressed rl.Texture2D) Pointer {
	return Pointer{
		regular_sprite: texture_regular,
		pressed_sprite: texture_pressed,
	}
}

func (p *Pointer) Unload() {
	rl.UnloadTexture(p.regular_sprite)
	rl.UnloadTexture(p.pressed_sprite)
}

func (p *Pointer) DrawPointer() {
	rl.HideCursor()
	mousepos := utils.GetViewportMousePosition()
	now := rl.GetTime()

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		p.lastPressedTime = now
	}

	sprite := p.regular_sprite
	if now-p.lastPressedTime < 0.1 {
		sprite = p.pressed_sprite
	}

	rl.DrawTexture(sprite, int32(mousepos.X), int32(mousepos.Y), globals.PointerColor)
}
