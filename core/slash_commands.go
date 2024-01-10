package core

import (
	"log"

	"github.com/GemAi/features"
	"github.com/bwmarrin/discordgo"
)

func RegisterSlashCommands(s *discordgo.Session, guildId string) {
	guildFeatures := features.Get(guildId)
	for featureId, enabled := range guildFeatures {
		if !enabled {
			continue
		}
		for _, cmd := range featureInfos.FeaturesMap[featureId].Commands {
			registerdCmd, err := s.ApplicationCommandCreate(s.State.User.ID, guildId, cmd.Command)
			if err != nil {
				log.Panicf("Cannot create '%v' command: %v", cmd.Command.Name, err)
			}
			cmd.Command = registerdCmd
		}
	}
}
