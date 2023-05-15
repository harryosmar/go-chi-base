package logger

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

var Logger = &log.Logger{
	Out:          os.Stdout,
	Formatter:    new(log.TextFormatter),
	Hooks:        make(log.LevelHooks),
	Level:        log.InfoLevel,
	ExitFunc:     os.Exit,
	ReportCaller: false,
}

func NewLogRequestEntry(logger *log.Logger, r *http.Request) *log.Entry {
	entry := log.NewEntry(logger)

	logFields := log.Fields{}

	logFields["req_id"] = middleware.GetReqID(r.Context())
	logFields["ts"] = time.Now().UTC().Format(time.RFC3339Nano)

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method

	logFields["remote_addr"] = r.RemoteAddr
	logFields["user_agent"] = r.UserAgent()

	payload, _ := httputil.DumpRequest(r, true)
	logFields["payload"] = fmt.Sprintf("%s", payload)

	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry = entry.Logger.WithFields(logFields)

	entry.Info("request started")

	return entry
}

type limitBuffer struct {
	*bytes.Buffer
	limit int
}

func (b limitBuffer) Write(p []byte) (n int, err error) {
	if b.Buffer.Len() >= b.limit {
		return len(p), nil
	}
	limit := b.limit
	if len(p) < limit {
		limit = len(p)
	}
	return b.Buffer.Write(p[:limit])
}

func (b limitBuffer) Read(p []byte) (n int, err error) {
	return b.Buffer.Read(p)
}

func NewLimitBuffer(size int) io.ReadWriter {
	return limitBuffer{
		Buffer: bytes.NewBuffer(make([]byte, 0, size)),
		limit:  size,
	}
}

func GetLogEntryFromCtx(ctx context.Context) *log.Entry {
	// Retrieve the log entry from the request context
	if logEntry, ok := ctx.Value("logEntry").(*log.Entry); ok {
		return logEntry
	}

	return log.NewEntry(Logger)
}
