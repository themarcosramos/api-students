package main

import (
	"log"

	"github.com/themarcosramos/api-students/api"
)

func main() {

  server := api.NewServer()	

  // Routes
  server.ConfigureRoutes()

  // Start server
  if  err := server.Start(); err!= nil {
	log.Fatal(err)
  }

}
