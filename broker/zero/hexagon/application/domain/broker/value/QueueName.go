package value

import (
	"github.com/alikarimii/zmqph/pkg/shared"
	"github.com/cockroachdb/errors"
)

// for sample
func BuildQueueName(q string) (QueueName, error) {
	wrapWithMsg := "BuildQueueName"
	if q == "" { //@TODO simple validation.but must use validator
		err := errors.New("empty input for queue name")
		err = shared.MarkAndWrapError(err, shared.ErrInputIsInvalid, wrapWithMsg)

		return "", err
	}
	return QueueName(q), nil
}

func RebuildQueueName(q string) QueueName {
	return QueueName(q)
}

type QueueName string

func (queueName QueueName) String() string {
	return string(queueName)
}
func (queueName QueueName) Equals(other QueueName) bool {
	return queueName.String() == other.String()
}
