package executors

import (
	"os"
	"snap-n-go/internal/factories"
	"snap-n-go/internal/logger"
	"snap-n-go/internal/types"
)

func Single(cp types.ConnectionParams, l *logger.Logger) {

	l.Info().Msg("Starting single command execution")
	strategyFn, ok := factories.StrategyFactory[cp.Engine]
	if !ok {
		l.Error().Msg("unsupported DB: " + cp.Engine)
		os.Exit(1)
	}

	commandFn, ok := factories.CommandFactory[cp.Command]
	if !ok {
		l.Error().Msg("unsupported Command: " + cp.Command)
		os.Exit(1)
	}

	strategy := strategyFn(cp)
	cmd := commandFn(strategy, cp)

	l.Info().Msg("Executing command: " + cp.Command)

	if err := cmd.Execute(); err != nil {
		l.Error().AnErr("Error while executing "+cp.Command+" on DBMS: "+cp.Engine, err).Send()
		os.Exit(1)
	}

	l.Info().Msg("operation completed successfully.")
}
