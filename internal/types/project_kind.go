package types

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

type ProjectKind string

const (
	KindTask    ProjectKind = "TASK"
	KindNote    ProjectKind = "NOTE"
	KindInbox   ProjectKind = "INBOX"
	KindUnknown ProjectKind = "UNKNOWN"
)

func RegisterProjectKindCompletions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{string(KindTask), string(KindNote)}, cobra.ShellCompDirectiveDefault
}

func (k *ProjectKind) UnmarshalJSON(data []byte) error {
	var kind string
	if err := json.Unmarshal(data, &kind); err != nil {
		return err
	}
	switch kind {
	case string(KindTask), string(KindNote), string(KindInbox), string(KindUnknown):
		*k = ProjectKind(kind)
	default:
		*k = KindUnknown
	}
	return nil
}

func (k ProjectKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(k))
}

func (k ProjectKind) String() string {
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

func (k *ProjectKind) Set(s string) error {
	switch s {
	case string(KindTask), string(KindNote), string(KindInbox):
		*k = ProjectKind(s)
	default:
		return fmt.Errorf("invalid project kind %q", s)
	}
	return nil
}

func (k *ProjectKind) Type() string {
	return "ProjectKind"
}
