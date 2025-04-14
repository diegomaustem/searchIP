package app

import "github.com/urfave/cli"

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Application of comand line"
	app.Usage = "Search IP'S and names of server in the internet"
	return app
}