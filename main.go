package main

import (
	"fmt"
	"os"
	"snap-n-go/cmd"
	"snap-n-go/internal/logger"
)

// Initialize the logger
var l *logger.Logger

func init() {
	var err error
	l, err = logger.NewLogger("", "main")
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
}

func main() {
	// Pass the logger to the root command
	l.Info().Msg("... Starting SnapNGo ...")
	if err := cmd.Execute(l); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
