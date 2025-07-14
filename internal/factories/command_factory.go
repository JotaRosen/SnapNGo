package factories

import (
	"snap-n-go/internal/commands"
	"snap-n-go/internal/dbstrategies"
	"snap-n-go/internal/types"
)

var CommandFactory = map[string]func(dbstrategies.DBStrategy, map[string]types.ConnectionParams) commands.Command{
	"ping": func(s dbstrategies.DBStrategy, _ map[string]types.ConnectionParams) commands.Command {
		return &commands.PingCommand{Strategy: s}
	},
	"backup": func(s dbstrategies.DBStrategy, _ map[string]types.ConnectionParams) commands.Command {
		return &commands.BackUpCommand{Strategy: s}
	},
	// resotre ,
	// etc
}
