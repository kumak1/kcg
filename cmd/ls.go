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
	"os"
	"strings"
	"text/tabwriter"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show repository list.",
	Run: func(cmd *cobra.Command, args []string) {
		repoFlag, _ := cmd.Flags().GetString("repo")
		groupFlag, _ := cmd.Flags().GetString("group")
		gitCommand := kcg.GitCommand(config)

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 1, '\t', 0)
		fmt.Fprintln(w, "NAME\tGROUPS\tREMOTE REPO\tLOCAL PATH")

		for index, repo := range gitCommand.List(repoFlag, groupFlag, "") {
			path := gitCommand.Path(repo)
			groups := strings.Join(repo.Groups, ",")
			fmt.Fprintln(w, index+"\t"+groups+"\t"+repo.Repo+"\t"+path)
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().String("repo", "", "repository name")
	lsCmd.Flags().String("group", "", "repository group name")
}
