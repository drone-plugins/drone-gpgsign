package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "gpgsign plugin"
	app.Usage = "gpgsign plugin"
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "key",
			Usage:  "use this private gpg key",
			EnvVar: "PLUGIN_KEY,GPGSIGN_KEY,GPG_KEY",
		},
		cli.StringFlag{
			Name:   "passphrase",
			Usage:  "passphrase for the key",
			EnvVar: "PLUGIN_PASSPHRASE,GPGSIGN_PASSPHRASE,GPG_PASSPHRASE",
		},
		cli.BoolFlag{
			Name:   "detach-sign",
			Usage:  "append detach sign flag",
			EnvVar: "PLUGIN_DETACH_SIGN",
		},
		cli.BoolFlag{
			Name:   "clear-sign",
			Usage:  "append clear sign flag",
			EnvVar: "PLUGIN_CLEAR_SIGN",
		},
		cli.StringSliceFlag{
			Name:   "files",
			Usage:  "list of files to sign",
			EnvVar: "PLUGIN_FILES,PLUGIN_FILE",
		},
		cli.StringSliceFlag{
			Name:   "excludes",
			Usage:  "list of exclude patters",
			EnvVar: "PLUGIN_EXCLUDES,PLUGIN_EXCLUDE",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Config: Config{
			Key:        c.String("key"),
			Passphrase: c.String("passphrase"),
			Detach:     c.Bool("detach-sign"),
			Clear:      c.Bool("clear-sign"),
			Files:      c.StringSlice("files"),
			Excludes:   c.StringSlice("excludes"),
		},
	}

	if plugin.Config.Key == "" {
		return errors.New("Missing private key")
	}

	if len(plugin.Config.Files) == 0 {
		return errors.New("Missing files list")
	}

	return plugin.Exec()
}
