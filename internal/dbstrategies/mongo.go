package dbstrategies

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
	"os/exec"
	"snap-n-go/internal/types"
	"time"
)

type MongoConnectionParams types.ConnectionParams

func buildURI(cp *MongoConnectionParams) string {
	var uri string
	// Construcción básica de la URI
	if cp.Username != "" && cp.Password != "" {
		// Escapar las credenciales para evitar problemas con caracteres especiales
		username := url.QueryEscape(cp.Username)
		password := url.QueryEscape(cp.Password)
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/", username, password, cp.Host, cp.Port)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s/", cp.Host, cp.Port)
	}

	// Agregar base de datos si está especificada
	if cp.DbName != "" {
		uri += "/" + cp.DbName
	}
	// Agregar parámetros de query por defecto
	params := []string{
		"authSource=admin", // Base de datos de autenticación por defecto
	}

	// Agregar parámetros a la URI
	uri += "?"
	for i, param := range params {
		if i > 0 {
			uri += "&"
		}
		uri += param
	}

	return uri
}

func (cp *MongoConnectionParams) getMongoClient() (*mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(buildURI(cp)))
	if err != nil {
		cancel()
	}
	return client, ctx, cancel
}

func (cp *MongoConnectionParams) Ping() error {
	client, ctx, cancel := cp.getMongoClient()

	//Both of these func need to be execute after command exeution

	defer cancel()
	defer client.Disconnect(ctx)
	return client.Ping(ctx, nil)
}

func (cp *MongoConnectionParams) BackUp() error {
	// The mongodump tool will establish its own connection.
	// Base arguments for the mongodump command

	//backupPath := "snapshot-" + time.Now().Format(time.RFC3339) // RFC3339  = "2006-01-02T15:04:05Z07:00"
	args := []string{
		"--host", cp.Host,
		"--port", cp.Port,
	}

	// Conditionally add arguments based on the connection parameters
	if cp.Username != "" {
		args = append(args, "--username", cp.Username)
	}
	if cp.Password != "" {
		args = append(args, "--password", cp.Password)
		// When using credentials, you often need to specify the authentication database.
		args = append(args, "--authenticationDatabase", "admin")
	}

	// If a specific database name is provided, only back up that database.
	// Otherwise, mongodump will back up all databases on the server.
	if cp.DbName != "" {
		backupPath := "snapshot-" + cp.DbName + "-" + time.Now().Format(time.RFC3339) // RFC3339  = "2006-01-02T15:04:05Z07:00"
		args = append(args, "--db", cp.DbName, "--out", backupPath)
	}

	// Create the command with our arguments
	cmd := exec.Command("mongodump", args...)

	// Execute the command and capture both stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		// If the command fails, return a detailed error including the output from mongodump,
		// which is very helpful for debugging connection issues or permissions problems.

		fmt.Errorf("mongodump command failed: %w\nOutput: %s", err, string(output))
		return err
	}
	return nil
}
