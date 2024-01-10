package features

import (
	"github.com/GemAi/configs"
	"github.com/GemAi/utils"
	"github.com/bwmarrin/discordgo"
)

type FeatureDoesNotExistError struct{}

var defaultFeatures DefaultFeatures
var Infos *FeatureInfos

func SetDefaultFeatures(df DefaultFeatures) {
	defaultFeatures = df
}

func SetFeatureInfos(fi *FeatureInfos) {
	Infos = fi
}

func (m *FeatureDoesNotExistError) Error() string {
	return "Feature does not exist"
}

func Enable(feature string, guildId string, s *discordgo.Session) error {
	guildConfig := configs.GetGuildParam(guildId)
	if guildConfig == nil {
		guildConfig = defaultFeatures
	}
	_, featureExists := guildConfig[feature]
	if !featureExists {
		return &FeatureDoesNotExistError{}
	}
	guildConfig[feature] = true
	configs.SetGuildParam(guildId, guildConfig)

	for _, cmd := range Infos.FeaturesMap[feature].Commands {
		registeredCmd, err := utils.RegisterSlashCommand(s, guildId, cmd.Command)
		if err != nil {
			return err
		}
		cmd.Command = registeredCmd
	}
	return nil
}

func Disable(feature string, guildId string, s *discordgo.Session) error {
	guildConfig := configs.GetGuildParam(guildId)
	if guildConfig == nil {
		guildConfig = defaultFeatures
	}
	_, featureExists := guildConfig[feature]
	if !featureExists || feature == "base" {
		return &FeatureDoesNotExistError{}
	}
	guildConfig[feature] = false
	configs.SetGuildParam(guildId, guildConfig)
	for _, cmd := range Infos.FeaturesMap[feature].Commands {
		err := utils.UnRegisterSlashCommandsFeature(s, guildId, cmd.Command.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func Get(guildId string) map[string]bool {
	guildConfig := configs.GetGuildParam(guildId)
	if guildConfig == nil {
		guildConfig = defaultFeatures
	}
	return guildConfig
}

func GetEnabled(feature string, guildId string) bool {
	return Get(guildId)[feature]
}

func Toggle(feature string, guildId string, s *discordgo.Session) error {
	if GetEnabled(feature, guildId) {
		return Disable(feature, guildId, s)
	}
	return Enable(feature, guildId, s)
}
