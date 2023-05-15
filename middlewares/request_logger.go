package middlewares

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	logger2 "github.com/harryosmar/go-chi-base/logger"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func RequestLogger(logger *log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := logger2.NewLogRequestEntry(logger, r)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			buf := logger2.NewLimitBuffer(150)
			ww.Tee(buf)
			reqID := middleware.GetReqID(r.Context())
			t1 := time.Now()

			defer func() {
				logFields := log.Fields{}
				logFields["req_id"] = reqID
				logFields["resp_status"] = ww.Status()
				logFields["resp_elapsed_ms"] = float64(time.Since(t1).Nanoseconds()) / 1000000.0
				logFields["resp_bytes_written"] = ww.BytesWritten()
				logFields["resp_headers"] = ww.Header()

				respContent, err := ioutil.ReadAll(buf)
				if err == nil {
					logFields["resp_content"] = string(respContent)
				}

				//entry = entry.WithFields(logFields) // append
				entry = entry.Logger.WithFields(logFields) // replace
				//entry.Write(ww.Status(), ww.BytesWritten(), ww.Header(), time.Since(t1), nil)
				if ww.Status() >= 400 {
					entry.Error("request completed with err")
				} else {
					entry.Info("request completed")
				}
			}()

			// Set the log entry as a value in the request context
			ctx := context.WithValue(r.Context(), "logEntry", log.NewEntry(logger).WithField("req_id", reqID))

			next.ServeHTTP(ww, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
