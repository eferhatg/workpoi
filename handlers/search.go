package handlers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/eferhatg/workpoi/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type searchargs struct {
    Popregions []models.RegionRow
    Popvenues []models.VenueRow
    ImgRoot string
}


func GetSearch(c *gin.Context) {



}
