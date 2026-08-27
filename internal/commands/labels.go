package commands

import (
	"context"
	"fmt"

	"gmauto/internal/gmail"

	"github.com/urfave/cli/v3"
)

func ShowLabelsCommand(ctx context.Context, cmd *cli.Command) error {
	srv := gmail.GetGmailService(ctx)
	labels := gmail.GetLabelsIds(srv)

	if len(labels) == 0 {
		fmt.Println("No labels found")
	} else {
		fmt.Println("Labels:")
		for name, id := range labels {
			fmt.Printf("  - %s (%s)\n", id, name)
		}
	}
	return nil
}
