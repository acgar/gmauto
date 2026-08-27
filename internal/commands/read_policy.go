package commands

import (
	"context"
	"fmt"
	"gmauto/internal/gmail"
	"gmauto/internal/policies"
	"gmauto/internal/utils"

	"github.com/urfave/cli/v3"
	sdkGmail "google.golang.org/api/gmail/v1"
)

func ExecReadPolicyCommand(ctx context.Context, cmd *cli.Command) error {
	srv := gmail.GetGmailService(ctx)
	labels := gmail.GetLabelsIds(srv)
	readPolicies := policies.LoadReadPolicies()

	for _, policy := range readPolicies {
		fmt.Println("Read policy for tag ", policy.TagName)
		executeReadPolicy(srv, labels[policy.TagName], policy.ExpirationInDays)
	}
	return nil
}

func executeReadPolicy(srv *sdkGmail.Service, labelId string, days int) {
	messages := gmail.GetMessagesWithLabelsAndQuery(srv, []string{labelId, gmail.UnreadLabelId}, "before:"+utils.DaysAgo(days))

	messagesIds := make([]string, 0, len(messages))
	for _, message := range messages {
		messagesIds = append(messagesIds, message.Id)
	}

	// Mark as read
	if len(messages) > 0 {
		gmail.RemoveLabelsFromMessages(srv, messagesIds, []string{gmail.UnreadLabelId})
		fmt.Printf("  - %d messages mark as read\n", len(messages))
	} else {
		fmt.Println("  - no messages found")
	}

}
