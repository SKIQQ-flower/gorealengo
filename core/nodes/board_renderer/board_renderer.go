package board_renderer

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/skwb/realengo-conflict/core/config"
	event "github.com/skwb/realengo-conflict/core/event"
	teams "github.com/skwb/realengo-conflict/core/nodes/teams"
	utils "github.com/skwb/realengo-conflict/core/utils"
)

var cfg *config.Config

type PieceType int

const (
	NoPiece PieceType = iota
	FactionLeader
	Knight
	Wizard
	Spy
	Archer
	Soldier
)

func PieceToString(piece PieceType) string {
	switch piece {
	case FactionLeader:
		return "Faction Leader"
	case Knight:
		return "Knight"
	case Wizard:
		return "Wizard"
	case Spy:
		return "Spy"
	case Archer:
		return "Archer"
	case Soldier:
		return "Soldier"
	default:
		return "No Piece"
	}
}

type Cell struct {
	Number      int
	PieceType   PieceType
	Rectangle   rl.Rectangle
	PieceName   string
	EnemyPiece  bool
	Hovered     bool
	IsBlackTile bool
}

type Board struct {
	PrimaryColor         rl.Color
	SecondaryColor       rl.Color
	CenterColor          rl.Color
	PiecePrimaryColor    rl.Color
	PieceSecondaryColor  rl.Color
	teams                []teams.Team
	ChessBoardSize       float32
	CellSize             int
	BoardWidth           int
	BoardHeight          int
	OutlineSize          float32
	rows                 int
	columns              int
	cells                []*Cell
	hoveredCell          *Cell
	enemySet             rl.Texture2D
	playerSet            rl.Texture2D
	paletteShader        rl.Shader
	primaryColorLoc      int32
	secondaryColorLoc    int32
	newPrimaryColorLoc   int32
	newSecondaryColorLoc int32
}

func (b *Board) Unload() {
	rl.UnloadTexture(b.enemySet)
	rl.UnloadTexture(b.playerSet)
	rl.UnloadShader(b.paletteShader)
}

func NewDefaultBoard(bus *event.SignalBus, teams []teams.Team) *Board {
	if cfg == nil {
		cfg, _ = config.LoadConfig()
	}
	primaryColor, _ := utils.HexToRGBA("#2c1e31")
	secondaryColor, _ := utils.HexToRGBA("#f7f3b7")
	piecesPrimaryColor, _ := utils.HexToRGBA("#ac2847")
	piecesSecondaryColor, _ := utils.HexToRGBA("#2c1e31")
	paletteShader := rl.LoadShader("", "./assets/shaders/palettechanger.fs")

	board := &Board{
		PrimaryColor:        rl.Color(primaryColor),
		SecondaryColor:      rl.Color(secondaryColor),
		PiecePrimaryColor:   rl.Color(piecesPrimaryColor),
		PieceSecondaryColor: rl.Color(piecesSecondaryColor),
		CenterColor:         utils.BlendColor(primaryColor, secondaryColor, 0.5),
		ChessBoardSize:      1,
		CellSize:            30,
		BoardWidth:          (7 * 30),
		BoardHeight:         (10 * 30),
		OutlineSize:         2,
		rows:                10,
		columns:             7,
		enemySet:            rl.LoadTexture("./assets/sprites/charactersets/default.png"),
		playerSet:           rl.LoadTexture("./assets/sprites/charactersets/default.png"),
		paletteShader:       paletteShader,
		teams:               teams,
	}
	rl.SetTextureFilter(board.enemySet, rl.TextureFilterNearest)
	rl.SetTextureFilter(board.playerSet, rl.TextureFilterNearest)
	board.primaryColorLoc = rl.GetShaderLocation(paletteShader, "colPrimary")
	board.secondaryColorLoc = rl.GetShaderLocation(paletteShader, "colSecondary")
	board.newPrimaryColorLoc = rl.GetShaderLocation(paletteShader, "newPrimary")
	board.newSecondaryColorLoc = rl.GetShaderLocation(paletteShader, "newSecondary")
	board.Create()
	return board
}

