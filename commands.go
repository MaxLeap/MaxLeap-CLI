package main
import (
	"fmt"
	"github.com/howeyc/gopass"
    "github.com/benile/cli"
	"github.com/benile/readlikeflags"
)

func getCommands()[]cli.Command{
	return []cli.Command{
		{
			Name:  "login",
			Usage: "login <username> -region <region,CN or US>",
			Action: func(c *cli.Context) {
				user := c.Args().Get(0)
				checkStrArg(user)
				region=c.String("region")
				if region==""{
					fmt.Println("miss region,please use -region <CN or US ...> to define the region")
					return
				}
				for i := 0; i < 3; i++ {
					fmt.Print("enter password:")
					passwd:=string(gopass.GetPasswd())
					if login(user, passwd) {
						//fmt.Println("login success")
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
				handle(func(a *app) {
					fmt.Print("upload")
					startWithProgress(func()int { return a.upload(path)})
				})
			},
		},
		{
			Name:  "log",
			Usage: "log [-l <info|error>] [-n <number of log>] [-s <number of skipped log>]",
			Action: func(c *cli.Context) {
				handle(func(a *app) {a.log(c.String("l"), c.Int("n"), c.Int("s"))  })
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
				handle(func (a *app){
					startWithProgress(func()int { return a.deploy(version) })
				});
			},
		},
		{
			Name:  "lv",
			Usage: "lv",
			Action: func(c *cli.Context) {
				handle(func(a *app) {
					fmt.Println(a.listAppVersions())})
			},
		},
		{
			Name:  "undeploy",
			Usage: "undeploy <version>",
			Action: func(c *cli.Context) {
				version := c.Args().First()
				checkStrArg(version)
				fmt.Print("undeploy")
				handle(func(a *app) { startWithProgress(func() int { return a.undeploy(version) }) })

			},
		},
		{
			Name:  "logout",
			Usage: "logout",
			Action: func(c *cli.Context) {
				readlikeflags.Exit()
			},
		},
	}
}
