package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/rgbkrk/iggy/langs"
)

func main() {
	app := cli.NewApp()
	app.Name = "iggy"

	//commands := make([]cli.Command, len(langs.Gitignores))

	commands := []cli.Command{}

	for lang, gitignore := range langs.Gitignores {
		command := cli.Command{
			Name:  lang,
			Usage: fmt.Sprintf("write the gitignore for %s", lang),
			Action: func(c *cli.Context) {
				f, err := os.Create(".gitignore")
				if err != nil {
					panic(err)
				}
				defer f.Close()

				f.WriteString(gitignore)
				f.Sync()
			},
		}

		commands = append(commands, command)
	}

	app.Commands = commands

	app.Run(os.Args)

}
