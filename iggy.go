package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/rgbkrk/iggy/langs"
)

func main() {
	app := cli.NewApp()
	app.Name = "iggy"

	//commands := make([]cli.Command, len(langs.Gitignores))

	commands := []cli.Command{}

	for l := range langs.Gitignores {
		lang := l
		command := cli.Command{
			Name:  lang,
			Usage: "",
			Action: func(c *cli.Context) {
				gitignore, ok := langs.Gitignores[lang]

				if !ok {
					panic(ok)
				}

				f, err := os.Create(".gitignore")
				if err != nil {
					panic(err)
				}
				defer f.Close()

				f.Write(gitignore)
				f.Sync()
			},
		}

		commands = append(commands, command)
	}

	app.Commands = commands

	app.Run(os.Args)

}
