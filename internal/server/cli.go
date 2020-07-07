package server

import (
	"os"
	"sort"

	"github.com/DirtyCajunRice/PeUD/internal/handlers"

	"github.com/urfave/cli/v2"
)

func CLI(env *handlers.Env) {
	app := &cli.App{
		Name:     "PeUD",
		HelpName: "peud",
		Version:  *env.Build.Version,
		Usage:    "Plex Ecosystem User Database",
		Action: func(c *cli.Context) error {
			// TODO: InsertPlexUsers validation for multi-option flags like database type
			env.Config.LoadFromEnv()
			Start(env)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "server-port",
				Aliases:     []string{"p"},
				Usage:       "Server listening port",
				EnvVars:     []string{"SERVER_PORT"},
				Value:       8888,
				Destination: &env.Config.APIServer.Port,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Usage:       "Enable debug logging",
				EnvVars:     []string{"DEBUG"},
				Value:       false,
				Destination: &env.Config.Debug,
			},
			&cli.StringFlag{
				Name:        "address",
				Aliases:     []string{"a"},
				Usage:       "Server hostname or IP",
				EnvVars:     []string{"SERVER_ADDRESS"},
				Value:       "0.0.0.0",
				Destination: &env.Config.APIServer.Address,
			},
			&cli.StringFlag{
				Name:        "database-type",
				Aliases:     []string{"T"},
				Usage:       "Database Backend [sqlite3, mysql, postgres]",
				EnvVars:     []string{"DATABASE_TYPE"},
				Value:       "sqlite3",
				Destination: &env.Config.Database.Type,
			},
			&cli.StringFlag{
				Name:        "database-name",
				Aliases:     []string{"t"},
				Usage:       "Database Name",
				EnvVars:     []string{"DATABASE_NAME"},
				Value:       "peud.db",
				Destination: &env.Config.Database.Name,
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	log := env.Log
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
