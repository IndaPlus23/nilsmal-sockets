package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/datasektionen/facebook/server/actions"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	r.GET("/blog_feed", actions.GetBlogFeed)
	r.GET("/send_websocket", actions.SendWebSocketMessageToClients)
}

func InitWebsocket(r *gin.RouterGroup) {
	r.GET("", actions.SendWebsocket)
}