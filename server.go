package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"plugin"
)

type Plugin interface {
	Start()
}

func Server(cmd *cobra.Command, agrs []string) {

	log.Printf(cmd.Short)

	log.Print("Config file used: ", viper.ConfigFileUsed())
	port := viper.GetString(HttpPortFlag)
	log.Printf("Using HTTP Port: %s", port)

	mod := "./bin/hello-plugin.go"
	startModule(mod)
}

func startModule(modPath string) {
	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open(modPath)
	if err != nil {
		log.Fatalln(err)
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		log.Fatalln(err)
	}
	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var mod Plugin
	mod, ok := symPlugin.(Plugin)
	if !ok {
		log.Fatalln("unexpected type from module symbol")
	}

	// 4. use the module
	mod.Start()
}
