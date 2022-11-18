/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/svenliebig/work-environment/pkg/ci"
	"github.com/svenliebig/work-environment/pkg/utils"
)

// ciCmd represents the ci command
var (
	ciCmd = &cobra.Command{
		Use:   "ci",
		Short: "Configure and use continuous integrations",
		Long: `Configure and use continuous integrations, add specficis CI's to your project or
create a globally available work environment CI.`,
	}
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Adds a new CI to your work environment",
		Long:  `Adds a new CI to your work environment`,
		Run: func(cmd *cobra.Command, args []string) {
			u, err := cmd.Flags().GetString("url")

			if err != nil {
				log.Fatal(err)
			} else {
				ur, err := url.ParseRequestURI(u)

				if err != nil {
					log.Fatalf("the url parameter does not satisfy an url format\n%s", err)
				}

				u = (&url.URL{
					Scheme: ur.Scheme,
					Host:   ur.Host,
				}).String()
			}

			ciType, err := cmd.Flags().GetString("type")

			if err != nil {
				log.Fatal(err)
			}

			auth, err := cmd.Flags().GetString("auth")

			if err != nil {
				log.Fatal(err)
			}

			name, err := cmd.Flags().GetString("name")

			if err != nil {
				log.Fatal(err)
			}

			p, err := utils.GetPath([]string{})

			if err != nil {
				log.Fatal(err)
			}

			err = ci.Create(p, u, ciType, name, auth)

			if err != nil {
				log.Fatal(err)
			}
		},
	}
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds a CI to your project",
		Long: `Adds a CI to your project, you have to be inside the project path or specify the
project identifier. The CI identifier is required, when you have more than one CI specified in yur
work environment.`,
		Run: func(cmd *cobra.Command, args []string) {
			ciId, err := cmd.Flags().GetString("ciIdentifier")

			if err != nil {
				log.Fatal(fmt.Errorf("err while trying to get variable %q. %w", "ciIdentifier", err))
			}

			bambooKey, err := cmd.Flags().GetString("key")

			if err != nil {
				log.Fatal(err)
			}

			project, err := cmd.Flags().GetString("projectIdentifier")

			if err != nil {
				log.Fatal(err)
			}

			suggest, err := cmd.Flags().GetBool("suggest")

			if err != nil {
				log.Fatal(err)
			}

			p, err := utils.GetPath([]string{})

			if err != nil {
				log.Fatal(err)
			}

			err = ci.Add(p, ciId, project, bambooKey, suggest)

			if err != nil {
				log.Fatal(err)
			}
		},
	}
	openCmd = &cobra.Command{
		Use:   "open",
		Short: "Opens the CI environment in your default browser",
		Long:  `Opens the CI environment in your default browser`,
		Run: func(cmd *cobra.Command, args []string) {
			p, err := utils.GetPath([]string{})

			if err != nil {
				log.Fatal(err)
			}

			err = ci.Open(p)

			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(ciCmd)
	ciCmd.AddCommand(createCmd)
	ciCmd.AddCommand(addCmd)
	ciCmd.AddCommand(openCmd)

	createCmd.Flags().StringP("url", "u", "", "the URL of the CI you want to add\nexample: 'https://bamboo.company.com'")
	createCmd.Flags().StringP("type", "t", "", "the CI type, currently available types are 'bamboo'\nexample: 'bamboo'")
	createCmd.Flags().StringP("auth", "a", "", "your base64 auth token for the CI environment\nexample: '8fmiam7dm/2o3m8cunskeswefwe'")
	createCmd.Flags().StringP("name", "n", "", "the unique Identifier for the CI in your work environment\nexample: 'my-bamboo-ci'")

	createCmd.MarkFlagRequired("url")
	createCmd.MarkFlagRequired("type")
	createCmd.MarkFlagRequired("auth")
	createCmd.MarkFlagRequired("name")

	addCmd.Flags().StringP("ciIdentifier", "c", "", "the identifier of the ci\nexample: 'my-bamboo'")
	addCmd.Flags().StringP("projectIdentifier", "p", "", "the identifier of the project\nexample: 'my-project'")
	addCmd.Flags().BoolP("suggest", "s", false, "if set, you will get suggestions of bamboo project keys")
	addCmd.Flags().StringP("key", "b", "", "the key identifier for the project in the ci, not relevant if suggest is set\nexmaple: 'PRS-SZ'")

	// addCmd÷MarkFlagsMutuallyExclusive("suggest", "bambooKey")
}