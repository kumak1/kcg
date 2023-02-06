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
}

// execSetupCmd represents the setup command
var execSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Run setup commands on each repository",
	Long:  `Running setup commands on each repository`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")

		kcg.ListParallelFor(func(key string, repoConf *kcg.RepositoryConfig) {
			resultOutput := ""
			resultError := false

			for _, command := range repoConf.Setup {
				expandEnvCommand := os.ExpandEnv(command)
				output, err := kcg.Run(repoConf, expandEnvCommand)
				resultOutput += "  "
				if err == nil {
					resultOutput += kcg.ValidMessage("run", command)
					if output != "" {
						resultOutput += output + "\n"
					}
				} else {
					resultOutput += kcg.ErrorMessage("run", command).Error()
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
		}, groupFlag, filterFlag)
	},
}

var execUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Run update commands on each repository",
	Long:  `Running update commands on each repository`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")

		kcg.ListParallelFor(func(key string, repoConf *kcg.RepositoryConfig) {
			resultOutput := ""
			resultError := false
			for _, command := range repoConf.Update {
				expandEnvCommand := os.ExpandEnv(command)
				output, err := kcg.Run(repoConf, expandEnvCommand)
				resultOutput += "  "
				if err == nil {
					resultOutput += kcg.ValidMessage("run", command)
					if output != "" {
						resultOutput += output + "\n"
					}
				} else {
					resultOutput += kcg.ErrorMessage("run", command).Error()
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
		}, groupFlag, filterFlag)
	},
}

var execListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show commands list on each repository",
	Long:  `Show commands list on each repository`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")
		quietFlag, _ := cmd.Flags().GetBool("quiet")

		execList := map[string][]string{}
		for key, repoConf := range kcg.List(groupFlag, filterFlag) {
			for execKey := range repoConf.Exec {
				execList[execKey] = append(execList[execKey], key)
			}
		}

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

func init() {
	rootCmd.AddCommand(execCmd)

	execCmd.AddCommand(execSetupCmd)
	assignSearchFlags(execSetupCmd)

	execCmd.AddCommand(execUpdateCmd)
	assignSearchFlags(execUpdateCmd)

	execCmd.AddCommand(execListCmd)
	assignSearchFlags(execListCmd)
	execListCmd.Flags().BoolP("quiet", "q", false, "Only display command")
}
