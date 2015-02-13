package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/ledis"
	"github.com/codegangsta/cli"

	"github.com/dockercn/wharf/email"
	"github.com/dockercn/wharf/models"
	_ "github.com/dockercn/wharf/routers"
)

var CmdWeb = cli.Command{
	Name:        "web",
	Usage:       "Start Wharf Web Service",
	Description: "Wharf is ContainerOps platform for the enterprise include Docker repositories storage, CI/CD and so on.",
	Action:      runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Value: "0.0.0.0",
			Usage: "Web listen IP, default: 0.0.0.0",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "80",
			Usage: "Web listen port, default: 80",
		},
		cli.StringFlag{
			Name:  "email",
			Value: "false",
			Usage: "Start email service",
		},
		cli.StringFlag{
			Name:  "conf",
			Value: "",
			Usage: "Web application conf path",
		},
	},
}

func runWeb(c *cli.Context) {
	var address, port string

	if len(c.String("address")) > 0 {
		address = strings.TrimSpace(c.String("address"))
	}

	if len(c.String("port")) > 0 {
		port = strconv.Itoa(c.Int("port"))
	}

	models.InitDb()

	if len(c.String("email")) > 0 {
		e, _ := strconv.ParseBool(c.String("email"))

		if e == true {
			email.StartService()
		}
	}

	beego.SetStaticPath(beego.AppConfig.String("docker::StaticPath"), fmt.Sprintf("%s/images", beego.AppConfig.String("docker::BasePath")))
	beego.SetStaticPath(beego.AppConfig.String("docker::Gravatar"), beego.AppConfig.String("docker::Gravatar"))

	beego.Run(fmt.Sprintf("%v:%v", address, port))
}
