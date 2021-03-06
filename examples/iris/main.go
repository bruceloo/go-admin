package main

import (
	"github.com/kataras/iris"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/adapter"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/engine"
)

func main() {
	app := iris.Default()

	eng := engine.Default()

	cfg := config.Config{
		DATABASE: []config.Database{
			{
				IP:           "127.0.0.1",
				PORT:         "3306",
				USER:         "root",
				PWD:          "root",
				NAME:         "godmin",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
		},
		AUTH_DOMAIN:  "localhost",
		LANGUAGE:     "cn",
		ADMIN_PREFIX: "admin",
	}

	adminPlugin := admin.NewAdmin(datamodel.TableFuncConfig)

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin).AddAdapter(new(adapter.Iris)).Use(app); err != nil {
		panic(err)
	}

	app.Run(iris.Addr(":8099"))
}
