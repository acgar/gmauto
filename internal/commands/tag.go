package commands

import (
	"context"
	"fmt"
	"gmauto/internal/gmail"
	"gmauto/internal/rules"
	"regexp"
	"strings"

	"github.com/urfave/cli/v3"
	sdkGmail "google.golang.org/api/gmail/v1"
)

func TagCommand(ctx context.Context, cmd *cli.Command) error {
	srv := gmail.GetGmailService(ctx)
	labels := gmail.GetLabelsIds(srv)

	inboxLabelId := labels[gmail.InboxTag]
	mailsToProcess := gmail.GetMessagesWithLabel(srv, inboxLabelId)

	if len(mailsToProcess) > 0 {
		mailsToProcessIds := make([]string, 0, len(mailsToProcess))
		for _, m := range mailsToProcess {
			mailsToProcessIds = append(mailsToProcessIds, m.Id)
		}

		// Process Rules
		mailsToTag := processTagRulesOnMails(srv, mailsToProcessIds)

		// Tag mails
		for tag, mailsId := range mailsToTag {
			labelId := labels[tag]
			gmail.AddLabelsToMessages(srv, mailsId, []string{labelId})
		}
		// Mark mails as processed
		gmail.RemoveLabelsFromMessages(srv, mailsToProcessIds, []string{inboxLabelId})
	} else {
		fmt.Println("No inbox mails to process")
	}
	return nil
}

func processTagRulesOnMails(srv *sdkGmail.Service, mailIds []string) map[string][]string {
	tagRulesArray := rules.LoadRules()
	mailsToTag := make(map[string][]string)
	fmt.Printf("%d messages to process:\n", len(mailIds))
	for _, mId := range mailIds {
		// for each message, check if satisfy tag rules.
		fmt.Printf("  - %s\n", mId)
		// fetch message metadata
		headers := gmail.GetMessageMetadataHeaders(srv, mId)
		// Extract message metadata
		from, subject := "", ""
		for _, header := range headers {
			fmt.Printf("    * %s: %s\n", header.Name, header.Value)
			if header.Name == "From" {
				from = header.Value
				// Clean from field:
				// Some-Company <noreply@communication.some-company.com> -> noreply@communication.some-company.com
				re := regexp.MustCompile(`<([^>]+)>`)
				matches := re.FindStringSubmatch(from)
				if len(matches) > 1 {
					from = matches[1]
				}
			} else if header.Name == "Subject" {
				subject = header.Value
			}
		}
		// Check any tags to add to this mail
		for _, tagRules := range tagRulesArray {
			if rules.CheckTagRules(from, subject, tagRules) {
				mailsToTag[tagRules.TagName] = append(mailsToTag[tagRules.TagName], mId)
			}
		}
	}

	fmt.Printf("Found %d tags to apply to mails\n", len(mailsToTag))
	for tag, mails := range mailsToTag {
		fmt.Printf("- %s (%s)\n", tag, strings.Join(mails, ","))
	}
	fmt.Println("------------------------------------------")
	return mailsToTag
}
