package dashboard

type Services interface {
	GetTotal(feature string) (int, error)
}

type Repositories interface {
	CountData(table string) (int, error)
	CountQueueData() (int, error)
}
