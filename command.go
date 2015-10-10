package main
import (
	"fmt"

)
type command struct {
	Name string
	Usage string
	Action func( * Context)
	Flags []Flag
}
type commands struct {
	commands []command
}
func newCommands() commands{
	return commands{getCommands()}
}
func(cmds commands) getCommand(name string) *command{
	for _,cmd:=range(cmds.commands){
		if cmd.Name==name{
			return &cmd
		}
	}
	return nil
}
func getCommands()[]command{
	return []command{
		{
			Name:  "use",
			Usage: "use <appname>",
			Action: func(c *Context) {
				app := c.Args().First()
				use(app)

			},
		},
		{
			Name:  "apps",
			Usage: "",
			Action: func(c *Context) {
				showApps()

			},
		},
		{
			Name:  "upload",
			Usage: "upload <filepath>",
			Action: func(c *Context) {
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
			Action: func(c *Context) {
				getApp().log(c.String("l"), c.Int("n"), c.Int("s"))
			},
			Flags: []Flag{
				StringFlag{
					Name:  "level,l",
					Value: "info",
					Usage: "log level",
				},
				IntFlag{
					Name:  "n",
					Value: 10,
					Usage: "number of row shown onetime",
				},
				IntFlag{
					Name:  "s",
					Value: -1,
					Usage: " number of row skipped",
				},
			},
		},
		{
			Name:  "deploy",
			Usage: "deploy <version>",
			Action: func(c *Context) {
				version := c.Args().First()
				checkStrArg(version)
				fmt.Print("deploy")
				startWithProgress(func() int { return getApp().deploy(version) })
			},
		},
		{
			Name:  "lv",
			Usage: "lv",
			Action: func(c *Context) {
				getApp().listAppVersions()
			},
		},
		{
			Name:  "undeploy",
			Usage: "undeploy <version>",
			Action: func(c *Context) {
				version := c.Args().First()
				checkStrArg(version)
				fmt.Print("undeploy")
				startWithProgress(func() int { return getApp().undeploy(version) })

			},
		},
		{
			Name:  "logout",
			Usage: "logout",
			Action: func(c *Context) {
				clear()
			},
		},
	}
}
