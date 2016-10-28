package handlers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/eferhatg/workpoi/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type homeargs struct {
    Popregions []models.RegionRow
    Popvenues []models.VenueRow
    ImgRoot string
}


func GetHome(c *gin.Context) {

	db := c.MustGet("db").(*sqlx.DB)
	cnf := c.MustGet("config").(*viper.Viper)

	v := models.NewVenue(db)
	vr, err := v.Select(40, 40)
	if err != nil {
		log.Fatal(err)
	}
	
	s := homeargs{ Popvenues: vr, ImgRoot:cnf.GetString("image.root") }

	c.HTML(http.StatusOK, "home.tmpl", s)

}
