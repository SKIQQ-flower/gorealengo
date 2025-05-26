package scenemanager

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/skwb/realengo-conflict/core/event"
)

type MenuScene struct {
	SetSceneFunc func(scene Scene)
	Bus          *event.SignalBus
}

func (m *MenuScene) Init() {

}

func (m *MenuScene) GetSceneName() string {
	return "Menu"
}

func (m *MenuScene) Unload() {

}

func (m *MenuScene) Update() {
	if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
		m.SetSceneFunc(&GameScene{
			Bus:          m.Bus,
			SetSceneFunc: m.SetSceneFunc,
		})
	}
}

func (m *MenuScene) DrawInScreen() {

}

func (m *MenuScene) Draw() {
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Menu - Pressione ENTER para começar", 100, 200, 20, rl.Black)
}
