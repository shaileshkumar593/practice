package middleware

import (
	"net/http"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func Logging(next http.Handler, logger kitlog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		_ = level.Info(logger).Log(
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
		)
	})
}
