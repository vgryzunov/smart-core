package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

func Server(cmd *cobra.Command, agrs []string) {

	log.Printf(cmd.Short)

	log.Print("Config file used: ", viper.ConfigFileUsed())
	port := viper.GetString(HttpPortFlag)
	log.Printf("Using HTTP Port: %s", port)
}

func tokenAuthOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("*** tokenAuthOnly handler ****")
		headers := r.Header
		xAuthToken := headers.Get("X-Auth-Token")
		if xAuthToken == "" {
			log.Printf("Not Authenticated. Missing authentication X-Auth-Token")
			http.Error(w, http.StatusText(403), 403)
			return
		}
		next.ServeHTTP(w, r)
	})
}
