package ui

type UIElement interface {
	Draw()
	SetPosition(x, y int32)
	GetSize() (w, h int32)
}
