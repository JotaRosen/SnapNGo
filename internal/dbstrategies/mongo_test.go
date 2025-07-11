package dbstrategies

import (
	"testing"
	"time"
)

func TestBuildURI(t *testing.T) {
	tests := []struct {
		name     string
		params   MongoConnectionParams
		expected string
	}{
		{
			name: "URI with username and password",
			params: MongoConnectionParams{
				Host:     "localhost",
				Port:     "27017",
				Username: "testuser",
				Password: "testpass",
				DbName:   "testdb",
			},
			expected: "mongodb://testuser:testpass@localhost:27017//testdb?authSource=admin",
		},
		{
			name: "URI without credentials",
			params: MongoConnectionParams{
				Host:   "localhost",
				Port:   "27017",
				DbName: "testdb",
			},
			expected: "mongodb://localhost:27017//testdb?authSource=admin",
		},
		{
			name: "URI without database name",
			params: MongoConnectionParams{
				Host:     "localhost",
				Port:     "27017",
				Username: "testuser",
				Password: "testpass",
			},
			expected: "mongodb://testuser:testpass@localhost:27017/?authSource=admin",
		},
		{
			name: "URI with special characters in credentials",
			params: MongoConnectionParams{
				Host:     "localhost",
				Port:     "27017",
				Username: "test@user",
				Password: "test:pass",
				DbName:   "testdb",
			},
			expected: "mongodb://test%40user:test%3Apass@localhost:27017//testdb?authSource=admin",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildURI(&tt.params)
			if result != tt.expected {
				t.Errorf("BuildURI() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestMongoConnectionParams_Ping_InvalidConnection(t *testing.T) {
	tests := []struct {
		name   string
		params MongoConnectionParams
	}{
		{
			name: "Invalid host",
			params: MongoConnectionParams{
				Host: "invalid-host-that-does-not-exist",
				Port: "27017",
			},
		},
		{
			name: "Invalid port",
			params: MongoConnectionParams{
				Host: "localhost",
				Port: "99999",
			},
		},
		{
			name: "Empty port",
			params: MongoConnectionParams{
				Host: "localhost",
				Port: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Ping()
			if err == nil {
				t.Errorf("Expected error for invalid connection parameters, but got nil")
			}
		})
	}
}

func TestMongoConnectionParams_Ping_Timeout(t *testing.T) {
	// Test that ping respects the timeout
	params := MongoConnectionParams{
		Host: "10.255.255.1", // Non-routable IP to simulate timeout
		Port: "27017",
	}

	start := time.Now()
	err := params.Ping()
	duration := time.Since(start)

	if err == nil {
		t.Errorf("Expected error for timeout, but got nil")
	}

	// Should timeout within reasonable time (5 seconds + some buffer)
	if duration > 10*time.Second {
		t.Errorf("Ping took too long: %v, expected around 5 seconds", duration)
	}
}

func TestMongoConnectionParams_ImplementsDBStrategy(t *testing.T) {
	// Test that MongoConnectionParams implements DBStrategy interface
	var strategy DBStrategy
	params := &MongoConnectionParams{
		Host: "localhost",
		Port: "27017",
	}

	strategy = params
	if strategy == nil {
		t.Errorf("MongoConnectionParams should implement DBStrategy interface")
	}

	// Test that the Ping method is available
	err := strategy.Ping()
	// We expect an error since we're not connecting to a real MongoDB instance
	if err == nil {
		t.Logf("Ping succeeded - this might indicate a real MongoDB instance is running")
	}
}

func TestMongoConnectionParams_WithCredentials(t *testing.T) {
	// Test ping with credentials (will fail but should not panic)
	params := MongoConnectionParams{
		Host:     "localhost",
		Port:     "27017",
		Username: "testuser",
		Password: "testpass",
		DbName:   "testdb",
	}

	err := params.Ping()
	// We expect an error since we're not connecting to a real MongoDB instance
	if err == nil {
		t.Logf("Ping with credentials succeeded - this might indicate a real MongoDB instance is running")
	}
}

func TestMongoConnectionParams_ContextCancellation(t *testing.T) {
	// Test that the ping operation respects context cancellation
	params := MongoConnectionParams{
		Host: "10.255.255.1", // Non-routable IP
		Port: "27017",
	}

	// This test verifies that the internal context timeout works
	// The Ping method creates its own context with 5-second timeout
	start := time.Now()
	err := params.Ping()
	duration := time.Since(start)

	if err == nil {
		t.Errorf("Expected error for unreachable host, but got nil")
	}

	// Should complete within reasonable time due to context timeout
	if duration > 8*time.Second {
		t.Errorf("Ping took too long: %v, expected to timeout around 5 seconds", duration)
	}
}
