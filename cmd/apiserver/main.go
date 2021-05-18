package main

import (
	"github.com/zn11ch/SimpleWebApp/internal/app/apiserver"
)

func main() {

	api := apiserver.New()
	api.Start()
}
