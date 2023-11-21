package main

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/version"
	"log"
)

func init() {
	log.Println(version.Name)
	log.Printf("Application Version: %s", version.Version)
}

func main() {

}
