package application

import (
	"github.com/spf13/viper"
	"testing"
)

func TestNew(t *testing.T) {

	c := viper.New()
	c.SetConfigName("config")
	c.AddConfigPath("../")
	c.SetConfigType("toml")
	err := c.ReadInConfig()
	if err != nil {
		t.Errorf("Config file not found...")

	}
	app, err := New(c)
	if err != nil {
		t.Fatalf("Cant create new application: %q", err)
	}

	if app.db == nil {
		t.Error("Application db object couldn't be created")
	}

	if app.sessionStore == nil {
		t.Error("Application session object couldn't be created")
	}

}
