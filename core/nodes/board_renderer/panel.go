package board_renderer

import (
	"fmt"
	"strconv"

	"github.com/skwb/realengo-conflict/core/globals"
	"github.com/skwb/realengo-conflict/core/ui/components"
)

func (b *Board) DrawPanel() {
	if b.hoveredCell == nil {
		return
	}

	var titleText string
	var descriptionText string

	if b.hoveredCell.PieceType != NoPiece {
		if b.hoveredCell.EnemyPiece {
			titleText = "Enemy "
		} else {
			titleText = "Your "
		}
		titleText += PieceToString(b.hoveredCell.PieceType)
		descriptionText = fmt.Sprintf("This cell is occupied by %s", PieceToString(b.hoveredCell.PieceType))
	} else {
		titleText = "Empty cell"
		descriptionText = "This cell is empty, there is nothing here. Maybe you want to use it."
	}

	titleLabel := components.NewLabel(titleText)
	numberLabel := components.NewLabel("#" + strconv.Itoa(b.hoveredCell.Number))

	descriptionLabel := components.NewLabel(descriptionText)

	header := &components.Container{
		Layout:  components.Horizontal,
		Spacing: 2,
	}
	header.Add(numberLabel)
	header.Add(titleLabel)

	panel := &components.Container{
		X:       0,
		Y:       0,
		Layout:  components.Vertical,
		Spacing: 2,
	}
	panel.Add(header)
	panel.Add(descriptionLabel)
	np := components.NewNinePatch(globals.ContainerNPatchTexture, 23, 12, 23, 23)
	panel.SetBackground(np)
	panel.Draw()
}
