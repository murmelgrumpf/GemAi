package core

import (
	"log"

	"github.com/GemAi/features"
	"github.com/GemAi/features/test"

	"github.com/bwmarrin/discordgo"
)

var allCmds = append(test.Cmds)
var allCmdsMap = make(map[string]*features.Cmd, len(allCmds))

func InitSlashCommands() {
	for _, cmd := range allCmds {
		allCmdsMap[cmd.Command.Name] = cmd
	}
}

func RegisterSlashCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if cmd, ok := allCmdsMap[i.ApplicationCommandData().Name]; ok {
		cmd.Function(s, i)
	}
}

func ApplySlashCommands(s *discordgo.Session, guildId string) {
	for _, cmd := range allCmds {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildId, cmd.Command)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", cmd.Command.Name, err)
		}
	}
}
