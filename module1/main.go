package main

import (
	"log"

	"github.com/xxii22w/grpc-example/proto"
)

func main() {
	// protoc --go_out=. --go_opt=paths=source_relative proto/hello.proto
	person := proto.Person{
		Name: "Cgh",
	}

	log.Println(person.GetName())
}