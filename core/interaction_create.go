package core

import (
	"strings"

	"github.com/GemAi/utils"
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Interaction.Type == discordgo.InteractionApplicationCommand {
		if cmd, ok := featureInfos.CommandsMap[i.ApplicationCommandData().Name]; ok {
			cmd.Function(s, i)
			return
		}
		utils.SendError(s, i, "This command does not exist. I don't even know how you called it.")
		return
	}
	if i.Interaction.Type == discordgo.InteractionMessageComponent {
		splitId := strings.Split(i.MessageComponentData().CustomID, "|")
		if cmd, ok := featureInfos.CommandsMap[splitId[0]]; ok && cmd.MsgInteractionFunction != nil {
			cmd.MsgInteractionFunction(s, i)
			return
		}
		utils.SendError(s, i, "This messageinteraction does not exist. I don't even know how you called it.")
		return
	}
}
