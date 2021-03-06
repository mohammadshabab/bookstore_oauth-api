package app

import (
	"github.com/gin-gonic/gin"

	"github.com/mohammadshabab/bookstore_oauth-api/src/http"
	"github.com/mohammadshabab/bookstore_oauth-api/src/repository/db"
	"github.com/mohammadshabab/bookstore_oauth-api/src/repository/rest"
	"github.com/mohammadshabab/bookstore_oauth-api/src/service/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// session, dbErr := cassandra.GetSession()
	// if dbErr != nil {
	// 	panic(dbErr)
	// }
	// session.Close()
	// dbRepository := db.NewRepository()
	// atService := access_token.NewService(dbRepository)
	//atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewAccessTokenHandler(access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
