package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Project utilities",
	Long:  `erp developer project utilities.`,
}

var projectInitCmd = &cobra.Command{
	Use:   "init PROJECT_PATH",
	Short: "Initialize a project",
	Long: `Initialize a new project in the current directory with the given path (e.g. github.com/myuser/my-erp-project). 
This will create:
- A go.mod file
- A erp.toml file
All parameters passed as command line arguments or env variables will be set in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create the go.mod file
		if len(args) == 0 {
			fmt.Println("You must specify a project path.")
			os.Exit(1)
		}
		projectPath := args[0]
		if err := runCommand("go", "mod", "init", projectPath); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Create the erp.toml file
		projectDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err = writeConfigFile(projectDir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var projectCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean the project directory",
	Long: `Clean the current directory from all generated and test artifacts.
You should use this command before committing your work.`,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("go", "mod", "edit", "-dropreplace", "github.com/Pedro-lmso-erp/pool@v1.0.2")
		if err := removeProjectDir(PoolDirRel); err != nil {
			fmt.Println(err)
		}
		if err := removeProjectDir(ResDirRel); err != nil {
			fmt.Println(err)
		}
		runCommand("go", "mod", "tidy")
	},
}

func removeProjectDir(dir string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	poolDir := filepath.Join(cwd, dir)
	return os.RemoveAll(poolDir)
}

func writeConfigFile(projectDir string) error {
	InitConfig()
	cfgFile := filepath.Join(projectDir, "erp.toml")
	return viper.WriteConfigAs(cfgFile)
}

func init() {
	erpCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectInitCmd)
	projectCmd.AddCommand(projectCleanCmd)
}
