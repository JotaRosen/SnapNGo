package dbstrategies

type DBStrategy interface {
	Ping() error
}
