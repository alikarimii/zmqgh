package shared

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type MessageID string

func GenerateMessageID() MessageID {
	return MessageID(uuid.New().String())
}

func BuildMessageID(value string) (MessageID, error) {
	if value == "" {
		err := errors.New("empty input for MessageID")
		err = MarkAndWrapError(err, ErrInputIsInvalid, "BuildMessageID")

		return "", err
	}

	id := MessageID(value)

	return id, nil
}

func RebuildMessageID(value string) MessageID {
	return MessageID(value)
}

func (id MessageID) String() string {
	return string(id)
}

func (id MessageID) Equals(other MessageID) bool {
	return id.String() == other.String()
}
