package factories

import (
	"snap-n-go/internal/commands"
	"snap-n-go/internal/dbstrategies"
	"snap-n-go/internal/types"
)

var CommandFactory = map[string]func(dbstrategies.DBStrategy, types.ConnectionParams) commands.Command{
	"ping": func(s dbstrategies.DBStrategy, _ types.ConnectionParams) commands.Command {
		return &commands.PingCommand{Strategy: s}
	},
	"backup": func(s dbstrategies.DBStrategy, _ types.ConnectionParams) commands.Command {
		return &commands.BackUpCommand{Strategy: s}
	},
	"restore": func(s dbstrategies.DBStrategy, _ types.ConnectionParams) commands.Command {
		return &commands.RestoreCommand{Strategy: s}
	},
	// incrementeal backup
	// diff backup
}
