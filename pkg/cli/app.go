package cli

import (
	"s3lambda/pkg/cli/command"

	"github.com/urfave/cli/v2"
)

func NewApplication() *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			command.RunS3ListenerLambdaCommand,
		},
	}
}
