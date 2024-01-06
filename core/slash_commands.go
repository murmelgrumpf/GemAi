package core

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegisterSlashCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if cmd, ok := featureInfos.CommandsMap[i.ApplicationCommandData().Name]; ok {
		cmd.Function(s, i)
	}
}

func ApplySlashCommands(s *discordgo.Session, guildId string) {
	for _, cmd := range featureInfos.Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildId, cmd.Command)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", cmd.Command.Name, err)
		}
	}
}
