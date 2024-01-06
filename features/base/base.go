package base

import (
	"github.com/GemAi/features"
	"github.com/GemAi/features/base/cmds"
)

var Feature = features.Feature{
	Id:          "base",
	Name:        "Base",
	Description: "The Base Commands. You can't disable this feature",
	Emoji:       "brick",
	GetCommands: func(infos *features.FeatureInfos) []*features.Cmd {
		return []*features.Cmd{
			cmds_base.FeatureEnable(infos),
			cmds_base.FeatureDisable(infos),
			cmds_base.Features(infos),
		}
	},
}
