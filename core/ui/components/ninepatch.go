package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type NinePatch struct {
	Tex   rl.Texture2D
	Patch rl.NPatchInfo
	Dest  rl.Rectangle
}

func NewNinePatch(tex rl.Texture2D, left, top, right, bottom int32) *NinePatch {
	return &NinePatch{
		Tex: tex,
		Patch: rl.NPatchInfo{
			Source: rl.NewRectangle(0, 0, float32(tex.Width), float32(tex.Height)),
			Left:   left,
			Top:    top,
			Right:  right,
			Bottom: bottom,
			Layout: 0,
		},
	}
}

func (n *NinePatch) SetDest(x, y, w, h int32) {
	n.Dest = rl.NewRectangle(float32(x), float32(y), float32(w), float32(h))
}

func (n *NinePatch) Draw() {
	rl.DrawTextureNPatch(
		n.Tex,
		n.Patch,
		n.Dest,
		rl.NewVector2(0, 0),
		0.0,
		rl.White,
	)
}
