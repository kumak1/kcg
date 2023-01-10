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
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "run setup commands on each repository",
	Long:  `Running setup commands on each repository`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")
		kcgCmd := kcg.Command(config)

		for index, repo := range kcgCmd.List(groupFlag, filterFlag) {
			fmt.Printf("  \x1b[32m%s\x1b[0m %s\n", "on", index)
			if len(repo.Setup) == 0 {
				fmt.Printf("    \x1b[33m%s\x1b[0m %s\n", "not exists", "setup command")
				continue
			}

			for _, command := range repo.Setup {
				fmt.Printf("    \x1b[32m%s\x1b[0m %s\n", "run", command)
				if err := kcgCmd.Run(repo, command); err != nil {
					fmt.Println(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	assignSearchFlags(setupCmd)
}
