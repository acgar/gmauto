package gmail

import (
	"log"

	"google.golang.org/api/gmail/v1"
)

func GetProfileEmail(srv *gmail.Service) string {
	profile, err := srv.Users.GetProfile("me").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	return profile.EmailAddress
}

func GetLabelsIds(srv *gmail.Service) map[string]string {
	labels := make(map[string]string)
	r, err := srv.Users.Labels.List("me").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}

	for _, l := range r.Labels {
		labels[l.Name] = l.Id
	}

	return labels
}

func AddLabelsToMessages(srv *gmail.Service, messagesIds []string, labelIds []string) {
	req := &gmail.BatchModifyMessagesRequest{
		Ids:         messagesIds,
		AddLabelIds: labelIds,
	}

	err := srv.Users.Messages.BatchModify("me", req).Do()
	if err != nil {
		log.Fatalf("Unable to add labels to messages %v", err)
	}
}

func RemoveLabelsFromMessages(srv *gmail.Service, messagesIds []string, labelIds []string) {
	req := &gmail.BatchModifyMessagesRequest{
		Ids:            messagesIds,
		RemoveLabelIds: labelIds,
	}

	err := srv.Users.Messages.BatchModify("me", req).Do()
	if err != nil {
		log.Fatalf("Unable to remove labels from messages %v", err)
	}
}

func GetMessagesWithLabel(srv *gmail.Service, labelId string) []*gmail.Message {
	r, err := srv.Users.Messages.List("me").LabelIds(labelId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve message from label: %v", err)
	}

	return r.Messages
}

// Todo: mejor older_than:30d ?
func GetMessagesWithLabelsAndQuery(srv *gmail.Service, labelIds []string, query string) []*gmail.Message {
	r, err := srv.Users.Messages.List("me").LabelIds(labelIds...).Q(query).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve messages: %v", err)
	}

	return r.Messages
}

func GetMessageMetadataHeaders(srv *gmail.Service, messageId string) []*gmail.MessagePartHeader {
	r, err := srv.Users.Messages.Get("me", messageId).Format("metadata").MetadataHeaders("subject", "from").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve message: %v", err)
	}

	return r.Payload.Headers
}
