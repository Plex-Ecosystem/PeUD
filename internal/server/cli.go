package server

import (
	"os"
	"sort"

	"github.com/DirtyCajunRice/PeUD/internal/handlers"

	"github.com/urfave/cli/v2"
)

func CLI(version, date *string, Env *handlers.Env) {
	app := &cli.App{
		Name:     "PeUD",
		HelpName: "peud",
		Version:  *version,
		Usage:    "Plex Ecosystem User Database",
		Action: func(c *cli.Context) error {
			// TODO: Add validation for multi-option flags like database type
			Start(version, date, Env)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "server-port",
				Aliases:     []string{"p"},
				Usage:       "Server listening port",
				EnvVars:     []string{"SERVER_PORT"},
				Value:       8888,
				Destination: &Env.Config.APIServer.Port,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Usage:       "Enable debug logging",
				EnvVars:     []string{"DEBUG"},
				Value:       false,
				Destination: &Env.Config.Debug,
			},
			&cli.StringFlag{
				Name:        "address",
				Aliases:     []string{"a"},
				Usage:       "Server hostname or IP",
				EnvVars:     []string{"SERVER_ADDRESS"},
				Value:       "0.0.0.0",
				Destination: &Env.Config.APIServer.Address,
			},
			&cli.StringFlag{
				Name:        "database-type",
				Aliases:     []string{"T"},
				Usage:       "Database Backend",
				EnvVars:     []string{"DATABASE_TYPE"},
				Value:       "sqlite",
				Destination: &Env.Config.Database.Type,
			},
			&cli.StringFlag{
				Name:        "database-name",
				Aliases:     []string{"t"},
				Usage:       "Database Name (For sqlite it is the filename)",
				EnvVars:     []string{"DATABASE_NAME"},
				Value:       "peud.db",
				Destination: &Env.Config.Database.Name,
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	log := Env.Log
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
