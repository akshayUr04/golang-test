package main

import "golang-test/pkg/api"

func main() {
	server := api.NewHttpServer()
	server.Start()
}
