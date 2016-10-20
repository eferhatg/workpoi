package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/eferhatg/workpoi/models"
	"github.com/jmoiron/sqlx"
)

func GetHome(c *gin.Context) {

	db:=c.MustGet("db").(*sqlx.DB)
	v:=models.NewVenue(db)
	vr,_:=v.GetOne()
	c.HTML(http.StatusOK, "home.html.tmpl", vr)
	
}
