package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "zcc"
	app.Usage = "zcloud code command line"
	app.Version = "0.1"
	if exists(getSessionPath()) == false && (len(os.Args) > 1 && os.Args[1] != "login") {
		fmt.Println("please login first,use 'login username password'")
		return
	}
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "login username password",
			Action: func(c *cli.Context) {
				user := c.Args().Get(0)
				passwd := c.Args().Get(1)
				login(user, passwd)

			},
		},
		{
			Name:  "use",
			Usage: "use appname",
			Action: func(c *cli.Context) {
				app := c.Args().First()
				use(app)

			},
		},
		{
			Name:  "apps",
			Usage: "",
			Action: func(c *cli.Context) {
				showApps()

			},
		},
		{
			Name:  "upload",
			Usage: "upload filepath",
			Action: func(c *cli.Context) {
				path := c.Args().First()
				getApp().upload(path)
			},
		},
		{
			Name:  "log",
			Usage: "log [-l <info|error>] [-n <number of log>] [-s <number of skipped log>] <appid>",
			Action: func(c *cli.Context) {
				getApp().log(c.String("l"), c.Int("n"), c.Int("s"))
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "level,l",
					Value: "info",
					Usage: "log level",
				},
				cli.IntFlag{
					Name:  "n",
					Value: 10,
					Usage: "number of row shown onetime",
				},
				cli.IntFlag{
					Name:  "s",
					Value: -1,
					Usage: " number of row skipped",
				},
			},
		},
		{
			Name:  "deploy",
			Usage: "deploy version",
			Action: func(c *cli.Context) {
				version := c.Args().First()
				getApp().deploy(version)
			},
		},
		{
			Name:  "lv",
			Usage: "list version",
			Action: func(c *cli.Context) {
				getApp().listAppVersions()
			},
		},
		{
			Name:  "undeploy",
			Usage: "undeploy",
			Action: func(c *cli.Context) {
				getApp().undeploy()
			},
		},
	}
	app.Run(os.Args)
}
