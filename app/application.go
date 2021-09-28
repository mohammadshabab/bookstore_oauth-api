package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammadshabab/bookstore_oauth-api/http"
	"github.com/mohammadshabab/bookstore_oauth-api/repository/db"
	"github.com/mohammadshabab/bookstore_oauth-api/src/domain/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// dbRepository := db.NewRepository()
	// atService := access_token.NewService(dbRepository)
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
