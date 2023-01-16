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
	"sync"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "delete merged branch on each repository dir",
	Long:  `Delete merged branch on each repository dir`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")
		kcgCmd := kcg.Command(config)

		var wg sync.WaitGroup

		for index, repo := range kcgCmd.List(groupFlag, filterFlag) {
			wg.Add(1)
			index := index
			repo := repo
			go func() {
				output, err := kcgCmd.Cleanup(repo)
				if err == nil {
					cmd.Printf(validMessageFormat, "✔", index)
					cmd.Print(string(output))
				} else {
					cmd.Printf(invalidMessageFormat, "X", index)
					cmd.Print(string(output))
					cmd.Println(err)
				}
				wg.Done()
			}()
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
	assignSearchFlags(cleanupCmd)
}
