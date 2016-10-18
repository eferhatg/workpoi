package application

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	// "net/http"

	// "$GO_BOOTSTRAP_REPO_NAME/$GO_BOOTSTRAP_REPO_USER/$GO_BOOTSTRAP_PROJECT_NAME/handlers"
	// "$GO_BOOTSTRAP_REPO_NAME/$GO_BOOTSTRAP_REPO_USER/$GO_BOOTSTRAP_PROJECT_NAME/middlewares"
)

func New(config *viper.Viper) (*Application, error) {

	conns := config.GetString(config.GetString("env") + ".database.connection_string")
	log.Println(conns)
	db, err := sqlx.Connect("postgres", conns)
	if err != nil {
		log.Error(err)
	}

	cookieStoreSecret := config.GetString(config.GetString("env") + ".cookie.secret")

	app := &Application{}
	log.Println(app)
	app.config = config
	app.db = db
	app.sessionStore = sessions.NewCookieStore([]byte(cookieStoreSecret))

	return app, err
}

// Application is the application object that runs HTTP server.
type Application struct {
	config       *viper.Viper
	db           *sqlx.DB
	sessionStore sessions.Store
}

func (app *Application) TestApplication() {

	log.Println(app)
	log.Println("test")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		first := Venue{}
		err := app.db.Get(&first, "SELECT name,address FROM venue limit 1")
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%#v\n", first)
		c.JSON(200, first)
	})
	r.Run() // listen and server on 0.0.0.0:8080
}

type Venue struct {
	Name    string `db:"name"`
	Address string `db:"address"`
}
