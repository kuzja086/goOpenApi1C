package main

import (
	"goOpenApi1C/internal/commands"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "generate",
		Version: "0.0.1",
		Usage:   "Генерирует спецификацию OpenApi для Http-сервисов 1С",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Value: "./src",
				Usage: "Путь к каталогу с сервисами",
			},
			&cli.StringFlag{
				Name:  "out",
				Value: "./build",
				Usage: "Каталог выгрузки результатов генерации",
			},
			&cli.StringFlag{
				Name:  "service",
				Usage: "Название сервиса для генерации",
			},
		},
		Action: func(c *cli.Context) error {
			commands.GenerateOpenApi(c)
			return log.Output(0, "Sucsess")
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
