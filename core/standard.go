package core

import rl "github.com/gen2brain/raylib-go/raylib"

func DefaultShortcuts() {
	if rl.IsKeyDown(rl.KeyLeftAlt) && rl.IsKeyDown(rl.KeyEnter) {
		rl.ToggleBorderlessWindowed()
	}
}
