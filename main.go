package main

import (
	"context"
	"gmauto/internal/commands"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "login",
				Usage: "Login commands",
				Commands: []*cli.Command{
					{
						Name:    "check",
						Aliases: []string{"c"},
						Usage:   "Show current login status",
						Action:  commands.CheckLoginCommand,
					},
				},
			},
			{
				Name:  "show",
				Usage: "Displays information",
				Commands: []*cli.Command{
					{
						Name:    "labels",
						Aliases: []string{"l"},
						Usage:   "Show current available gmail labels",
						Action:  commands.ShowLabelsCommand,
					},
					{
						Name:    "tag-rules",
						Aliases: []string{"r"},
						Usage:   "Show current tag rules",
						Action:  commands.ShowRulesCommand,
					},
					{
						Name:    "policies",
						Aliases: []string{"p"},
						Usage:   "Show current configured policies",
						Action:  commands.ShowPoliciesCommand,
					},
				},
			},
			{
				Name:   "tag",
				Usage:  "Tag mails with user defined rules",
				Action: commands.TagCommand,
			},
			{
				Name:  "exec",
				Usage: "Executes policies",
				Commands: []*cli.Command{
					{
						Name:    "read-policy",
						Aliases: []string{"read"},
						Usage:   "Executes read policy. Mark messages as read.",
						Action:  commands.ExecReadPolicyCommand,
					},
					{
						Name:    "trash-policy",
						Aliases: []string{"del"},
						Usage:   "Executes trash policy. Deletes messages.",
						Action:  commands.ExecTrashPolicyCommand,
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
