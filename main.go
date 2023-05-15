package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/harryosmar/go-chi-base/app/user/factories"
	"github.com/harryosmar/go-chi-base/logger"
	"github.com/harryosmar/go-chi-base/middlewares"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func init() {
	// setup config
	viper.SetConfigFile("config.json") // Set the configuration file path
	viper.SetConfigType("json")        // Set the configuration file type
	viper.AutomaticEnv()               // Enable automatic environment variable binding
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// setup logger
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	middleware.DefaultLogger = middlewares.RequestLogger(logger.Logger)

	// setup metrics
	middlewares.InitPrometheusMetrics()
}

func main() {
	router := chi.NewRouter()

	// A good base middlewares stack
	router.Use(middlewares.CustomRequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.PrometheusMiddleware)
	router.Use(middlewares.ResponseSetHeaderRequestId)

	// Create a new Prometheus handler
	handler := promhttp.Handler()
	router.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	apiRouter()

	// Mount the sub-router to the main router
	router.Mount("/api/v1/user", apiRouter())

	log.Infof("Starting server on port %d", viper.GetInt("port"))
	err := http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("port")), router)
	if err != nil {
		log.Fatal(err)
	}
}

// A completely separate router for administrator routes
func apiRouter() http.Handler {
	userHandler := factories.MakeUserHandler()

	router := chi.NewRouter()
	router.Use(middlewares.ResponseSetContentTypeJSON)

	router.Post("/login", userHandler.Login)

	return router
}
