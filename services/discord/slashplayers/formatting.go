package slashplayers

import (
	bunString "GoBun/functional/strings"
	"fmt"
)

func joinNames(names []string, verb string) string {
	if len(names) == 0 {
		return ""
	}
	if len(names) == 1 {
		return fmt.Sprintf("%s is %s.", names[0], verb)
	}
	last := names[len(names)-1]
	init := names[:len(names)-1]
	return fmt.Sprintf("%s and %s are %s.", bunString.Join(init, ", "), last, verb)
}

func formatResponse(players, spectators []string) string {
	playerPart := joinNames(players, "racing")
	spectatorPart := joinNames(spectators, "watching")

	if playerPart == "" && spectatorPart == "" {
		return "Nobody is playing."
	}

	response := playerPart + " " + spectatorPart
	response = bunString.Escape(response, '\\', '\\')
	response = bunString.Escape(response, '*', '\\')
	response = bunString.Escape(response, '_', '\\')
	response = bunString.Escape(response, '|', '\\')
	response = bunString.Escape(response, ':', '\\')
	return response
}
