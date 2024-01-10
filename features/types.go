package features

import (
	"github.com/bwmarrin/discordgo"
)

type InteractionFunc func(*discordgo.Session, *discordgo.InteractionCreate)

type Cmd struct {
	Command                *discordgo.ApplicationCommand
	Function               InteractionFunc
	MsgInteractionFunction InteractionFunc
}

type Feature struct {
	Id          string
	Name        string
	Description string
	Emoji       string
	GetCommands func() []*Cmd
	Commands    []*Cmd
}

type FeatureInfos struct {
	Features    *[2]*Feature
	FeaturesMap map[string]*Feature

	CommandsMap map[string]*Cmd

	FeatureChoices  FeatureChoices
	FeatureIdNames  map[string]string
	DefaultFeatures DefaultFeatures
}

type FeatureChoices []*discordgo.ApplicationCommandOptionChoice
type DefaultFeatures map[string]bool
