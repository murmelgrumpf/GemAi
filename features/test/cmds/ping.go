package cmds_test

import (
	"github.com/GemAi/features"
	"github.com/GemAi/utils"

	"github.com/bwmarrin/discordgo"
)

func Ping() *features.Cmd {
	return &features.Cmd{
		Command: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Sends pong back",
		},
		Function: pingFunction,
	}
}

func pingFunction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	utils.InteractionRespond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong",
		},
	})
}
