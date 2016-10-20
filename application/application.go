package application

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"

	"github.com/eferhatg/workpoi/handlers"
	// "net/http"

	// "$GO_BOOTSTRAP_REPO_NAME/$GO_BOOTSTRAP_REPO_USER/$GO_BOOTSTRAP_PROJECT_NAME/handlers"
	// "$GO_BOOTSTRAP_REPO_NAME/$GO_BOOTSTRAP_REPO_USER/$GO_BOOTSTRAP_PROJECT_NAME/middlewares"
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

	router.GET("/", handlers.GetHome)

	port := app.config.GetString(app.config.GetString("env") + ".env.port")
	os.Setenv("PORT",port)


	log.Printf("Application Starting on %s", os.Getenv("PORT"))
	router.Run()

}

func (app *Application) appMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app", app)
		c.Next()
	}
}


