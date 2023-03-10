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
	"github.com/kumak1/kcg/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "run `git clone` each repository",
	Long:  `Running git clone command each repository`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")

		pkg.ListParallelFor(func(key string, repoConf *pkg.RepositoryConfig) {
			if output, err := pkg.Clone(repoConf); err == nil {
				cmd.Printf(pkg.ValidMessage("✔", key))
				if output != "" {
					cmd.Println(output)
				}
			} else {
				cmd.Print(pkg.ErrorMessage("X", key))
				cmd.Print(output + err.Error())
			}
		}, groupFlag, filterFlag)

		if config.Ghq {
			for index, repo := range pkg.List("", "") {
				if path, exists := pkg.Path(repo); path != "" && exists {
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
