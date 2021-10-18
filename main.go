package main

import (
	"fmt"
	"main/helper"
	"main/route"
	"main/db"
	_ "main/config"
	"time"
)


func main() {

	db, e := db.MySQL()
	helper.PanicIfError(e)
	eb := db.Ping()
	helper.PanicIfError(eb)
	fmt.Println("database running success")
	
	dt := time.Now()
	fmt.Println(dt.Format("02-Jan-2006"))
	// Routing
	r := route.NewRoute()
	r.Run() 
	

	
}

