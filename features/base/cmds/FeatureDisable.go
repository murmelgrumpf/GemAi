package cmds_base

import (
	"github.com/GemAi/features"
	"github.com/GemAi/utils"

	"github.com/bwmarrin/discordgo"
)

var featureIdNamesDisable map[string]string

func FeatureDisable(infos *features.FeatureInfos) *features.Cmd {
	featureIdNamesDisable = infos.FeatureIdNames
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
					Choices:     infos.FeatureChoices,
				},
			},
		},
		Function: featureDisableFunction,
	}
}

func featureDisableFunction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	feature := utils.Options(i)["feature"].StringValue()
	err := features.Disable(feature, i.GuildID)
	returnMessage := "***" + featureIdNamesDisable[feature] + "*** was disabled :white_check_mark:"
	if err != nil {
		returnMessage = "***" + featureIdNamesDisable[feature] + "*** could not be disabled :x:"
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: returnMessage,
		},
	})
}
