package utils

import (
	"github.com/bwmarrin/discordgo"
)

func Options(i *discordgo.InteractionCreate) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	return optionMap
}

func RegisterSlashCommand(s *discordgo.Session, guildId string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	return s.ApplicationCommandCreate(s.State.User.ID, guildId, cmd)
}

func UnRegisterSlashCommandsFeature(s *discordgo.Session, guildId string, cmdId string) error {
	err := s.ApplicationCommandDelete(s.State.User.ID, guildId, cmdId)
	return err
}
