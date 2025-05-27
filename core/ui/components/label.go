package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Label struct {
	Text     string
	X, Y     int32
	FontSize float32
	Color    rl.Color
}

func (l *Label) Draw() {
	rl.DrawText(l.Text, l.X, l.Y, int32(l.FontSize), l.Color)
}

func (l *Label) SetPosition(x, y int32) {
	l.X = x
	l.Y = y
}

func (l *Label) GetSize() (int32, int32) {
	w := int32(rl.MeasureText(l.Text, int32(l.FontSize)))
	return w, int32(l.FontSize)
}

func NewLabel(text string) *Label {
	return &Label{
		Text:     text,
		FontSize: 12,
		Color:    rl.Black,
	}
}
