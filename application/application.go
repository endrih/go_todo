package application

import (
	"database/sql"
	"endrih/go_todo/config"
	"endrih/go_todo/data"
	"log"
	"os"

	"github.com/gorilla/sessions"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	DB       *sql.DB
	Session  sessions.Store
	Config   *config.AppConfig
}

var App *Application = &Application{}

func (app *Application) Initialize() {
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.Config = config.Initialize()
	app.DB = data.Initialize(app.Config.DbConfig)
	app.initializeSessionStore()
}

func (app *Application) initializeSessionStore() {
	key := app.Config.SESSION_KEY // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30          // 30 days
	isProd := app.Config.IS_PROD

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	app.Session = store
}
