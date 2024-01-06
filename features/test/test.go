package test

import (
	"github.com/GemAi/features"
	"github.com/GemAi/features/test/cmds"
)

var Feature = features.Feature{
	Id:          "test",
	Name:        "Test",
	Description: "Testcommands, can be ignored",
	Emoji:       "gear",
	GetCommands: func(infos *features.FeatureInfos) []*features.Cmd {
		return []*features.Cmd{cmds_test.Ping(infos)}
	},
}
