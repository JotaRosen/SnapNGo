package dbstrategies

type DBStrategy interface {
	Ping() error
	BackUp() error
}
