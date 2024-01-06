package features

import (
	"github.com/GemAi/configs"
)

type FeatureDoesNotExistError struct{}

var defaultFeatures DefaultFeatures

func SetDefaultFeatures(df DefaultFeatures) {
	defaultFeatures = df
}

func (m *FeatureDoesNotExistError) Error() string {
	return "Feature does not exist"
}

func Enable(feature string, guildId string) error {
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
	return nil
}

func Disable(feature string, guildId string) error {
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
	return nil
}

func Get(guildId string) map[string]bool {
	guildConfig := configs.GetGuildParam(guildId)
	if guildConfig == nil {
		guildConfig = defaultFeatures
	}
	return guildConfig
}
