package features

import (
	"github.com/bwmarrin/discordgo"
)

type CmdFunc func(*discordgo.Session, *discordgo.InteractionCreate)

type Cmd struct {
	Command  *discordgo.ApplicationCommand
	Function CmdFunc
}
