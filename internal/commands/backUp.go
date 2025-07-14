package commands

import "snap-n-go/internal/dbstrategies"

type BackUpCommand struct {
	Strategy dbstrategies.DBStrategy
}

func (p *BackUpCommand) Execute() error {
	return p.Strategy.BackUp()
}
