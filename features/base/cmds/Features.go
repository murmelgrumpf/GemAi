package cmds_base

import (
	"strings"

	"github.com/GemAi/features"
	"github.com/GemAi/utils"

	"github.com/bwmarrin/discordgo"
)

func Features() *features.Cmd {
	return &features.Cmd{
		Command: &discordgo.ApplicationCommand{
			Name:        "features",
			Description: "Shows all the available features",
		},
		Function:               featuresFunction,
		MsgInteractionFunction: featuresMessageInteractionFunction,
	}
}

func featuresFunction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	utils.InteractionRespond(s, i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: generateResponseData(s, i),
	})
}

func generateResponseData(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	componentsFirstRow := []discordgo.MessageComponent{}
	embedFields := []*discordgo.MessageEmbedField{}

	for _, feature := range features.Infos.Features {
		componentsFirstRow = append(componentsFirstRow, discordgo.Button{
			Label:    feature.Name,
			Style:    utils.BoolButtonStye(features.GetEnabled(feature.Id, i.GuildID)),
			CustomID: "features|" + feature.Id,
			Disabled: feature.Id == "base",
			Emoji:    utils.DefaultEmoji(feature.Emoji),
		})

		embedFields = append(embedFields, &discordgo.MessageEmbedField{
			Name:  "[ " + feature.Emoji + " ] " + feature.Name + " - *(" + enabledDisabledBool(features.GetEnabled(feature.Id, i.GuildID)) + ")*",
			Value: feature.Description,
		})
	}
	components := []discordgo.MessageComponent{discordgo.ActionsRow{Components: componentsFirstRow}}

	return &discordgo.InteractionResponseData{
		Components: components,
		Embeds: []*discordgo.MessageEmbed{
			{
				Type:        discordgo.EmbedTypeRich,
				Title:       "Features",
				Description: "Enable or disable available features.",
				Color:       utils.EmbedColorInfo,
				Fields:      embedFields,
			},
		},
	}
}

func featuresMessageInteractionFunction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	splitId := strings.Split(i.MessageComponentData().CustomID, "|")
	err := features.Toggle(splitId[1], i.GuildID, s)
	if err != nil {
		utils.SendError(s, i, err.Error())
		return
	}

	newData := generateResponseData(s, i)
	utils.InteractionRespond(
		s, i,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Components: newData.Components,
				Embeds:     newData.Embeds,
			},
		},
	)
}

func enabledDisabledBool(enabled bool) string {
	if enabled {
		return "enabled"
	}
	return "disabled"
}
