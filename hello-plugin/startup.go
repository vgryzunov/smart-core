package main

import (
	"github.com/go-openapi/loads"
	"github.com/vgryzunov/smart-core/hello-plugin/swagger/restapi"
	"github.com/vgryzunov/smart-core/hello-plugin/swagger/restapi/operations"
)

type plugin string

var Plugin plugin

func Start(port int) error {
	// Load JSON generated from swagger.yml
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return err
	}

	api := operations.NewHelloAPI(swaggerSpec)

	// Create REST API server
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = port

	return server.Serve()
}
