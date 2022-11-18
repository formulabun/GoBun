package slashplayers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"GoBun/functional/array"
	dContext "GoBun/services/discord/context"
	"GoBun/services/discord/env"
	translator "GoBun/services/translator/client"

	"github.com/bwmarrin/discordgo"
)

var trClient *translator.APIClient
var logger = log.New(os.Stdout, "/players | ", log.LstdFlags)
var request translator.ApiPlayerinfoGetRequest

var command = &discordgo.ApplicationCommand{
	Name:        "players",
	Description: "Get player info",
}

func Start(c dContext.DiscordContext) {
	config := translator.NewConfiguration()
	trClient = translator.NewAPIClient(config)

	logger.Println("Trying to get a connection to the translator service.")
	request = trClient.DefaultApi.PlayerinfoGet(context.Background())
	_, _, err := request.Execute()
	for err != nil {
		<-time.After(time.Second * 5)
		_, _, err = request.Execute()
		logger.Println(err)
	}
	logger.Println("Connection gotten, registering command.")

	command, err = c.S.ApplicationCommandCreate(env.APPICATIONID, env.TESTGUILD, command)
	if err != nil {
		logger.Fatal(err)
	}

	destroy := c.S.AddHandler(reply)

	for _ = range c.Cancel {
	}

	logger.Println("Destroying command")
	destroy()
}

func reply(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	logger.Println("Interaction created")
	data, _, _ := request.Execute()

	players := array.Filter(data.PlayerInfo, func(p translator.PlayerInfoEntry) bool {
		return *p.Team != 255
	})
	playerNames := array.Map(players, func(p translator.PlayerInfoEntry) string {
		return p.GetName()
	})
	spectators := array.Filter(data.PlayerInfo, func(p translator.PlayerInfoEntry) bool {
		return *p.Team == 255
	})
	spectatorNames := array.Map(spectators, func(p translator.PlayerInfoEntry) string {
		return p.GetName()
	})
	response := formatResponse(playerNames, spectatorNames)

	logger.Printf("Sending %s\n", response)
	s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%s", response),
		},
	})
	logger.Println("Interaction Ending")
}
