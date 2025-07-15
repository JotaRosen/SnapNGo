package factories

import (
	"snap-n-go/internal/dbstrategies"
	"snap-n-go/internal/types"
)

// Recibe parámetros y retorna la estrategia adecuada
var StrategyFactory = map[string]func(types.ConnectionParams) dbstrategies.DBStrategy{
	"mongo": func(cp types.ConnectionParams) dbstrategies.DBStrategy {
		return &dbstrategies.MongoConnectionParams{
			Command:  cp.Command,
			Engine:   cp.Engine,
			Host:     cp.Host,
			Port:     cp.Port,
			Username: cp.Username,
			Password: cp.Password,
			DbName:   cp.DbName,
		}
	},
}
