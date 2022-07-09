package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/serjbibox/rest-api-template/handlers"
)

func (s *Server) InitRoutes() {
	c := gin.LoggerConfig{
		Output:    os.Stdout,
		SkipPaths: []string{"/scipped-path"},
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				params.ClientIP,
				params.TimeStamp.Format(time.RFC1123),
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)
		},
	}
	s.router.Use(gin.LoggerWithConfig(c))
	s.router.Use(gin.Recovery())
	rg := s.router.Group("/user")
	{
		rg.Use(auth())
		rg.GET("/:id", handlers.GetUser)
	}
	s.router.GET("/test", func(c *gin.Context) {
		c.String(200, "test")
	})
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println("c.GetHeader()", c.GetHeader("Authorization"))
		if len(authHeader) == 0 {
			c.String(http.StatusBadRequest, errors.New("authorization is required Header").Error())
			c.Abort()
		}
		if authHeader != "apiKey" {
			c.String(http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to this operation: api_key=%s", authHeader).Error())
			c.Abort()
		}
		c.Next()
	}
}
