package kcg

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Ghq   bool
	Repos map[string]*RepositoryConfig
}

type RepositoryConfig struct {
	Repo   string
	Path   string
	Groups []string
	Setup  []string
}

func WriteConfig(path string) {
	if path != "" {
		viper.SetConfigFile(path)
	}

	if viper.ConfigFileUsed() == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.SetConfigFile(filepath.Join(home, ".kcg"))
	}

	viper.SetTypeByDefaultValue(false)
	viper.WriteConfig()
}
