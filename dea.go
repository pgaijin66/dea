package dea

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type emailProviders struct {
	allowList []string `mapstructure:"allowList.providers"`
	blockList []string `mapstructure:"blockList.providers"`
}

func loadConfig() (emailProviders, error) {

	viper.SetConfigName("dea")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return emailProviders{}, fmt.Errorf("error loading config file: %s", err)
	}

	var providers emailProviders

	err = viper.Unmarshal(&providers)
	if err != nil {
		return emailProviders{}, fmt.Errorf("fatal error config file: %s", err)
	}

	providers.allowList = viper.GetStringSlice("allowList.providers")
	providers.blockList = viper.GetStringSlice("blockList.providers")

	return providers, nil
}

func IsDisposableEmail(email string) (bool, error) {
	providers, _ := loadConfig()

	// Check if provider is in list of blocked email providers
	for _, blocked := range providers.blockList {
		if strings.HasSuffix(email, blocked) {
			// If provider is in blocked list, check if provider is explicitly allowed
			for _, allowed := range providers.allowList {
				if strings.HasSuffix(email, allowed) {
					return false, nil
				} else {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
