package cmds_test

import (
	"fmt"

	"github.com/GemAi/features"

	"github.com/bwmarrin/discordgo"
)

func Ping(_ *features.FeatureInfos) *features.Cmd {
	return &features.Cmd{
		Command: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Sends pong back",
		},
		Function: pingFunction,
	}
}

func pingFunction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	fmt.Println("hallo")

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong",
		},
	})
}
