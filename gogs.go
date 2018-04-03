// +build go1.6

// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Gogs is a painless self-hosted Git Service.
package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/gogits/gogs/cmd"
	"github.com/gogits/gogs/pkg/setting"
)

<<<<<<< HEAD
const APP_VER = "0.11.36.0306"
=======
//const APP_VER = "0.11.34.1122"
const APP_VER = "10.0.0"
>>>>>>> 58cd90196b58bd706010d6ac0222ee36d36f93c9

func init() {
	setting.AppVer = APP_VER
}

func main() {
	app := cli.NewApp()
	app.Name = "Gogs"
	app.Usage = "A painless self-hosted Git service"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.Web,
		cmd.Serv,
		cmd.Hook,
		cmd.Cert,
		cmd.Admin,
		cmd.Import,
		cmd.Backup,
		cmd.Restore,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
