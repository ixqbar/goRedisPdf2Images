package main

import (
	"app/common"
	"app/core"
	"github.com/urfave/cli"
	"os"
)

func main()  {
	app := cli.NewApp()
	app.Name = "goRedisPdf2Image"
	app.Version = common.VERSION
	app.Description = "pdf to images"
	app.Usage = ""

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config,c",
			Usage: "load configuration from `FILE`",
		},
	}

	app.Action = func(c *cli.Context) error {
		configFile := c.String("config")
		if len(configFile) == 0 {
			cli.ShowAppHelp(c)
			return nil
		}

		common.Logger.Printf("run with config file %s", configFile)

		if _, err := common.ParseConfig(configFile); err != nil {
			return err
		}

		return core.Run()
	}

	err := app.Run(os.Args)
	if err != nil {
		common.Logger.Print(err)
	}
}
