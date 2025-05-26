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
	var signal_bus = event.NewSignalBus()
	var manager = &scenemanager.SceneManager{}
	var cfg, _ = config.LoadConfig()
	rl.InitWindow(cfg.Window.WindowWidth, cfg.Window.WindowHeight, cfg.Game.GameName)
	globals.ContainerNPatchTexture = rl.LoadTexture("./assets/sprites/npatchs/container.png")
	var pointer = pointer.NewPointer(rl.LoadTexture("./assets/sprites/pointer/pressed.png"), rl.LoadTexture("./assets/sprites/pointer/regular.png"))
	rl.SetWindowIcon(*rl.LoadImage("assets/sprites/gameicon.png"))
	var virtualWidth = cfg.Window.ViewportWidth
	var virtualHeight = cfg.Window.ViewportHeight
	var menu *scenemanager.MenuScene = &scenemanager.MenuScene{
		Bus: signal_bus,
		SetSceneFunc: func(s scenemanager.Scene) {
			manager.SetScene(s)
		},
	}
	manager.SetScene(menu)
	defer rl.CloseWindow()
	var viewport = rl.LoadRenderTexture(virtualWidth, virtualHeight)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		manager.Update()
		DefaultShortcuts()
		rl.BeginTextureMode(viewport)
		rl.ClearBackground(rl.Black)
		manager.Draw()
		pointer.DrawPointer()
		rl.EndTextureMode()

		screenWidth := float32(rl.GetScreenWidth())
		screenHeight := float32(rl.GetScreenHeight())
		screenRatio := screenWidth / screenHeight
		targetRatio := float32(virtualWidth) / float32(virtualHeight)

		var scale, destWidth, destHeight, destX, destY float32
		if screenRatio >= targetRatio {
			scale = screenHeight / float32(virtualHeight)
			destWidth = float32(virtualWidth) * scale
			destHeight = screenHeight
			destX = (screenWidth - destWidth) / 2
			destY = 0
		} else {
			scale = screenWidth / float32(virtualWidth)
			destWidth = screenWidth
			destHeight = float32(virtualHeight) * scale
			destX = 0
			destY = (screenHeight - destHeight) / 2
		}

		src := rl.Rectangle{X: 0, Y: 0, Width: float32(virtualWidth), Height: -float32(virtualHeight)}
		dest := rl.Rectangle{X: destX, Y: destY, Width: destWidth, Height: destHeight}
		origin := rl.Vector2{X: 0, Y: 0}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawTexturePro(viewport.Texture, src, dest, origin, 0, rl.White)

		manager.DrawInScreen()

		rl.EndDrawing()
	}
	rl.UnloadTexture(globals.ContainerNPatchTexture)
	pointer.Unload()
}
