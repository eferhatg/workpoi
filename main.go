package main

import (
	
	//"database/sql"
	log "github.com/Sirupsen/logrus"

	"github.com/eferhatg/workpoi/application"
	"github.com/spf13/viper"
	//"log"
)

func newConfig() (*viper.Viper, error) {
	
	c := viper.New()
	c.SetConfigName("config")
	c.AddConfigPath(".")
	c.SetConfigType("toml")
	err := c.ReadInConfig()
	if err != nil {
		log.Println("Config file not found...")
		
	}

	return c, nil
}


func main() {

	config, err := newConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(config.GetString("env"))
	 app, err := application.New(config)
	 if err != nil {
	 	log.Fatal(err)
	 }

	app.TestApplication()

	// Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}

	// if you have null fields and use SELECT *, you must use sql.Null* in your struct
	//  places := []Place{}
	//  err = db.Select(&places, "SELECT * FROM place ORDER BY telcode ASC")

	// log.Println("You fuck up")
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {

	// 	first := Venue{}
	// 	err = db.Get(&first, "SELECT name,address FROM venue limit 1")
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// 	log.Printf("%#v\n", first)

	// 	c.JSON(200, first)
	// })
	// r.Run() // listen and server on 0.0.0.0:8080
}





