package server

import (
	"os"
	"sort"

	"github.com/DirtyCajunRice/PeUD/internal/handlers"

	"github.com/urfave/cli/v2"
)

func CLI(version, date *string, handlerEnv *handlers.Env) {
	app := &cli.App{
		Name:     "PeUD",
		HelpName: "peud",
		Version:  *version,
		Usage:    "Plex Ecosystem User Database",
		Action: func(c *cli.Context) error {
			Start(version, date, handlerEnv)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "server-port",
				Aliases:     []string{"p"},
				Usage:       "Server listening port",
				EnvVars:     []string{"SERVER_PORT"},
				Value:       8888,
				Destination: &handlerEnv.Config.APIServer.Port,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Usage:       "Enable debug logging",
				EnvVars:     []string{"DEBUG"},
				Value:       false,
				Destination: &handlerEnv.Config.Debug,
			},
			&cli.StringFlag{
				Name:        "address",
				Aliases:     []string{"a"},
				Usage:       "Server hostname or IP",
				EnvVars:     []string{"SERVER_ADDRESS"},
				Value:       "0.0.0.0",
				Destination: &handlerEnv.Config.APIServer.Address,
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	log := handlerEnv.Log
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
