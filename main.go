package main

import (
	_ "github.com/damillano93/beego_yuno/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

func init(){
	databaseName, _ := beego.AppConfig.String("databaseName")
	databasePassword, _ := beego.AppConfig.String("databasePassword")
	databaseHost, _ := beego.AppConfig.String("databaseHost")
	databaseUserName, _ := beego.AppConfig.String("databaseUserName")
	databaseUrl := "postgres://"+databaseUserName+":"+databasePassword+"@"+databaseHost+"/"+databaseName+"?sslmode=disable&search_path=public"+""
	orm.RegisterDataBase("default", "postgres", databaseUrl)


}

func main() {
    //orm.Debug=true
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
        AllowHeaders: []string{"Origin", "x-requested-with",
            "content-type",
            "accept",
            "origin",
            "authorization",
            "x-csrftoken"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	beego.Run()
}

