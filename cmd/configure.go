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
	"bytes"
	"fmt"
	kcgExec "github.com/kumak1/kcg/internal"
	"github.com/kumak1/kcg/pkg"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Operate config file",
	Long:  `Operate config file`,
}

var configureInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty config file",
	Long:  `Create an empty config file`,
	Run: func(cmd *cobra.Command, args []string) {
		initRepo()
		viper.Set("ghq", pkg.GhqValid())
		viper.Set("repos", config.Repos)
		path, _ := cmd.Flags().GetString("path")

		if err := WriteConfig(path); err == nil {
			cmd.Println("Create config file at: " + viper.ConfigFileUsed())
		} else {
			cmd.PrintErrln("The config file could not write")
		}
	},
}

var configureSetCmd = &cobra.Command{
	Use:   "set <name>",
	Short: "Add repository config",
	Long:  `Add repository config`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initRepo()
		if _, ok := config.Repos[args[0]]; !ok {
			config.Repos[args[0]] = &pkg.RepositoryConfig{}
			config.Repos[args[0]].Exec = map[string][]string{}
			config.Repos[args[0]].BranchAlias = map[string]string{}
		}
		if repo, _ := cmd.Flags().GetString("repo"); repo != "" {
			config.Repos[args[0]].Repo = repo

			if config.Ghq {
				if path, ok := pkg.Path(config.Repos[args[0]]); ok {
					config.Repos[args[0]].Path = path
				}
			}
		}
		if path, _ := cmd.Flags().GetString("path"); path != "" {
			config.Repos[args[0]].Path = path
		}
		if branchAlias, _ := cmd.Flags().GetStringArray("branch-alias"); len(branchAlias) != 0 {
			config.Repos[args[0]].SetAlias(branchAlias)
		}
		if group, _ := cmd.Flags().GetStringArray("group"); len(group) != 0 {
			config.Repos[args[0]].Group = group
		}
		if err := UpdateConfig(); err != nil {
			cmd.PrintErrln("The config file could not write")
		}
	},
}

var configureImportCmd = &cobra.Command{
	Use:   "import",
	Short: "import specified config file into default file",
	Long:  `import specified config file into default file`,
	Run: func(cmd *cobra.Command, args []string) {
		importPath, _ := cmd.Flags().GetString("path")
		importUrl, _ := cmd.Flags().GetString("url")
		useGhq, _ := cmd.Flags().GetBool("ghq")

		if importPath == "" && importUrl == "" && !useGhq {
			cmd.PrintErrln(cmd.Help())
			return
		}

		initRepo()
		tempConfig := config

		if importPath != "" {
			if importConfig, err := importConfigFile(importPath); err == nil {
				for index, repo := range importConfig.Repos {
					tempConfig.Repos[index] = repo
				}
			} else {
				cmd.PrintErrln(err)
			}
		}

		if importUrl != "" {
			if importConfig, err := importConfigUrl(importUrl); err == nil {
				for index, repo := range importConfig.Repos {
					tempConfig.Repos[index] = repo
				}
			} else {
				cmd.PrintErrln(err)
			}
		}

		if useGhq {
			if pkg.GhqValid() {
				tempConfig.Ghq = true
				pkg.Initialize(tempConfig)
				for index, repo := range pkg.GhqList() {
					if _, ok := tempConfig.Repos[index]; !ok {
						tempConfig.Repos[index] = &pkg.RepositoryConfig{}
					}
					tempConfig.Repos[index].Repo = repo

					if path, ok := pkg.Path(tempConfig.Repos[index]); ok {
						tempConfig.Repos[index].Path = path
					}
				}
			} else {
				cmd.PrintErr(pkg.ErrorMessage("invalid", "ghq command is not available"))
			}
		}

		viper.Set("ghq", pkg.GhqValid())
		viper.Set("repos", tempConfig.Repos)

		if err := WriteConfig(cfgFile); err != nil {
			cmd.PrintErrln("The config file could not write")
		}
	},
}

var configureExportCmd = &cobra.Command{
	Use:   "export",
	Short: "export config",
	Long:  `export config`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")

		viper.Reset()
		viper.Set("repos", pkg.List(groupFlag, filterFlag))

		if bs, err := yaml.Marshal(viper.AllSettings()); err == nil {
			fmt.Print(string(bs))
		} else {
			cmd.PrintErr(pkg.ErrorMessage("invalid", err.Error()))
		}
	},
}

var configureAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add repository config",
	Long:  `Add repository config`,
}

var configureAddAliasCmd = &cobra.Command{
	Use:   "branch-alias <name> <alias:value>",
	Short: "Add branch-alias config",
	Long:  `Add branch-alias config`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if _, ok := config.Repos[args[0]]; ok {
			config.Repos[args[0]].AddAlias(args[1])
			if err := UpdateConfig(); err != nil {
				cmd.PrintErrln("The config file could not write")
			}
		} else {
			cmd.PrintErr(pkg.ErrorMessage("not exists", args[0]))
		}
	},
}

var configureAddGroupCmd = &cobra.Command{
	Use:   "group <name> <group_name>",
	Short: "Add group config",
	Long:  `Add group config`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if _, ok := config.Repos[args[0]]; ok {
			config.Repos[args[0]].Group = append(config.Repos[args[0]].Group, args[1])
			if err := UpdateConfig(); err != nil {
				cmd.PrintErrln("The config file could not write")
			}
		} else {
			cmd.PrintErr(pkg.ErrorMessage("not exists", args[0]))
		}
	},
}

var configureDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete repository config",
	Long:  `Delete repository config`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initRepo()
		delete(config.Repos, args[0])
		if err := UpdateConfig(); err != nil {
			cmd.PrintErrln("The config file could not write")
		}
	},
}

func initRepo() {
	if config.Repos == nil {
		config.Repos = map[string]*pkg.RepositoryConfig{}
	}
}

func UpdateConfig() error {
	viper.Set("repos", config.Repos)
	if err := WriteConfig(""); err != nil {
		return fmt.Errorf("the config file could not write")
	}

	return nil
}

func WriteConfig(path string) error {
	if path != "" {
		viper.SetConfigFile(path)
	}

	if viper.ConfigFileUsed() == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.SetConfigFile(filepath.Join(home, ".kcg"))
	}

	viper.SetTypeByDefaultValue(false)
	return viper.WriteConfig()
}

func importConfigFile(path string) (pkg.Config, error) {
	var importConfig pkg.Config

	if !kcgExec.New().FileExists(path) {
		return importConfig, pkg.ErrorMessage("not exists", path)
	}

	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return importConfig, pkg.ErrorMessage("invalid", "cant read config file")
	}

	if err := viper.Unmarshal(&importConfig); err != nil {
		return importConfig, pkg.ErrorMessage("invalid", "cant unmarshal config")
	}

	return importConfig, nil
}

func importConfigUrl(url string) (pkg.Config, error) {
	var importConfig pkg.Config

	if res, err := http.Get(url); err == nil {
		if bodyBytes, err := io.ReadAll(res.Body); err == nil {
			if err := viper.ReadConfig(bytes.NewBuffer(bodyBytes)); err != nil {
				return importConfig, err
			}
		} else {
			return importConfig, err
		}
	} else {
		return importConfig, err
	}

	if err := viper.Unmarshal(&importConfig); err != nil {
		return importConfig, pkg.ErrorMessage("invalid", "cant unmarshal config")

	}

	return importConfig, nil
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.AddCommand(configureInitCmd)
	configureInitCmd.Flags().String("path", "", "write config file path")

	configureCmd.AddCommand(configureSetCmd)
	configureSetCmd.Flags().String("repo", "", "remote repository")
	configureSetCmd.Flags().String("path", "", "local dir")
	configureSetCmd.Flags().StringArray("branch-alias", []string{}, "specify like \"NAME:VALUE\"")
	configureSetCmd.Flags().StringArray("group", []string{}, "group")

	configureCmd.AddCommand(configureImportCmd)
	configureImportCmd.Flags().Bool("ghq", false, "Import from 'ghq list'")
	configureImportCmd.Flags().String("path", "", "configure file path")
	configureImportCmd.Flags().String("url", "", "configure file url")

	configureCmd.AddCommand(configureExportCmd)
	assignSearchFlags(configureExportCmd)

	configureCmd.AddCommand(configureAddCmd)
	configureAddCmd.AddCommand(configureAddGroupCmd)
	configureAddCmd.AddCommand(configureAddAliasCmd)

	configureCmd.AddCommand(configureDeleteCmd)
}
