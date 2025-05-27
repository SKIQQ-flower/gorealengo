package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/skwb/realengo-conflict/core/config"
	"github.com/skwb/realengo-conflict/core/event"
	"github.com/skwb/realengo-conflict/core/globals"
	"github.com/skwb/realengo-conflict/core/nodes/pointer"
	scenemanager "github.com/skwb/realengo-conflict/core/scenes"
)

func StartGame() {
	signalBus := event.NewSignalBus()
	manager := &scenemanager.SceneManager{}
	cfg, _ := config.LoadConfig()

	rl.InitWindow(cfg.Window.WindowWidth, cfg.Window.WindowHeight, cfg.Game.GameName)
	defer rl.CloseWindow()

	virtualW := float32(cfg.Window.ViewportWidth)
	virtualH := float32(cfg.Window.ViewportHeight)

	renderTexture := rl.LoadRenderTexture(int32(virtualW), int32(virtualH))
	defer rl.UnloadRenderTexture(renderTexture)

	globals.ContainerNPatchTexture = rl.LoadTexture("./assets/sprites/npatchs/container.png")
	defer rl.UnloadTexture(globals.ContainerNPatchTexture)

	p := pointer.NewPointer(
		rl.LoadTexture("./assets/sprites/pointer/pressed.png"),
		rl.LoadTexture("./assets/sprites/pointer/regular.png"),
	)
	defer p.Unload()

	rl.SetWindowIcon(*rl.LoadImage("assets/sprites/gameicon.png"))
	rl.SetTargetFPS(60)

	menu := &scenemanager.MenuScene{
		Bus: signalBus,
		SetSceneFunc: func(s scenemanager.Scene) {
			manager.SetScene(s)
		},
	}
	manager.SetScene(menu)

	for !rl.WindowShouldClose() {
		screenW := float32(rl.GetScreenWidth())
		screenH := float32(rl.GetScreenHeight())

		manager.Update()
		DefaultShortcuts()

		rl.BeginTextureMode(renderTexture)
		rl.ClearBackground(rl.Black)

		manager.Draw()
		p.DrawPointer()

		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		src := rl.NewRectangle(0, 0, virtualW, -virtualH)
		dst := rl.NewRectangle(0, 0, screenW, screenH)
		rl.DrawTexturePro(renderTexture.Texture, src, dst, rl.NewVector2(0, 0), 0, rl.White)

		manager.DrawInScreen()
		rl.EndDrawing()
	}
}
