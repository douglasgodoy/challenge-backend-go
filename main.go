package main

import (
	env "api/infra/dotenv"
	db "api/infra/pg"
	"api/router"
)

func main() {
	env.InitializeEnvs()
	db.InitializePg()
	router.Initialize()
}
