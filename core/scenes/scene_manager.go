package scenemanager

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/skwb/realengo-conflict/core/config"
	"github.com/skwb/realengo-conflict/core/event"
)

type Scene interface {
	Update()
	Draw()
	Unload()
	Init()
	DrawInScreen()
	GetSceneName() string
}

type SceneManager struct {
	Bus     *event.SignalBus
	current Scene
}

func (s *SceneManager) SetScene(scene Scene) {
	var cfg, _ = config.LoadConfig()
	if scene != nil {
		rl.SetWindowTitle(cfg.Game.GameName + " : " + scene.GetSceneName())
		if s.current != nil {
			s.current.Unload()
		}
		scene.Init()
		s.current = scene
	}
}

func (s *SceneManager) Init() {
	if s.current != nil {
		s.current.Init()
	}
}

func (s *SceneManager) Update() {
	if s.current != nil {
		s.current.Update()
	}
}

func (s *SceneManager) Draw() {
	if s.current != nil {
		s.current.Draw()
	}
}

func (s *SceneManager) DrawInScreen() {
	if s.current != nil {
		s.current.DrawInScreen()
	}
}
