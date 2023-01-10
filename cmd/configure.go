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

		if importFromGhq, _ := cmd.Flags().GetBool("import-from-ghq"); importFromGhq {
			for index, repo := range kcgExec.GhqList() {
				if _, ok := config.Repos[index]; !ok {
					config.Repos[index] = &kcg.RepositoryConfig{}
				}
				config.Repos[index].Repo = repo

				if path, err := kcgExec.GhqPath(repo); err == nil {
					config.Repos[index].Path = path
				}
			}
		}

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
		viper.Set("repos", config.Repos)
		WriteConfig("")
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

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.AddCommand(configureInitCmd)
	configureInitCmd.Flags().String("path", "", "write config file path")
	configureInitCmd.Flags().Bool("import-from-ghq", false, "create from `ghq list`")

	configureCmd.AddCommand(configureSetCmd)
	configureSetCmd.Flags().String("repo", "", "remote repository")
	configureSetCmd.Flags().String("path", "", "local dir")
	configureSetCmd.Flags().StringArray("branch-alias", []string{}, "specify like \"NAME:VALUE\"")
	configureSetCmd.Flags().StringArray("group", []string{}, "group")
	configureSetCmd.Flags().StringArray("setup", []string{}, "setup command")

	configureCmd.AddCommand(configureDeleteCmd)
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
