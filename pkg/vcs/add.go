package vcs

import (
	"fmt"

	"github.com/svenliebig/work-environment/pkg/context"
	"github.com/svenliebig/work-environment/pkg/utils/cli"
)

func Add(ctx context.ProjectContext) error {
	// TODO this needs to be fixed
	c := ctx.Configuration()

	options := make([]string, 0, len(c.VCSEnvironments))
	for _, v := range c.VCSEnvironments {
		options = append(options, v.Identifier)
	}
	env := cli.Select("Select the VCS environment to add", options)

	vcse, err := c.GetVCSEnvironmentById(env)

	if err != nil {
		return err
	}

	err = ctx.UpdateVCS(vcse, "")

	if err != nil {
		return err
	}

	client, err := UseClient(ctx, vcse)

	if err != nil {
		return err
	}

	configuration, err := client.Configure()

	if err != nil {
		return err
	}

	err = ctx.UpdateVCS(vcse, configuration)

	if err != nil {
		return err
	}

	err = ctx.Close()

	if err != nil {
		return err
	}

	fmt.Printf(
		"%s the vsce '%s' to the project '%s'.\n",
		cli.Colorize(cli.Green, "Added"),
		cli.Colorize(cli.Purple, vcse.Identifier),
		cli.Colorize(cli.Purple, ctx.Project().Identifier),
	)

	return nil
}
