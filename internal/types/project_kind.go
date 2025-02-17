package types

import "encoding/json"

type ProjectKind string

const (
	KindTask    ProjectKind = "TASK"
	KindNote    ProjectKind = "NOTE"
	KindInbox   ProjectKind = "INBOX"
	KindUnknown ProjectKind = "UNKNOWN"
)

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
