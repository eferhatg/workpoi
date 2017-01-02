package application

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	//"html/template"
	"os"

	//"errors"
	"github.com/eferhatg/workpoi/handlers"
	//"github.com/gin-gonic/contrib/renders/multitemplate"
	"net/http"
)

type Application struct {
	config       *viper.Viper
	db           *sqlx.DB
	sessionStore sessions.Store
}

func New(config *viper.Viper) (*Application, error) {

	conns := config.GetString(config.GetString("env") + ".database.connection_string")

	db, err := sqlx.Connect("postgres", conns)
	if err != nil {
		log.Error(err)
	}

	cookieStoreSecret := config.GetString(config.GetString("env") + ".cookie.secret")

	app := &Application{}

	app.config = config
	app.db = db
	app.sessionStore = sessions.NewCookieStore([]byte(cookieStoreSecret))

	return app, err
}

func (app *Application) Start() {

	router := gin.Default()
	router.Use(app.appMiddleware())
	router.Use(app.errorHandling())
	router.LoadHTMLGlob("templates/*")
	router.Static("/css", "./public/css")
	router.Static("/img", "./public/img")
	router.Static("/js", "./public/js")
	router.Static("/media", "./public/media")
	router.Static("/images", "./public/images")
	router.Static("/fonts", "./public/fonts")
	
	router.GET("/", handlers.GetHome)
	router.GET("/search", handlers.GetSearch)
	router.GET("/calisilacak-mekanlar", handlers.GetVenues)

	port := app.config.GetString(app.config.GetString("env") + ".env.port")
	os.Setenv("PORT", port)

	// router.NoRoute(func(c *gin.Context) {
	// 	c.Error(errors.New("Route not found"))
	// 	//c.HTML(http.StatusNotFound, "error.tmpl",nil)
	// })

	log.Printf("Application Starting on %s", os.Getenv("PORT"))
	router.Run()

}

// func createTemplateRender() multitemplate.Render {
// 	r := multitemplate.New()
// 	r.AddFromFiles("index", "./templates/base.tmpl", "./templates/home.tmpl")
// 	return r
// }

func (app *Application) appMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", app.db)
		c.Set("session", app.sessionStore)
		c.Set("config", app.config)

		c.Next()
	}
}

func (app *Application) errorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors.Errors()) > 0 {
			c.HTML(http.StatusInternalServerError, "error.tmpl", nil)
		}
	}
}
