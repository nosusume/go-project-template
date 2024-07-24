package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	db     *sqlx.DB
	engine *gin.Engine
}

func NewApp(db *sqlx.DB, engine *gin.Engine) *App {
	return &App{
		db:     db,
		engine: engine,
	}
}

func (app App) GetEngine() *gin.Engine {
	return app.engine
}

func (app App) GetDB() *sqlx.DB {
	return app.db
}
