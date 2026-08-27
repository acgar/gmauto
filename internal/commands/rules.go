package commands

import (
	"context"
	"gmauto/internal/rules"

	"github.com/urfave/cli/v3"
)

func ShowRulesCommand(ctx context.Context, cmd *cli.Command) error {
	rules.Display(rules.LoadRules())
	return nil
}
