package main

import (
	"fmt"
	"os"
	"cli"
	"github.com/benile/readlikeflags"
	"github.com/howeyc/gopass"
)

func main() {
	app := cli.NewApp()
	app.Name = "maxleap"
	app.Usage = "zcloud code command line"
	app.Version = "0.2"
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "login <username> -region <region,CN or US>",
			Action: func(c *cli.Context) {
				user := c.Args().Get(0)
				checkStrArg(user)
				region = c.String("region")
				if region == "" {
					fmt.Println("miss region,please use -region <CN or US ...> to define the region")
					return
				}
				for i := 0; i < 3; i++ {
					fmt.Print("enter password:")
					passwd := string(gopass.GetPasswdMasked())
					if login(user, passwd) {
						errorHandler:= func(err error) {
							fmt.Println(err.Error())
						}
						readlikeflags.StartSession(
							readlikeflags.Options{
								getCommands(),nil,app.Usage,app.Version,errorHandler,app.Name})
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
	}
	app.Run(os.Args)
}
