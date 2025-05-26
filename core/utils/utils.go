package utils

import (
	"fmt"
	"image/color"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/skwb/realengo-conflict/core/config"
)

func InvertColor(color rl.Color) rl.Color {
	return rl.Color{R: 255 - color.R, G: 255 - color.G, B: 255 - color.B, A: color.A}
}

func RectAddPadding(rect rl.Rectangle, padding float32) rl.Rectangle {
	return rl.Rectangle{
		X:      rect.X - padding,
		Y:      rect.Y - padding,
		Width:  rect.Width + 2*padding,
		Height: rect.Height + 2*padding,
	}
}

func GetViewportMousePosition() rl.Vector2 {
	screenMousePos := rl.GetMousePosition()
	var cfg, _ = config.LoadConfig()

	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	viewportWidth := float32(cfg.Window.ViewportWidth)
	viewportHeight := float32(cfg.Window.ViewportHeight)

	screenRatio := screenWidth / screenHeight
	viewportRatio := viewportWidth / viewportHeight

	var scale float32
	var offsetX, offsetY float32

	if screenRatio >= viewportRatio {
		scale = screenHeight / viewportHeight
		offsetX = (screenWidth - viewportWidth*scale) / 2
		offsetY = 0
	} else {
		scale = screenWidth / viewportWidth
		offsetX = 0
		offsetY = (screenHeight - viewportHeight*scale) / 2
	}

	x := (screenMousePos.X - offsetX) / scale
	y := (screenMousePos.Y - offsetY) / scale

	return rl.NewVector2(x, y)
}

func BlendColor(a, b rl.Color, t float32) rl.Color {
	return rl.Color{
		R: uint8(float32(a.R)*(1-t) + float32(b.R)*t),
		G: uint8(float32(a.G)*(1-t) + float32(b.G)*t),
		B: uint8(float32(a.B)*(1-t) + float32(b.B)*t),
		A: uint8(float32(a.A)*(1-t) + float32(b.A)*t),
	}
}

func ScreenToViewport(screenPos rl.Vector2, cfg *config.Config) rl.Vector2 {
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	viewportWidth := float32(cfg.Window.ViewportWidth)
	viewportHeight := float32(cfg.Window.ViewportHeight)

	return rl.NewVector2(screenPos.X*viewportWidth/screenWidth, screenPos.Y*viewportHeight/screenHeight)
}

func HexToRGBA(hex string) (color.RGBA, error) {
	hex = strings.TrimPrefix(hex, "#")

	var r, g, b, a uint8 = 0, 0, 0, 255

	switch len(hex) {
	case 3:
		if _, err := fmt.Sscanf(hex, "%1x%1x%1x", &r, &g, &b); err != nil {
			return color.RGBA{}, err
		}
		r *= 17
		g *= 17
		b *= 17
	case 6:
		if _, err := fmt.Sscanf(hex, "%2x%2x%2x", &r, &g, &b); err != nil {
			return color.RGBA{}, err
		}
	case 8:
		if _, err := fmt.Sscanf(hex, "%2x%2x%2x%2x", &r, &g, &b, &a); err != nil {
			return color.RGBA{}, err
		}
	default:
		return color.RGBA{}, fmt.Errorf("invalid hex format")
	}

	return color.RGBA{R: r, G: g, B: b, A: a}, nil
}
