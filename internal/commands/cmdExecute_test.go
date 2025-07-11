package commands

import (
	"snap-n-go/internal/dbstrategies"
	"testing"
)

func TestPingCommand_Execute(t *testing.T) {
	// Test the PingCommand with MongoDB strategy
	mongoParams := &dbstrategies.MongoConnectionParams{
		Host: "invalid-host-for-testing",
		Port: "27017",
	}

	pingCmd := &PingCommand{
		Strategy: mongoParams,
	}

	err := pingCmd.Execute()
	if err == nil {
		t.Errorf("Expected error when pinging invalid host, but got nil")
	}
}
