package main

import "log"

func main() {
	log.Println("Starting")

	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

// Run is responsbile for initializing and starting our gRPC server
func Run() error {
	return nil
}
