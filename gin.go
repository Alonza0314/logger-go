package loggergo

import (
	"strings"

	"github.com/Alonza0314/logger-go/v2/model"
)

// ginWriter adapts a model.LoggerInterface into an io.Writer, so it can be
// assigned to gin's gin.DefaultWriter / gin.DefaultErrorWriter. This file does
// not import gin-gonic/gin itself -- an io.Writer is all gin needs, so
// consumers of this library don't have to take on a gin dependency just to
// get a tagged logger.
//
// gin.Logger() (and gin.Default()) capture gin.DefaultWriter once, at the
// time the middleware is constructed -- so assign the writer before the
// gin.Engine is created:
//
//	gin.DefaultWriter = loggergo.NewGinWriter(logger.WithTag("GIN"))
//	gin.DefaultErrorWriter = loggergo.NewGinWriter(logger.WithTag("GIN"))
type ginWriter struct {
	log model.LoggerInterface
}

// NewGinWriter returns an io.Writer that forwards everything written to it
// into log via Infoln.
func NewGinWriter(log model.LoggerInterface) *ginWriter {
	return &ginWriter{log: log}
}

func (w *ginWriter) Write(p []byte) (int, error) {
	w.log.Infoln(strings.TrimRight(string(p), "\n"))
	return len(p), nil
}
