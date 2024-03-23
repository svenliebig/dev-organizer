/*
Copyright © 2024 Sven Liebig
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	_ "github.com/svenliebig/work-environment/pkg/cd/clients"
	_ "github.com/svenliebig/work-environment/pkg/ci/bamboo"
	"github.com/svenliebig/work-environment/pkg/we"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "we",
		Short: "the work environment cli help you to organize and maintain a productive work environment",
		Long: `the work environment cli help you to organize and maintain a productive work environment
by providing a link between the tool you use the most (the cli) and the applications
that you have to use (like CI, CD, Dev/QA Stages, etc).

The goal is, to get less disrupted in your workflow and more time for important 
things, rather than waiting for lists of ci plans to fetch or waiting for bloated
web applications (*cough cough* jira *cough*) to execute a base functionality.`,
	}
	rootInfoCmd = &cobra.Command{
		Use:   "info",
		Short: "print the contextual information of the current working directory",
		Long: `print the contextual information of the current working directory
like configured repository providers, ci/cd providers, etc.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := we.Info()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(rootInfoCmd)
}
