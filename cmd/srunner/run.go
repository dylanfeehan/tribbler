package main

import "github.com/dylanfeehan/tribbler/pkg/tribserver"

func main() {
	tribserver.RunServer("localhost", "8080")
}
