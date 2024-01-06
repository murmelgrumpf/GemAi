package core

import (
	"github.com/GemAi/features"
	"github.com/GemAi/features/base"
	"github.com/GemAi/features/test"
	"github.com/bwmarrin/discordgo"
)

var allFeatures = [...]*features.Feature{
	&base.Feature,
	&test.Feature,
}

var featuresMap = map[string]*features.Feature{}

var featureChoices = features.FeatureChoices{}
var featureIdNames = map[string]string{}
var defaultFeatures = features.DefaultFeatures{}

var commands = []*features.Cmd{}
var commandsMap = map[string]*features.Cmd{}

var featureInfos = features.FeatureInfos{
	FeaturesMap: featuresMap,

	FeatureChoices:  featureChoices,
	FeatureIdNames:  featureIdNames,
	DefaultFeatures: defaultFeatures,

	Commands:    commands,
	CommandsMap: commandsMap,
}

func InitFeatures() {
	for _, feat := range allFeatures {
		featuresMap[feat.Id] = feat
		featureChoices = append(featureChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  feat.Name,
			Value: feat.Id,
		})

		featureIdNames[feat.Id] = feat.Name
		defaultFeatures[feat.Id] = feat.Id == "base"
	}
	features.SetDefaultFeatures(defaultFeatures)

	for _, feat := range allFeatures {

		commands = append(commands, feat.GetCommands(&featureInfos)...)

	}

	for _, command := range commands {
		commandsMap[command.Command.Name] = command
	}
}
