package commands

import "snap-n-go/internal/dbstrategies"

type PingCommand struct {
	Strategy dbstrategies.DBStrategy
}

func (p *PingCommand) Execute() error {
	return p.Strategy.Ping()
}
