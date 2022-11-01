package context

import "github.com/bwmarrin/discordgo"

type DiscordContext struct {
	S      *discordgo.Session
	Cancel chan struct{}
}
