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
	"github.com/spf13/viper"
	"sync"

	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "run `git clone` each repository",
	Long:  `Running git clone command each repository`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")

		var wg sync.WaitGroup

		for index, repo := range kcg.List(groupFlag, filterFlag) {
			wg.Add(1)
			index := index
			repo := repo
			go func() {
				output, err := kcg.Clone(repo)
				if err == nil {
					cmd.Printf(kcg.ValidMessage("✔", index))
					if output != "" {
						cmd.Println(output)
					}
				} else {
					cmd.Print(kcg.ErrorMessage("X", index))
					cmd.Print(output + err.Error())
				}
				wg.Done()
			}()
		}

		wg.Wait()

		if config.Ghq {
			for index, repo := range kcg.List("", "") {
				if path, exists := kcg.Path(repo); path != "" && exists {
					config.Repos[index].Path = path
				}
			}
			viper.Set("repos", config.Repos)

			if err := WriteConfig(""); err != nil {
				cmd.PrintErrln("The config file could not write")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	assignSearchFlags(cloneCmd)
}
