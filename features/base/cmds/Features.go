package cmds_base

import (
	"github.com/GemAi/features"

	"github.com/bwmarrin/discordgo"
)

func Features(_ *features.FeatureInfos) *features.Cmd {
	return &features.Cmd{
		Command: &discordgo.ApplicationCommand{
			Name:        "features",
			Description: "Shows all the available features",
		},
		Function: featuresFunction,
	}
}

func featuresFunction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong",
		},
	})
}
