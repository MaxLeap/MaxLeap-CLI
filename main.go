package main

import (
	"fmt"
	"os"
	"github.com/howeyc/gopass"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "zcc"
	app.Usage = "zcloud code command line"
	app.Version = "0.1"
	if exists(getSessionPath()) == false && (len(os.Args) > 1 && os.Args[1] != "login") {
		fmt.Println("please login first,use 'login username'")
		return
	}
	host=getHostString()
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "login <username> [-region <region,default:"+ region+ "("+host+")>]",
			Action: func(c *cli.Context) {
				user := c.Args().Get(0)
				checkStrArg(user)
				region=c.String("region")
				if region=="CN" {
					host=CN
				}else if region=="US"{
					host=US
				}else {
					host=region
				}
				for i := 0; i < 3; i++ {
					fmt.Print("enter password:")
					passwd:=string(gopass.GetPasswd())
					if login(user, passwd) {
						fmt.Println("login success")
						persistHostString()
						break
					} else {
						if i < 2 {
							fmt.Println("Permission denied, please try again.")
						} else {
							fmt.Println("Permission denied")
						}
					}
				}

			},
			Flags:[]cli.Flag{
				cli.StringFlag{
					Name:  "region",
					Value: region,
					Usage: "choose region,US or CN ...",
				},
			},
		},
		{
			Name:  "use",
			Usage: "use <appname>",
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
			Usage: "upload <filepath>",
			Action: func(c *cli.Context) {
				path := c.Args().First()
				fn := func() int {
					return getApp().upload(path)
				}
				fmt.Print("upload")
				startWithProgress(fn)
			},
		},
		{
			Name:  "log",
			Usage: "log [-l <info|error>] [-n <number of log>] [-s <number of skipped log>]",
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
			Usage: "deploy <version>",
			Action: func(c *cli.Context) {
				version := c.Args().First()
				checkStrArg(version)
				fmt.Print("deploy")
				startWithProgress(func() int { return getApp().deploy(version) })
			},
		},
		{
			Name:  "lv",
			Usage: "lv",
			Action: func(c *cli.Context) {
				getApp().listAppVersions()
			},
		},
		{
			Name:  "undeploy",
			Usage: "undeploy <version>",
			Action: func(c *cli.Context) {
				version := c.Args().First()
				checkStrArg(version)
				fmt.Print("undeploy")
				startWithProgress(func() int { return getApp().undeploy(version) })

			},
		},
		{
			Name:  "logout",
			Usage: "logout",
			Action: func(c *cli.Context) {
				clear()
			},
		},
	}
	app.Run(os.Args)
}
