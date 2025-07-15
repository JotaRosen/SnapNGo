package commands

import "snap-n-go/internal/dbstrategies"

type RestoreCommand struct {
	Strategy dbstrategies.DBStrategy
}

func (p *RestoreCommand) Execute() error {
	return p.Strategy.Restore()
}
