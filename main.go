package main

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	var host string;
	var port int;
	var username string;
	var password string;

	ascii := `
 ________         __  __           
/        |       /  |/  |          
$$$$$$$$/______  $$ |$$ | __    __ 
   $$ | /      \ $$ |$$ |/  |  /  |
   $$ | $$$$$$  |$$ |$$ |$$ |  $$ |
   $$ | /    $$ |$$ |$$ |$$ |  $$ |
   $$ |/$$$$$$$ |$$ |$$ |$$ \__$$ |
   $$ |$$    $$ |$$ |$$ |$$    $$ |
   $$/  $$$$$$$/ $$/ $$/  $$$$$$$ |
                         /  \__$$ |
                         $$    $$/ 
                          $$$$$$/  

	`;

	app := &cli.App{
		Name:  "tallycli",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "127.0.0.1",
				Usage: "TallyDB server host",
				Destination: &host,
			},
			&cli.IntFlag{
				Name:  "port",
				Value: 4053,
				Usage: "TallyDB server port",
				Destination: &port,
			},
			&cli.StringFlag{
				Name:  "username",
				Aliases: []string{ "u" },
				Value: "username",
				Usage: "Username",
				Destination: &username,
			},
			&cli.StringFlag{
				Name:  "password",
				Aliases: []string{ "p" },
				Value: "password",
				Usage: "Password",
				Destination: &password,
			},
		},
		Usage: "CLI tool to connect to TallyDB servers",
		Action: func(*cli.Context) error {
			color.Red(ascii)

			color.White("✅ TallyCLI v1.0.0")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}