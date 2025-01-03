package main

import (
	"flag"
	"log"
	"pregogen"
)

func main() {
	typeName := flag.String("type", "", "type name to generate methods for")
	fileName := flag.String("file", "", "Go source file to parse containing the struct type")
	gentype := flag.String("gen", "bytesBuffer", "use this method to generate")
	flag.Parse()

	if *fileName == "" {
		log.Fatal("file flag is required")
	}

	log.Println("generate type:", *gentype)

	err := pregogen.RunGenerator(*typeName, *fileName, *gentype)
	if err != nil {
		log.Panic(err)
	}
}
