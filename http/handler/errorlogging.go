package handler

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/pufferpanel/apufferi/http"
	"github.com/pufferpanel/apufferi/logging"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				http.Respond(c).Fail().Status(500).Code(http.UNKNOWN).Message("unexpected error").Data(err).Send()
				logging.Errorf("Error handling route\n%+v\n%s", err, debug.Stack())
				c.Abort()
			}
		}()

		c.Next()
	}
}
