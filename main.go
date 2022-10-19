package main

import (
	"HTTPProxy/proxy"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	config, err := ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("read config.yaml error: %s", err)
	}

	err = config.Vaildation()
	if err != nil {
		log.Fatalf("config.yaml error: %s", err)
	}

	router := mux.NewRouter()
	for _, location := range config.Location {
		httpProxy, err := proxy.NewHTTPProxy(location.ProxyPass, location.BalanceMode)
		if err != nil {
			log.Fatalf("create proxy error: %s", err)
		}

		if config.HealthCheck {
			httpProxy.HealthCheck(config.HealthCheckInterval)
		}
		router.Handle(location.Pattern, httpProxy)
	}

	if config.MaxAllowed > 0 {
		router.Use(maxAllowedMiddleware(config.MaxAllowed))
	}

	svr := http.Server{
		Addr:    ":" + strconv.Itoa(config.Port),
		Handler: router,
	}

	// print config detail
	config.Print()

	// listen and serve
	if config.Schema == "http" {
		err := svr.ListenAndServe()
		if err != nil {
			log.Fatalf("listen and serve error: %s", err)
		}
	}
}
