package main

import (
	"app_redes/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
