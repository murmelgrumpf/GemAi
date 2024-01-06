package cmds_base

import (
	"github.com/GemAi/features"
	"github.com/GemAi/utils"

	"github.com/bwmarrin/discordgo"
)

var featureIdNamesEnable map[string]string

func FeatureEnable(infos *features.FeatureInfos) *features.Cmd {
	featureIdNamesEnable = infos.FeatureIdNames
	return &features.Cmd{
		Command: &discordgo.ApplicationCommand{
			Name:        "feature-enable",
			Description: "Enables the specified feature",
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
		Function: featureEnableFunction,
	}
}

func featureEnableFunction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	feature := utils.Options(i)["feature"].StringValue()
	err := features.Enable(feature, i.GuildID)
	returnMessage := "***" + featureIdNamesEnable[feature] + "*** was enabled :white_check_mark:"
	if err != nil {
		returnMessage = "***" + featureIdNamesEnable[feature] + "*** could not be enabled :x:"
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: returnMessage,
		},
	})
}
