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

	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "A brief description of your command",
	Long:  ``,
}

var configureInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		initRepo()
		viper.Set("ghq", kcgExec.IsCommandAvailable("ghq"))
		viper.Set("repos", config.Repos)
		path, _ := cmd.Flags().GetString("path")
		kcg.WriteConfig(path)
		fmt.Println("Create config file at: " + viper.ConfigFileUsed())
	},
}

var configureSetCmd = &cobra.Command{
	Use:   "set <name>",
	Short: "A brief description of your command",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initRepo()
		if _, ok := config.Repos[args[0]]; !ok {
			config.Repos[args[0]] = &kcg.RepositoryConfig{}
		}
		if repo, _ := cmd.Flags().GetString("repo"); repo != "" {
			config.Repos[args[0]].Repo = repo
		}
		if path, _ := cmd.Flags().GetString("path"); path != "" {
			config.Repos[args[0]].Path = path
		}
		if groups, _ := cmd.Flags().GetStringArray("groups"); len(groups) != 0 {
			config.Repos[args[0]].Groups = groups
		}
		if setup, _ := cmd.Flags().GetStringArray("setup"); len(setup) != 0 {
			config.Repos[args[0]].Setup = setup
		}
		viper.Set("repos", config.Repos)
		kcg.WriteConfig("")
	},
}

var configureDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "A brief description of your command",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initRepo()
		delete(config.Repos, args[0])
		viper.Set("repos", config.Repos)
		kcg.WriteConfig("")
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.AddCommand(configureInitCmd)
	configureInitCmd.Flags().String("path", "", "write config file path")

	configureCmd.AddCommand(configureSetCmd)
	configureSetCmd.Flags().String("repo", "", "remote repository (required)")
	configureSetCmd.Flags().String("path", "", "local dir")
	configureSetCmd.Flags().StringArray("groups", []string{}, "group")
	configureSetCmd.MarkFlagRequired("repo")

	configureCmd.AddCommand(configureDeleteCmd)
}

func initRepo() {
	if config.Repos == nil {
		config.Repos = map[string]*kcg.RepositoryConfig{}
	}
}

//kcg config init
//kcg config add NAME --repo --path --groups --setup
//kcg config set NAME --repo --path --groups --setup
//kcg config delete NAME