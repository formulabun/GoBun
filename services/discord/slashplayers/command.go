package slashplayers

import (
	"context"
	"log"
	"time"

	dContext "GoBun/services/discord/context"
	"GoBun/services/discord/env"
	translator "GoBun/services/translator/client"

	"github.com/bwmarrin/discordgo"
)

var trClient *translator.APIClient
var logger = log.Default()

var command = &discordgo.ApplicationCommand{
	Name:        "players",
	Description: "Get player info",
}

func Start(c dContext.DiscordContext) {
	logger.SetPrefix("/players | ")
	config := translator.NewConfiguration()
	trClient = translator.NewAPIClient(config)

	request := trClient.DefaultApi.ServerinfoGet(context.Background())
	_, _, err := request.Execute()
	for err != nil {
		<-time.After(time.Second * 5)
		_, _, err = request.Execute()
	}
	logger.Println("Connection gotten, registering command")
	command, err = c.S.ApplicationCommandCreate(env.APPICATIONID, env.TESTGUILD, command)
	if err != nil {
		logger.Fatal(err)
	}

	destroy := c.S.AddHandler(reply)

	for _ = range c.Cancel {
	}

	destroy()
}

func reply(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	//interactionData := interaction.ApplicationCommandData()
	s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hello",
		},
	})
}
