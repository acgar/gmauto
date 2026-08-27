package commands

import (
	"context"
	"fmt"
	"gmauto/internal/policies"
	"strconv"

	"github.com/urfave/cli/v3"
)

func ShowPoliciesCommand(ctx context.Context, cmd *cli.Command) error {
	readPolicies := policies.LoadReadPolicies()
	trashPolicies := policies.LoadTrashPolicies()

	fmt.Println("✉️ Read policies:")
	for _, policy := range readPolicies {
		fmt.Println("  - " + policy.TagName + ": " + strconv.Itoa(policy.ExpirationInDays) + " days")
	}

	fmt.Println("🗑️ Trash policies:")
	for _, policy := range trashPolicies {
		fmt.Println("  - " + policy.TagName + ": " + strconv.Itoa(policy.ExpirationInDays) + " days")
	}

	return nil
}
