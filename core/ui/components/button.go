package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	Text                string
	X, Y, Width, Height int32
	OnClick             func()
	Hovered             bool
}

func (b *Button) Draw() {
	mouse := rl.GetMousePosition()
	if rl.CheckCollisionPointRec(mouse, rl.NewRectangle(float32(b.X), float32(b.Y), float32(b.Width), float32(b.Height))) {
		b.Hovered = true
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && b.OnClick != nil {
			b.OnClick()
		}
	} else {
		b.Hovered = false
	}

	color := rl.DarkGray
	if b.Hovered {
		color = rl.Gray
	}

	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, color)
	rl.DrawText(b.Text, b.X+10, b.Y+10, 20, rl.White)
}

func (b *Button) SetPosition(x, y int32) {
	b.X = x
	b.Y = y
}

func (b *Button) GetSize() (int32, int32) {
	return b.Width, b.Height
}
