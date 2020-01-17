package main

import (
	"fmt"
	"log"
	"net/http"
	"project1/config"
	usercontroller "project1/controllers/user"
	"time"
)

type CustomMux struct {
	http.ServeMux
}

func (c CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if config.Configuration().Log.Verbose {
		log.Println("Incoming request from", r.Host, "accessing", r.URL.String())
	}
	c.ServeMux.ServeHTTP(w, r)
}

func main() {
	router := new(CustomMux)

	router.HandleFunc("/user", usercontroller.FindAll)

	server := new(http.Server)
	server.Handler = router
	server.ReadTimeout = config.Configuration().Server.ReadTimeout * time.Second
	server.WriteTimeout = config.Configuration().Server.WriteTimeout * time.Second
	server.Addr = fmt.Sprintf("%s:%d", config.Configuration().Server.Host, config.Configuration().Server.Port)

	if config.Configuration().Log.Verbose {
		log.Printf("Starting server at %s \n", server.Addr)
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