func (b *Board) centerOffset() rl.Vector2 {
	return rl.NewVector2(
		float32(cfg.Window.ViewportWidth)/2-float32(b.BoardWidth)/2,
		float32(cfg.Window.ViewportHeight)/2-float32(b.BoardHeight)/2,
	)
}

func (b *Board) CreateCells() {
	b.cells = nil
	offset := b.centerOffset()

	scaledCellSize := float32(b.CellSize) * b.ChessBoardSize

	for row := 0; row < b.rows; row++ {
		for col := 0; col < b.columns; col++ {
			cellIndex := row*b.columns + col
			x := float32(col) * scaledCellSize
			y := float32(row) * scaledCellSize

			cell := Cell{
				Number:      cellIndex,
				IsBlackTile: !((row+col)%2 == 0),
				Rectangle: rl.NewRectangle(
					x+offset.X,
					y+offset.Y,
					scaledCellSize,
					scaledCellSize,
				),
			}

			if cellIndex == 31 {
				cell.Rectangle.Height = scaledCellSize * 2
				cell.IsBlackTile = false
			}

			b.cells = append(b.cells, &cell)
		}
	}
}

func (b *Board) Create() {
	b.CreateCells()
	b.CreatePieces()
	b.loadNames()
}

func (b *Board) DrawCells() {
	offset := b.centerOffset()
	shadowOffset := rl.NewVector2(offset.X*1.05, offset.Y*1.45)
	rl.DrawRectangle(int32(shadowOffset.X), int32(shadowOffset.Y), int32(b.BoardWidth), int32(b.BoardHeight), b.PrimaryColor)

	for _, cell := range b.cells {
		if cell.Number == 38 {
			continue
		}

		var color rl.Color
		switch {
		case cell.Number == 31:
			color = b.CenterColor
		case cell.IsBlackTile:
			color = b.PrimaryColor
		default:
			color = b.SecondaryColor
		}

		rl.DrawRectangle(int32(cell.Rectangle.X), int32(cell.Rectangle.Y), int32(cell.Rectangle.Width), int32(cell.Rectangle.Height), color)
	}

	if b.hoveredCell != nil {
		var outlineColor rl.Color
		switch b.hoveredCell.IsBlackTile {
		case true:
			outlineColor = utils.InvertColor(b.PrimaryColor)
		default:
			outlineColor = utils.InvertColor(b.SecondaryColor)
		}
		rl.DrawRectangleLinesEx(b.hoveredCell.Rectangle, b.OutlineSize, outlineColor)
	}
	rl.DrawRectangleLinesEx(utils.RectAddPadding(rl.NewRectangle(offset.X, offset.Y, float32(b.BoardWidth), float32(b.BoardHeight)), b.OutlineSize), b.OutlineSize, rl.NewColor(255, 255, 255, 255))
}

func (b *Board) HandleInput(bus *event.SignalBus) {
	mousePos := utils.GetViewportMousePosition()

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && b.hoveredCell != nil && rl.CheckCollisionPointRec(mousePos, b.hoveredCell.Rectangle) {
		bus.Emit(event.SignalCellPressed, b.hoveredCell)
	}

	found := false
	for _, cell := range b.cells {
		if rl.CheckCollisionPointRec(mousePos, cell.Rectangle) {
			if b.hoveredCell != cell {
				if b.hoveredCell != nil {
					bus.Emit(event.SignalCellHovered, b.hoveredCell)
					b.hoveredCell.Hovered = false
				}
				b.hoveredCell = cell
				bus.Emit(event.SignalCellUnhovered, cell)
				cell.Hovered = true
			}
			found = true
			break
		}
	}

	if !found && b.hoveredCell != nil {
		bus.Emit(event.SignalCellUnhovered, b.hoveredCell)
		b.hoveredCell.Hovered = false
		b.hoveredCell = nil
	}
}

func (b *Board) Draw() {
	b.DrawCells()
	b.DrawPieces()
	b.DrawPanel()
}
