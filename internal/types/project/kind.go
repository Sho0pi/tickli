package project

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

type Kind string

const (
	KindTask    Kind = "TASK"
	KindNote    Kind = "NOTE"
	KindInbox   Kind = "INBOX"
	KindUnknown Kind = "UNKNOWN"
)

var KindCompletion = []cobra.Completion{
	cobra.CompletionWithDesc(string(KindTask), "Task project"),
	cobra.CompletionWithDesc(string(KindNote), "Note project"),
}

var KindCompletionFunc = cobra.FixedCompletions(KindCompletion, cobra.ShellCompDirectiveNoFileComp)

func (k *Kind) UnmarshalJSON(data []byte) error {
	var kind string
	if err := json.Unmarshal(data, &kind); err != nil {
		return err
	}
	switch kind {
	case string(KindTask), string(KindNote), string(KindInbox), string(KindUnknown):
		*k = Kind(kind)
	default:
		*k = KindUnknown
	}
	return nil
}

func (k Kind) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(k))
}

func (k Kind) String() string {
	switch k {
	case KindTask:
		return "📝Task"
	case KindNote:
		return "📖Note"
	case KindInbox:
		return "📥Inbox"
	default:
		return "🔧Unknown"
	}
}

func (k *Kind) Set(s string) error {
	switch s {
	case string(KindTask), string(KindNote), string(KindInbox):
		*k = Kind(s)
	default:
		return fmt.Errorf("invalid project kind %q", s)
	}
	return nil
}

func (k *Kind) Type() string {
	return "ProjectKind"
}
