package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bitcoin-trading-system/bitcoin-slack-notification/config"
	"github.com/bitcoin-trading-system/bitcoin-slack-notification/handler"
)

func NewRouter(cfg config.Config) (*gin.Engine, error) {
	h, err := handler.NewHandler(cfg)
	if err != nil {
		return nil, err
	}


	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.POST("/message", h.PostMessage)

	return router, nil
}

func Run(r *gin.Engine, cfg config.Config) error {
	return r.Run(":" + cfg.BaseConfig.Port)
}
