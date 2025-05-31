package cmd

import (
	"github.com/spf13/cobra"
	"new-blog/cmd/gen"
	"os"
)

var (
	dir, _ = os.Getwd()
)

// gen model name
// gen module name
// gen repository name
// gen service name

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate code for models, modules, repositories, and services",
	Long:  `Generate code for models, modules, repositories, and services`,
}

var genModelCmd = &cobra.Command{
	Use:   "model",
	Short: "Generate model code",
	Run: func(cmd *cobra.Command, args []string) {
		generateModel(args)
	},
}

var genModuleCmd = &cobra.Command{
	Use:   "module",
	Short: "Generate module code",
	Run: func(cmd *cobra.Command, args []string) {
		generateModule(args)
	},
}

var genServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Generate service code",
	Run: func(cmd *cobra.Command, args []string) {
		generateService(args)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.AddCommand(
		genModelCmd,
		genModuleCmd,
		genServiceCmd,
	)
}

func generateModel(args []string) {
	err := gen.GenerateModel(args)
	if err != nil {
		return
	}
}

func generateModule(args []string) {
	err := gen.GenerateModule(args)
	if err != nil {
		return
	}
}

func generateService(args []string) {
	err := gen.GenerateService(args)
	if err != nil {
		return
	}
}
