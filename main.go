package main

import (
	"fmt"
	"main/config"
	"main/helper"
	"main/route"
	"time"
)

type formB struct {
  Username string `json:"username"`
  Password string `json:"password"`
}


func main() {

	db, e := config.MySQL()
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

