package features

import (
	"github.com/bwmarrin/discordgo"
)

type CmdFunc func(*discordgo.Session, *discordgo.InteractionCreate)

type Cmd struct {
	Command  *discordgo.ApplicationCommand
	Function CmdFunc
}

type Feature struct {
	Id          string
	Name        string
	Description string
	Emoji       string
	GetCommands func(*FeatureInfos) []*Cmd
}

type FeatureInfos struct {
	FeaturesMap map[string]*Feature

	Commands    []*Cmd
	CommandsMap map[string]*Cmd

	FeatureChoices  FeatureChoices
	FeatureIdNames  map[string]string
	DefaultFeatures DefaultFeatures
}

type FeatureChoices []*discordgo.ApplicationCommandOptionChoice
type DefaultFeatures map[string]bool
