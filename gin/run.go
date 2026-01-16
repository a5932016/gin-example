package gin

import (
	"log"
	"net/http"

	"github.com/a5932016/gin_example/gin/services"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func Run() {
	webService := &http.Server{
		Addr:         ":8081",
		Handler:      services.RunWebService(),
		ReadTimeout:  0,
		WriteTimeout: 0,
	}

	g.Go(func() error {
		return webService.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
