package app

import "github.com/urfave/cli"

// A função gerar vai retornar a aplicação de linha de comando, pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPS e Nomes de Servidor na Internet"

	return app
}
