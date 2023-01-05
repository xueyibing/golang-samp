package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"os"

	mysql_samp "wxcloudrun-golang/databases"
	"wxcloudrun-golang/router"

)

func main() {
	//if err := db.Init(); err != nil {
	//	panic(fmt.Sprintf("mysql init failed with %+v", err))
	//}

	source := "%s:%s@tcp(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	user := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")
	addr := os.Getenv("MYSQL_ADDRESS")
	dataBase := os.Getenv("MYSQL_DATABASE")
	if dataBase == "" {
		dataBase = "golang_demo"
	}
	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	fmt.Println("start init mysql with ", source)
	err := mysql_samp.Init(
		addr,
		user,
		pwd,
		dataBase)
	if err != nil {
		panic(err)
	}


	r := router.InitRouter()
	r.Run(":80" )


}
