package board_renderer

import (
	"log"
	"math/rand"
	"os"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (b *Board) DrawPieces() {
	primaryValues := []float32{
		float32(b.PiecePrimaryColor.R) / 255,
		float32(b.PiecePrimaryColor.G) / 255,
		float32(b.PiecePrimaryColor.B) / 255,
		float32(b.PiecePrimaryColor.A) / 255,
	}
	rl.SetShaderValue(b.paletteShader, b.primaryColorLoc, primaryValues, rl.ShaderUniformVec4)

	secondaryValues := []float32{
		float32(b.PieceSecondaryColor.R) / 255,
		float32(b.PieceSecondaryColor.G) / 255,
		float32(b.PieceSecondaryColor.B) / 255,
		float32(b.PieceSecondaryColor.A) / 255,
	}
	rl.SetShaderValue(b.paletteShader, b.secondaryColorLoc, secondaryValues, rl.ShaderUniformVec4)

	for _, cell := range b.cells {
		if cell.PieceType == NoPiece {
			continue
		}

		var (
			textureSet rl.Texture2D
			teamColors []rl.Color
		)

		if cell.EnemyPiece {
			textureSet = b.enemySet
			if len(b.teams) > 0 {
				teamColors = b.teams[0].Colors
			}
		} else {
			textureSet = b.playerSet
			if len(b.teams) > 1 {
				teamColors = b.teams[1].Colors
			}
		}

		if len(teamColors) < 2 {
			continue
		}

		newPrimary := []float32{
			float32(teamColors[0].R) / 255,
			float32(teamColors[0].G) / 255,
			float32(teamColors[0].B) / 255,
			float32(teamColors[0].A) / 255,
		}
		rl.SetShaderValue(b.paletteShader, b.newPrimaryColorLoc, newPrimary, rl.ShaderUniformVec4)

		newSecondary := []float32{
			float32(teamColors[1].R) / 255,
			float32(teamColors[1].G) / 255,
			float32(teamColors[1].B) / 255,
			float32(teamColors[1].A) / 255,
		}
		rl.SetShaderValue(b.paletteShader, b.newSecondaryColorLoc, newSecondary, rl.ShaderUniformVec4)
		frameWidth := float32(textureSet.Width) / 6
		source := rl.NewRectangle(
			float32((int32(cell.PieceType)-1)*int32(frameWidth)),
			0,
			30,
			28,
		)
		rl.BeginShaderMode(b.paletteShader)
		rl.DrawTexturePro(
			textureSet,
			source,
			cell.Rectangle,
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)
		rl.EndShaderMode()
	}
}

func (b *Board) loadNames() {
	data, err := os.ReadFile("./assets/static/names.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	names := strings.Split(strings.TrimSpace(string(data)), "\n")

	for i, cell := range b.cells {
		if cell.PieceType == NoPiece {
			continue
		}
		randomIndex := rand.Intn(len(names))
		b.cells[i].PieceName = names[randomIndex]
	}
}

func (b *Board) CreatePieces() {
	piecesOrder := []PieceType{Spy, Wizard, Knight, FactionLeader, Knight, Wizard, Spy, Archer, Soldier, Soldier, Soldier, Soldier, Soldier, Archer, Soldier, NoPiece, NoPiece, NoPiece, NoPiece, NoPiece, Soldier}

	for i := range b.cells {
		b.cells[i].PieceType = NoPiece
		b.cells[i].EnemyPiece = false
	}

	for row := 0; row < 3; row++ {
		for col := 0; col < b.columns; col++ {
			cellIndex := row*b.columns + col
			if cellIndex < len(piecesOrder) {
				b.cells[cellIndex].PieceType = piecesOrder[cellIndex]
				b.cells[cellIndex].EnemyPiece = true
			}
		}
	}

	for row := 7; row < 10; row++ {
		for col := 0; col < b.columns; col++ {
			mirroredRow := 9 - row
			mirroredCol := b.columns - 1 - col
			mirroredIndex := mirroredRow*b.columns + mirroredCol
			if mirroredIndex < len(piecesOrder) {
				cellIndex := row*b.columns + col
				b.cells[cellIndex].PieceType = piecesOrder[mirroredIndex]
				b.cells[cellIndex].EnemyPiece = false
			}
		}
	}
}
