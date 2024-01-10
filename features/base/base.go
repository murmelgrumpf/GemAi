package base

import (
	"github.com/GemAi/features"
	"github.com/GemAi/features/base/cmds"
)

var Feature = features.Feature{
	Id:          "base",
	Name:        "Base",
	Description: "The Base Commands. You can't disable this feature",
	Emoji:       "🧱",
	GetCommands: func() []*features.Cmd {
		return []*features.Cmd{
			cmds_base.FeatureEnable(),
			cmds_base.FeatureDisable(),
			cmds_base.Features(),
		}
	},
}
