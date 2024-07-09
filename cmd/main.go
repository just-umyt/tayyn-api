package main

import (
	"log"
	"os"

	"github.com/just-umyt/blUg/internal/env"
	"github.com/just-umyt/blUg/internal/router"
)

func main() {

	env.LoadEnv()

	app := router.NewApp()

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}