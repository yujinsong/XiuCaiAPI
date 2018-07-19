package main

import (
	"github.com/XiuCai/XiuCaiAPI/common"
	"github.com/XiuCai/XiuCaiAPI/utils"
	"github.com/XiuCai/XiuCaiAPI/utils/mysql-utils"
	"github.com/XiuCai/XiuCaiAPI/routers"
	"github.com/codegangsta/negroni"
	"net/http"
	"log"
	"fmt"
	"flag"
)

var ENV string

func main() {
	flag.StringVar(&ENV, "env", "prod", "env settings")
	flag.Parse()

	utils.CreateLogger("xiucai.log")
	ENV = "test"
	if ENV == "test" || ENV == "local" {
		mysql_utils.SetConfig(mysql_utils.LoadConfigFile("./mysql_json_local.conf"))
	} else {
		mysql_utils.SetConfig(mysql_utils.LoadConfigFile("./mysql_json_prod.conf"))
	}
	mysql_utils.GetInstance()
	common.StartUp()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr : ":8080",
		Handler: n,
	}

	log.Println("Listing...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
