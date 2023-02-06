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
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "delete merged branch on each repository dir",
	Long:  `Delete merged branch on each repository dir`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")

		kcg.ListParallelFor(func(key string, repoConf *kcg.RepositoryConfig) {
			if output, err := kcg.Cleanup(repoConf); err == nil {
				cmd.Print(kcg.ValidMessage("✔", key))
				if output != "" {
					cmd.Println(output)
				}
			} else {
				cmd.Print(kcg.ErrorMessage("X", key))
				cmd.Print(output + err.Error())
			}
		}, groupFlag, filterFlag)
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
	assignSearchFlags(cleanupCmd)
}
