package status

import (
	"context"
	"fmt"
	"log"
	"time"

	dContext "GoBun/services/discord/context"
	translator "GoBun/services/translator/client"

	"github.com/bwmarrin/discordgo"
)

var trClient *translator.APIClient

func Start(c dContext.DiscordContext) {
	ticker := time.NewTicker(5*time.Second)

	config := translator.NewConfiguration()
	trClient = translator.NewAPIClient(config)

	go setupTimer(ticker, c.S)

	for _ = range c.Cancel {
	}
	ticker.Stop()
}

func setupTimer(tick *time.Ticker, s *discordgo.Session) {
	for _ = range tick.C {
		updateStatus(s)
	}
}

func makeStatusData(info *translator.ServerInfo) discordgo.UpdateStatusData {
	i := 0
	var statusText string
	playerCount := int(info.GetNumberOfPlayer())
	switch playerCount {
	case 0:
		statusText = "an empty map"
	case 1:
		statusText = fmt.Sprintf("%d player race", playerCount)
	default:
		statusText = fmt.Sprintf("%d players race", playerCount)
	}

	status := discordgo.Activity{
		statusText,
		discordgo.ActivityTypeWatching,
		"",
		time.Now(),
		"",
		"",
		"",
		discordgo.TimeStamps{},
		discordgo.Emoji{},
		discordgo.Party{},
		discordgo.Assets{},
		discordgo.Secrets{},
		false,
		0,
	}

	return discordgo.UpdateStatusData{
		&i,
		[]*discordgo.Activity{&status},
		false,
		"Hello",
	}
}

func updateStatus(s *discordgo.Session) {
	resp, _, err := trClient.DefaultApi.ServerinfoGet(context.Background()).Execute()
	if err != nil {
		log.Print(fmt.Sprintf("Could not get player count: %s", err))
		return
	}
	err = s.UpdateStatusComplex(makeStatusData(resp))
	if err != nil {
		log.Print(fmt.Sprintf("Could not update status: %s", err))
		return
	}
}
