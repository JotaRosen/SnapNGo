package dbstrategies

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
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

func (cp *MongoConnectionParams) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(buildURI(cp)))
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx) //defer needed since we are client is returned and exectuion es after function clousure

	return client.Ping(ctx, nil)
}
