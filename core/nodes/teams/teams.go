package teams

import (
	"encoding/json"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"github.com/samber/lo/mutable"
	"github.com/skwb/realengo-conflict/core/event"
	"github.com/skwb/realengo-conflict/core/utils"
)

type Team struct {
	Colors    []rl.Color
	Name      string
	isCurrent bool
}

func SortTeams(bus *event.SignalBus) []Team {
	return SortTeamsFile("./assets/static/teams.json", bus)
}

func SortTeamsFile(path string, bus *event.SignalBus) []Team {
	var selected_teams []Team

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Error().Msgf("Teams file not found: %s", path)
		return createDefaultTeams(bus)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Err(err).Msg("Error reading teams file")
		return createDefaultTeams(bus)
	}

	var teamsMap map[string][]string
	err = json.Unmarshal(data, &teamsMap)
	if err != nil {
		log.Err(err).Msg("Error decoding teams JSON")
		return createDefaultTeams(bus)
	}

	teamNames := lo.Keys(teamsMap)
	if len(teamNames) < 2 {
		log.Error().Msg("Less than 2 teams found in JSON")
		return createDefaultTeams(bus)
	}

	var teamPool []Team
	for name, colorHexes := range teamsMap {
		if len(colorHexes) < 2 {
			log.Warn().Msgf("Team %s does not have at least 2 colors", name)
			continue
		}

		var colors []rl.Color
		for _, hex := range colorHexes {
			rgba, err := utils.HexToRGBA(hex)
			if err != nil {
				log.Warn().Msgf("Invalid color for team %s: %s", name, hex)
				continue
			}
			r, g, b, a := rgba.RGBA()
			colors = append(colors, rl.NewColor(uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)))
		}

		if len(colors) >= 2 {
			teamPool = append(teamPool, Team{Colors: colors, Name: name})
		}
	}

	if len(teamPool) < 2 {
		log.Warn().Msg("Could not find two valid teams in file, using defaults")
		return createDefaultTeams(bus)
	}

	const maxAttempts = 10
	for attempt := 0; attempt < maxAttempts; attempt++ {
		shuffled := teamPool
		mutable.Shuffle(shuffled)
		t1, t2 := shuffled[0], shuffled[1]

		if !teamsShareColors(t1, t2) {
			t1.isCurrent = false
			t2.isCurrent = true
			selected_teams = []Team{t1, t2}
			bus.Emit(event.SignalTeamsSorted, selected_teams)
			return selected_teams
		}
		log.Warn().Msgf("Attempt %d: Skipping team pair with overlapping colors: %s and %s", attempt+1, t1.Name, t2.Name)
	}

	log.Warn().Msg("Failed to find non-overlapping teams, using defaults")
	return createDefaultTeams(bus)
}

func teamsShareColors(t1, t2 Team) bool {
	for _, c1 := range t1.Colors {
		for _, c2 := range t2.Colors {
			if c1 == c2 {
				return true
			}
		}
	}
	return false
}

func createDefaultTeams(bus *event.SignalBus) []Team {
	log.Info().Msg("Creating default teams")

	defaultTeams := []Team{
		{Name: "Red Team", Colors: []rl.Color{rl.NewColor(200, 50, 50, 255), rl.NewColor(255, 150, 150, 255)}},
		{Name: "Blue Team", Colors: []rl.Color{rl.NewColor(50, 50, 200, 255), rl.NewColor(150, 150, 255, 255)}},
	}

	bus.Emit(event.SignalTeamsSorted, defaultTeams)
	return defaultTeams
}
