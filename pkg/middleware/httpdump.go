package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http/httputil"
)

func RequestDump(verbose bool, logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if verbose {
			body, err := httputil.DumpRequest(c.Request, true)
			if err != nil {
				logger.Err(err).Msg("DUMP REQUEST FAILED")
			} else {
				logger.Info().Msgf("DUMP REQUEST SUCCESS \n", body)
			}
		}
		c.Next()
	}
}
