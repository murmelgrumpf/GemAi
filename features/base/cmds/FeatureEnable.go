package cmds_base

import (
	"github.com/GemAi/features"
	"github.com/GemAi/utils"

	"github.com/bwmarrin/discordgo"
)

func FeatureEnable() *features.Cmd {
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
					Choices:     features.Infos.FeatureChoices,
				},
			},
		},
		Function: featureEnableFunction,
	}
}

func featureEnableFunction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	feature := utils.Options(i)["feature"].StringValue()
	err := features.Enable(feature, i.GuildID, s)
	returnMessage := "***" + features.Infos.FeatureIdNames[feature] + "*** was enabled :white_check_mark:"
	if err != nil {
		returnMessage = "***" + features.Infos.FeatureIdNames[feature] + "*** could not be enabled :x:"
	}
	utils.InteractionRespond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: returnMessage,
		},
	})
}
