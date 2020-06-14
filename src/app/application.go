package app

import (
	"github.com/edilbertloquine/go-microservices/oauth-api/src/app/http"
	"github.com/edilbertloquine/go-microservices/oauth-api/src/clients/cassandra"
	"github.com/edilbertloquine/go-microservices/oauth-api/src/domain/access_token"
	"github.com/edilbertloquine/go-microservices/oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication -
func StartApplication() {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	session.Close()

	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
