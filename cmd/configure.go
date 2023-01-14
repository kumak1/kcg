/*
Copyright © 2023 kumak1

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/kumak1/kcg/kcg"
	kcgExec "github.com/kumak1/kcg/kcg/exec"
	"github.com/spf13/viper"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Operate config file",
	Long:  `Operate config file`,
}

var configureInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty config file",
	Long:  `Create an empty config file`,
	Run: func(cmd *cobra.Command, args []string) {
		initRepo()
		viper.Set("ghq", kcgExec.IsCommandAvailable("ghq"))
		viper.Set("repos", config.Repos)
		path, _ := cmd.Flags().GetString("path")
		WriteConfig(path)
		fmt.Println("Create config file at: " + viper.ConfigFileUsed())
	},
}

var configureSetCmd = &cobra.Command{
	Use:   "set <name>",
	Short: "Add repository config",
	Long:  `Add repository config`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initRepo()
		if _, ok := config.Repos[args[0]]; !ok {
			config.Repos[args[0]] = &kcg.RepositoryConfig{}
		}
		if repo, _ := cmd.Flags().GetString("repo"); repo != "" {
			config.Repos[args[0]].Repo = repo

			if config.Ghq {
				if path, err := kcgExec.GhqPath(repo); err == nil {
					config.Repos[args[0]].Path = path
				}
			}
		}
		if path, _ := cmd.Flags().GetString("path"); path != "" {
			config.Repos[args[0]].Path = path
		}
		if branchAlias, _ := cmd.Flags().GetStringArray("branch-alias"); len(branchAlias) != 0 {
			config.Repos[args[0]].Alias = branchAlias
		}
		if group, _ := cmd.Flags().GetStringArray("group"); len(group) != 0 {
			config.Repos[args[0]].Group = group
		}
		if setup, _ := cmd.Flags().GetStringArray("setup"); len(setup) != 0 {
			config.Repos[args[0]].Setup = setup
		}
		if update, _ := cmd.Flags().GetStringArray("update"); len(update) != 0 {
			config.Repos[args[0]].Update = update
		}
		viper.Set("repos", config.Repos)
		WriteConfig("")
	},
}

var configureImportCmd = &cobra.Command{
	Use:   "import",
	Short: "import specified config file into default file",
	Long:  `import specified config file into default file`,
	Run: func(cmd *cobra.Command, args []string) {
		importFilePath, _ := cmd.Flags().GetString("path")
		useGhq, _ := cmd.Flags().GetBool("ghq")

		if importFilePath == "" && !useGhq {
			cmd.PrintErrln(cmd.Help())
			return
		}

		initRepo()
		tempConfig := config

		if importFilePath != "" {
			if importConfig, err := importConfig(importFilePath); err == nil {
				for index, repo := range importConfig.Repos {
					tempConfig.Repos[index] = repo
				}
			} else {
				cmd.PrintErrln(err)
			}
		}

		if useGhq {
			for index, repo := range kcgExec.GhqList() {
				if _, ok := tempConfig.Repos[index]; !ok {
					tempConfig.Repos[index] = &kcg.RepositoryConfig{}
				}
				tempConfig.Repos[index].Repo = repo

				if path, err := kcgExec.GhqPath(repo); err == nil {
					tempConfig.Repos[index].Path = path
				}
			}
		}

		viper.Set("ghq", kcgExec.IsCommandAvailable("ghq"))
		viper.Set("repos", tempConfig.Repos)
		WriteConfig(cfgFile)
	},
}

var configureAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add repository config",
	Long:  `Add repository config`,
}

var configureAddAliasCmd = &cobra.Command{
	Use:   "branch-alias <name> <alias:value>",
	Short: "Add branch-alias config",
	Long:  `Add branch-alias config`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if _, ok := config.Repos[args[0]]; ok {
			config.Repos[args[0]].Alias = append(config.Repos[args[0]].Alias, args[1])
			viper.Set("repos", config.Repos)
			WriteConfig("")
		} else {
			fmt.Printf("    \x1b[31m%s\x1b[0m %s", "not exists", args[0])
		}
	},
}

var configureAddGroupCmd = &cobra.Command{
	Use:   "group <name> <group_name>",
	Short: "Add group config",
	Long:  `Add group config`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if _, ok := config.Repos[args[0]]; ok {
			config.Repos[args[0]].Group = append(config.Repos[args[0]].Group, args[1])
			viper.Set("repos", config.Repos)
			WriteConfig("")
		} else {
			fmt.Printf("    \x1b[31m%s\x1b[0m %s", "not exists", args[0])
		}
	},
}

var configureAddSetupCmd = &cobra.Command{
	Use:   "setup <name> <command>",
	Short: "Add setup command",
	Long:  `Add setup command`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if _, ok := config.Repos[args[0]]; ok {
			config.Repos[args[0]].Setup = append(config.Repos[args[0]].Setup, args[1])
			viper.Set("repos", config.Repos)
			WriteConfig("")
		} else {
			fmt.Printf("    \x1b[31m%s\x1b[0m %s", "not exists", args[0])
		}
	},
}

var configureAddUpdateCmd = &cobra.Command{
	Use:   "update <name> <command>",
	Short: "Add update command",
	Long:  `Add update command`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if _, ok := config.Repos[args[0]]; ok {
			config.Repos[args[0]].Update = append(config.Repos[args[0]].Update, args[1])
			viper.Set("repos", config.Repos)
			WriteConfig("")
		} else {
			fmt.Printf("    \x1b[31m%s\x1b[0m %s", "not exists", args[0])
		}
	},
}

var configureDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete repository config",
	Long:  `Delete repository config`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initRepo()
		delete(config.Repos, args[0])
		viper.Set("repos", config.Repos)
		WriteConfig("")
	},
}

func initRepo() {
	if config.Repos == nil {
		config.Repos = map[string]*kcg.RepositoryConfig{}
	}
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

func importConfig(path string) (kcg.Config, error) {
	var importConfig kcg.Config

	if !kcgExec.FileExists(path) {
		return importConfig, fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "not exists", path)
	}

	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return importConfig, fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "invalid", "cant read config file")
	}

	if err := viper.Unmarshal(&importConfig); err != nil {
		return importConfig, fmt.Errorf("    \x1b[31m%s\x1b[0m %s", "invalid", "cant unmarshal config")
	}

	return importConfig, nil
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.AddCommand(configureInitCmd)
	configureInitCmd.Flags().String("path", "", "write config file path")

	configureCmd.AddCommand(configureSetCmd)
	configureSetCmd.Flags().String("repo", "", "remote repository")
	configureSetCmd.Flags().String("path", "", "local dir")
	configureSetCmd.Flags().StringArray("branch-alias", []string{}, "specify like \"NAME:VALUE\"")
	configureSetCmd.Flags().StringArray("group", []string{}, "group")
	configureSetCmd.Flags().StringArray("setup", []string{}, "setup command")
	configureSetCmd.Flags().StringArray("update", []string{}, "update command")

	configureCmd.AddCommand(configureImportCmd)
	configureImportCmd.Flags().Bool("ghq", false, "Import from 'ghq list'")
	configureImportCmd.Flags().String("path", "", "configure file path")

	configureCmd.AddCommand(configureAddCmd)
	configureAddCmd.AddCommand(configureAddGroupCmd)
	configureAddCmd.AddCommand(configureAddAliasCmd)
	configureAddCmd.AddCommand(configureAddSetupCmd)
	configureAddCmd.AddCommand(configureAddUpdateCmd)

	configureCmd.AddCommand(configureDeleteCmd)
}
