package executors

import (
	"encoding/json"
	"fmt"
	"os"
	"snap-n-go/internal/factories"
	"snap-n-go/internal/logger"
	"snap-n-go/internal/types"
	"strconv"
	"sync"
)

func Concurrent(dbFile string, l *logger.Logger) {

	l.Info().Msg("Starting Concurrent command execution")
	cpArr := getConnectionParams(dbFile)

	var wg sync.WaitGroup

	for i, cp := range cpArr {

		wg.Add(1)

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

		l.Info().Msg("Executing Config: " + strconv.Itoa(i))

		go func() {
			defer wg.Done()
			err := cmd.Execute()
			if err != nil {
				l.Error().AnErr("Error while executing "+cp.Command+" on DBMS: "+cp.Engine, err).Send()
				os.Exit(1)
			} else {
				defer l.Info().Msg("operation completed successfully for command: " + cp.Command)
			}
		}()
	}

	wg.Wait()

}

func getConnectionParams(filePath string) []types.ConnectionParams {

	dir, _ := os.Getwd()

	byteArr, err := os.ReadFile(dir + "/" + filePath)
	if err != nil {
		fmt.Errorf("Error reading file: %v", err)
	}

	var cpArr []types.ConnectionParams

	uErr := json.Unmarshal(byteArr, &cpArr)
	if uErr != nil {
		fmt.Errorf("Error unmarshalling JSON: %v", uErr)
	}

	return cpArr
}
