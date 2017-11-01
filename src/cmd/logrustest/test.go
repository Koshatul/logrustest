package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func testCommand(context *cli.Context) error {
	logrus.Debug("testCommand():start")
	logrus.Debug("testCommand():end")
	return nil
}
