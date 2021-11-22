package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bugg123/golang-microservices/product-api/data"
	"github.com/bugg123/golang-microservices/product-api/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	v := data.NewValidation()

	ph := handlers.NewProducts(l, v)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.Update)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.Create)
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	swaggerOpts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml", Path: "docs2"}
	swagger := middleware.SwaggerUI(swaggerOpts, nil)
	rapiOpts := middleware.RapiDocOpts{SpecURL: "/swagger.yaml", Path: "docs3"}
	rapi := middleware.RapiDoc(rapiOpts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/docs2", swagger)
	getRouter.Handle("/docs3", rapi)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
