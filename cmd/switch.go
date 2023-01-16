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
	"github.com/spf13/cobra"
	"strings"
	"sync"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch <branch>",
	Short: "run `git switch` on each repository dir",
	Long:  `Running git switch command on each repository dir`,
	//Args: cobra.ExactArgs(1),
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return fmt.Errorf("missing branch argument\n")
		}
		return nil
	},
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
				output, err := kcgCmd.Switch(repo, args[0])
				if err == nil {
					cmd.Printf(validMessageFormat, "✔", index)
					message := string(output)
					if !strings.Contains(message, "Already on") {
						cmd.Print(message)
					}
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
	rootCmd.AddCommand(switchCmd)
	assignSearchFlags(switchCmd)
}
