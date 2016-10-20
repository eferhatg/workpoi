package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/eferhatg/workpoi/application"
	"github.com/spf13/viper"
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

	app.Start()

}
