package slashaoc

import (
	dContext "GoBun/services/discord/context"
	"GoBun/services/discord/env"
	"bytes"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

type aocData struct {
	content string
}

var logger = log.New(os.Stdout, "/AoC | ", log.LstdFlags)
var nonePermission = int64(0)

var data *members = nil
var dataLock = sync.Mutex{}
var staleLock = sync.Mutex{}

var command = &discordgo.ApplicationCommand{
	Name:                     "aoc",
	Description:              "Get Advent of Code private leaderboard",
	DefaultMemberPermissions: &nonePermission,
}

func Start(c dContext.DiscordContext) {
	logger.Println("Registering command")
	_, err := c.S.ApplicationCommandCreate(env.APPICATIONID, env.TESTGUILD, command)
	if err != nil {
		logger.Printf("Couldn't register my command: \n", err)
		return
	}

	destroy := c.S.AddHandler(reply)

	for _ = range c.Cancel {
	}

	logger.Println("Destroying command")
	destroy()
}

func timeout() {
	if staleLock.TryLock() {
		defer staleLock.Unlock()
		<-time.After(30 * time.Minute)
		dataLock.Lock()
		data = nil
		dataLock.Unlock()
	}
}

func reply(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	go timeout()

	dataLock.Lock()
	if data == nil {
		d, _ := fetchScores()
		data = &d
	}

	response := format(data)
	dataLock.Unlock()

	go s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &response,
	})
}

func format(data *members) discordgo.InteractionResponseData {
	var result bytes.Buffer
	var embed discordgo.MessageEmbed
	embed.Title = ":christmas_tree: Formulabun Advent of Code leaderboard :christmas_tree:"
	embed.Fields = make([]*discordgo.MessageEmbedField, len(*data))
	embed.Color = 0x62ff29
	embed.URL = "https://adventofcode.com/2022/"
	embed.Description = "Join our leaderboard with `2605039`"
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: "Happy coding!",
	}
	for i, m := range *data {
		var field = discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%s :star: %d", m.Name, m.Stars),
			Value:  fmt.Sprintf("%d", m.Score),
			Inline: false,
		}
		embed.Fields[i] = &field
	}

	return discordgo.InteractionResponseData{
		Content: result.String(),
		Embeds:  []*discordgo.MessageEmbed{&embed},
	}
}
