package handlers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/eferhatg/workpoi/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func GetHome(c *gin.Context) {

	db := c.MustGet("db").(*sqlx.DB)
	v := models.NewVenue(db)
	vr, err := v.Select(5,5)
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "home.tmpl", vr)

}
