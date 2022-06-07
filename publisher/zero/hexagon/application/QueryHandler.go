package application

func NewQueryHandler(
	forMessageCount ForMessageCount,
) *QueryHanndler {
	return &QueryHanndler{forMessageCount}
}

type QueryHanndler struct {
	forMessageCount ForMessageCount
}

func (queryHandler QueryHanndler) MessageCount() int64 {
	return queryHandler.forMessageCount()
}
