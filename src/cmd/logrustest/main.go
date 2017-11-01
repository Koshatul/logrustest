package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {

	app := cli.NewApp()

	app.Name = "logrustest"
	app.Usage = "Logrus Test Example"
	app.Version = "0.1.0"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:   "test",
			Usage:  "Test Command",
			Action: testCommand,
		},
	}

	app.Run(os.Args)
}
