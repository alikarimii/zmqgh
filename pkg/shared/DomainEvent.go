package shared

type DomainEvent interface {
	Meta() EventMeta
	IsFailureEvent() bool
	FailureReason() error
}
