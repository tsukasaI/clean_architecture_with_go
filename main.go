package main

import (
	"clean_architecture_with_go/adapter/api"
	"clean_architecture_with_go/infrastructure"
)

func main() {
	infrastructure.InitDB()

	e := api.InitRouter()
	e.Run()
}
