package scenemanager

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/skwb/realengo-conflict/core/event"
	"github.com/skwb/realengo-conflict/core/globals"
	boardrenderer "github.com/skwb/realengo-conflict/core/nodes/board_renderer"
	"github.com/skwb/realengo-conflict/core/nodes/gradient"
	"github.com/skwb/realengo-conflict/core/nodes/pointer"
	"github.com/skwb/realengo-conflict/core/nodes/teams"
)

type GameScene struct {
	SetSceneFunc func(scene Scene)
	Bus          *event.SignalBus
	board        *boardrenderer.Board
	gradient     *gradient.MovingGradient
	Pointer      *pointer.Pointer
}

func (g *GameScene) CheckForReroll() {
	if rl.IsKeyPressed(rl.KeyR) {
		var new_teams = teams.SortTeams(g.Bus)
		g.board = boardrenderer.NewDefaultBoard(g.Bus, new_teams)
		g.gradient.ChangeColors(new_teams[1].Colors[1], new_teams[0].Colors[1])
		globals.PointerColor = new_teams[1].Colors[0]
	}
}

func (m *GameScene) Unload() {
	m.board.Unload()

}

func (g *GameScene) Init() {
	var teams = teams.SortTeams(g.Bus)
	g.board = boardrenderer.NewDefaultBoard(g.Bus, teams)
	g.gradient = gradient.NewMovingGradient(teams[1].Colors[1], teams[0].Colors[1])
	globals.PointerColor = teams[1].Colors[0]
}

func (g *GameScene) Update() {
	g.CheckForReroll()
	g.board.HandleInput(g.Bus)
}

func (m *GameScene) GetSceneName() string {
	return "Game"
}

func (g *GameScene) DrawInScreen() {
}

func (g *GameScene) Draw() {
	g.gradient.Draw()
	g.board.Draw()
}
