package main

import (
	"fmt"
	"main/config"
	"main/db"
	"main/helper"
	"main/route"
	"time"
)

func main() {
	// initialize config
	config.Init()

	db, e := db.MySQL()
	helper.PanicIfError(e)
	eb := db.Ping()
	helper.PanicIfError(eb)
	fmt.Println("database running success")

	dt := time.Now()
	fmt.Println(dt.Format("02-Jan-2006"))
	// Routing
	r := route.NewRoute()

	port := fmt.Sprintf(`:%s`, config.Configuration.AppPort)
	r.Run(port)

}
