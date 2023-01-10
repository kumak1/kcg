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
		kcgCmd := kcg.Command(config)

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 1, '\t', 0)
		fmt.Fprintln(w, "NAME\tCURRENT BRANCH\tGROUP\tREMOTE REPO\tLOCAL PATH")

		for index, repo := range kcgCmd.List(groupFlag, filterFlag) {
			path, _ := kcgCmd.Path(repo)
			group := strings.Join(repo.Group, ",")
			branch := kcgCmd.CurrentBranch(repo)
			fmt.Fprintln(w, index+"\t"+branch+"\t"+group+"\t"+repo.Repo+"\t"+path)
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	assignSearchFlags(lsCmd)
}
