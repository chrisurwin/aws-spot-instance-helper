package main

import (
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/chrisurwin/aws-spot-instance-helper/agent"
	"github.com/urfave/cli"
)

//VERSION of the program
var VERSION = "v0.1.0-dev"

func beforeApp(c *cli.Context) error {
	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "aws-spot-instance-helper"
	app.Version = VERSION
	app.Usage = "Evacuates an AWS Spot Instance host when marked for termination"
	app.Action = start
	app.Before = beforeApp
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug,d",
			Usage:  "Debug logging",
			EnvVar: "DEBUG",
		},
		cli.DurationFlag{
			Name:   "poll-interval,i",
			Value:  5 * time.Second,
			Usage:  "Polling interval for checks",
			EnvVar: "POLL_INTERVAL",
		},
		cli.StringFlag{
			Name:   "cattleURL,u",
			Usage:  "Cattle URL",
			EnvVar: "CATTLE_URL",
		},
		cli.StringFlag{
			Name:   "cattleAccessKey,ck",
			Usage:  "Cattle Access Key",
			EnvVar: "CATTLE_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "cattleSecretKey,cs",
			Usage:  "Cattle Secret Key",
			EnvVar: "CATTLE_SECRET_KEY",
		},
	}
	app.Run(os.Args)
}

func start(c *cli.Context) error {
	if c.String("cattleURL") == "" {
		return fmt.Errorf("Cattle URL required")
	}
	if c.String("cattleAccessKey") == "" {
		return fmt.Errorf("Cattle Access Key required")
	}
	if c.String("cattleSecretKey") == "" {
		return fmt.Errorf("Cattle Secret Key required")
	}
	a := agent.NewAgent(c.Duration("poll-interval"), c.String("cattleURL"), c.String("cattleAccessKey"), c.String("cattleSecretKey"))
	return a.Start()
}
