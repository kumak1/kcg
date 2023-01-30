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
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")
		quietFlag, _ := cmd.Flags().GetBool("quiet")
		allFlag, _ := cmd.Flags().GetBool("all")

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 1, '\t', 0)

		if !quietFlag {
			if allFlag {
				fmt.Fprintln(w, "NAME\tCURRENT BRANCH\tGROUP\tBRANCH ALIAS\tREMOTE REPO\tLOCAL PATH")
			} else {
				fmt.Fprintln(w, "NAME\tCURRENT BRANCH\tGROUP")
			}
		}

		for index, repo := range kcg.List(groupFlag, filterFlag) {
			if quietFlag {
				fmt.Fprintln(w, index)
			} else {
				path, _ := kcg.Path(repo)
				branch := kcg.CurrentBranch(repo)
				group := strings.Join(repo.Group, ",")

				if allFlag {
					branchAlias := strings.Join(repo.Alias, ",")
					fmt.Fprintln(w, index+"\t"+branch+"\t"+group+"\t"+branchAlias+"\t"+repo.Repo+"\t"+path)
				} else {
					fmt.Fprintln(w, index+"\t"+branch+"\t"+group)
				}
			}
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolP("quiet", "q", false, "Only display repository name")
	lsCmd.Flags().BoolP("all", "a", false, "Display all repository setting")
	assignSearchFlags(lsCmd)
}
