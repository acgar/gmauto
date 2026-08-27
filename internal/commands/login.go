package commands

import (
	"context"
	"fmt"
	"gmauto/internal/gmail"
	"strings"

	"github.com/urfave/cli/v3"
)

func CheckLoginCommand(ctx context.Context, cmd *cli.Command) error {
	srv := gmail.GetGmailService(ctx)
	email := gmail.GetProfileEmail(srv)
	if len(strings.TrimSpace(email)) > 0 {
		fmt.Println("Login OK with '" + email + "' account.")
	} else {
		fmt.Println("No account logged.")
	}

	return nil
}
