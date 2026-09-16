package commands

import (
	"context"
	"fmt"
	"gmauto/internal/gmail"
	"gmauto/internal/policies"
	"strconv"

	"github.com/urfave/cli/v3"
	sdkGmail "google.golang.org/api/gmail/v1"
)

func ExecTrashPolicyCommand(ctx context.Context, cmd *cli.Command) error {
	srv := gmail.GetGmailService(ctx)
	labels := gmail.GetLabelsIds(srv)
	trashPolicies := policies.LoadTrashPolicies()

	for _, policy := range trashPolicies {
		fmt.Println("Trash policy for tag ", policy.TagName)
		executeTrashPolicy(srv, labels[policy.TagName], policy.ExpirationInDays)
	}
	return nil
}

func executeTrashPolicy(srv *sdkGmail.Service, labelId string, days int) {
	query := "older_than:" + strconv.Itoa(days) + "d"
	messages := gmail.GetMessagesWithLabelsAndQuery(srv, []string{labelId}, query)

	messagesIds := make([]string, 0, len(messages))
	for _, message := range messages {
		messagesIds = append(messagesIds, message.Id)
	}

	// Remove messages
	if len(messages) > 0 {
		gmail.AddLabelsToMessages(srv, messagesIds, []string{gmail.TrashLabelId})
		fmt.Printf("  - %d messages sent to trash\n", len(messages))
	} else {
		fmt.Println("  - no messages found")
	}

}
