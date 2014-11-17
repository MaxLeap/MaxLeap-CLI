package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "zcc"
	app.Usage = "zcloud code command line"
	app.Version = "0.1"
	////cmd := os.Args[1]
	//cmd := "deploy"
	//if cmd == "login" {
	//	user := os.Args[2]
	//	passwd := os.Args[3]
	//	login(user, passwd)
	//}
	//if cmd == "deploy" {
	//	user := "/Users/ben/Downloads/cloud-code-template-java-1.0-SNAPSHOT-mod.zip"
	//	deploy(user)
	//}
	//if cmd == "listapps" {
	//	showApps()
	//}
	//if cmd == "use" {
	//	use(os.Args[2])
	//}
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "login username password",
			Action: func(c *cli.Context) {
				user := c.Args().First()
				passwd := c.Args().First()
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
			Name:  "lsapps",
			Usage: "",
			Action: func(c *cli.Context) {
				showApps()

			},
		},
		{
			Name:  "deploy",
			Usage: "deploy filepath",
			Action: func(c *cli.Context) {
				path := c.Args().First()
				deploy(path)

			},
		},
	}
	app.Run(os.Args)
}
