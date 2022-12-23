package main

import (
	"log"

	"github.com/karanbirsingh7/go-grpc/internal/db"
	"github.com/karanbirsingh7/go-grpc/internal/rocket"
)

func main() {
	log.Println("Starting")

	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

// Run is responsbile for initializing and starting our gRPC server
func Run() error {
	// connec to DB
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	// DB migrate on startup
	if err := rocketStore.Migrate(); err != nil {
		return err
	}

	//
	_ = rocket.New(&rocketStore)
	return nil
}
