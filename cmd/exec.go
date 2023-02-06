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
	"github.com/kumak1/kcg/kcg"
	"github.com/spf13/cobra"
	"os"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run commands on each repository",
	Long:  `Run commands on each repository`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Print(cmd.Help())
			return
		}

		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")

		execCommandName := args[0]
		resultOutput := ""
		resultError := false
		validCommand := false

		kcg.ListParallelFor(func(key string, repoConf *kcg.RepositoryConfig) {
			for execKey, execStrings := range repoConf.Exec {
				if execKey == execCommandName {
					validCommand = true
					for _, cmdString := range execStrings {
						expandEnvCommand := os.ExpandEnv(cmdString)
						output, err := kcg.Run(repoConf, expandEnvCommand)

						if err == nil {
							resultOutput += kcg.ValidMessage("run", cmdString)
							if output != "" {
								resultOutput += output + "\n"
							}
						} else {
							resultOutput += kcg.ErrorMessage("run", cmdString).Error()
							if output != "" {
								resultOutput += output + "\n"
							}
							resultOutput += err.Error() + "\n"
							resultError = true
							break
						}
					}

					if resultError {
						cmd.Print(kcg.ErrorMessage("X", key))
					} else {
						cmd.Printf(kcg.ValidMessage("✔", key))
					}

					if resultOutput != "" {
						cmd.Print(resultOutput)
					}
				}
			}
		}, groupFlag, filterFlag)

		if !validCommand {
			cmd.Print(cmd.Help())
		}
	},
}

func execList(group string, filter string) map[string][]string {
	execList := map[string][]string{}
	for key, repoConf := range kcg.List(group, filter) {
		for execKey := range repoConf.Exec {
			execList[execKey] = append(execList[execKey], key)
		}
	}
	return execList
}

var execListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show commands list on each repository",
	Long:  `Show commands list on each repository`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")
		quietFlag, _ := cmd.Flags().GetBool("quiet")

		execList := execList(groupFlag, filterFlag)
		for listKey, listValue := range execList {
			if quietFlag {
				cmd.Println(listKey)
			} else {
				cmd.Println(listKey + ":")
				for _, repoKey := range listValue {
					cmd.Println("  " + repoKey)
				}
			}
		}
	},
}

var execSetCmd = &cobra.Command{
	Use:   "set <name>",
	Short: "Set command on specify repository",
	Long:  `Set command on specify repository`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoName := args[0]
		if _, ok := config.Repos[repoName]; !ok {
			cmd.PrintErrln(kcg.ErrorMessage("invalid name", repoName))
			cmd.PrintErrln(cmd.Help())
			os.Exit(1)
		}

		commandName, _ := cmd.Flags().GetString("name")
		commandStrings, _ := cmd.Flags().GetStringArray("command")

		if err := cmd.MarkFlagRequired("name"); commandName == "" || err != nil {
			cmd.PrintErr(kcg.ErrorMessage("invalid", "flag needs an argument: --name"))
			return
		}
		if err := cmd.MarkFlagRequired("command"); len(commandStrings) == 0 || commandStrings[0] == "" || err != nil {
			cmd.PrintErr(kcg.ErrorMessage("invalid", "flag needs an argument: --command"))
			return
		}

		if config.Repos[repoName].Exec == nil {
			config.Repos[repoName].Exec = map[string][]string{}
		}
		config.Repos[repoName].Exec[commandName] = commandStrings

		if err := UpdateConfig(); err != nil {
			cmd.PrintErrln("The config file could not write")
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	assignSearchFlags(execCmd)

	execCmd.AddCommand(execListCmd)
	assignSearchFlags(execListCmd)
	execListCmd.Flags().BoolP("quiet", "q", false, "Only display command")

	execCmd.AddCommand(execSetCmd)
	execSetCmd.Flags().StringP("name", "n", "", "Execute command name (required)")
	execSetCmd.Flags().StringArrayP("command", "c", []string{}, "Execute command string (required)")
}
