package cmds_base

import (
	"github.com/GemAi/features"
	"github.com/GemAi/utils"

	"github.com/bwmarrin/discordgo"
)

func FeatureDisable() *features.Cmd {
	return &features.Cmd{
		Command: &discordgo.ApplicationCommand{
			Name:        "feature-disable",
			Description: "Disables the specified feature",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "feature",
					Description: "Feature",
					Required:    true,
					Choices:     features.Infos.FeatureChoices,
				},
			},
		},
		Function: featureDisableFunction,
	}
}

func featureDisableFunction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	feature := utils.Options(i)["feature"].StringValue()
	err := features.Disable(feature, i.GuildID, s)
	returnMessage := "***" + features.Infos.FeatureIdNames[feature] + "*** was disabled :white_check_mark:"
	if err != nil {
		returnMessage = "***" + features.Infos.FeatureIdNames[feature] + "*** could not be disabled :x:"
	}
	utils.InteractionRespond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: returnMessage,
		},
	})
}
