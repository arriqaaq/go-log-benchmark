package main

import (
	"fmt"
	zerolog "github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		log.Print("hello")
	})
	http.HandleFunc("/logrus", func(w http.ResponseWriter, r *http.Request) {
		logrus.Print("hello")
	})
	http.HandleFunc("/fmt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("hello")
	})
	http.HandleFunc("/zerolog", func(w http.ResponseWriter, r *http.Request) {
		zerolog.Print("hello")
	})
	http.HandleFunc("/zap", func(w http.ResponseWriter, r *http.Request) {
		sugar.Infof("hello")
	})

	//10 prints
	http.HandleFunc("/log100", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 10; i++ {
			log.Print("hello")
		}
	})
	http.HandleFunc("/logrus100", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 10; i++ {
			logrus.Print("hello")
		}
	})
	http.HandleFunc("/fmt100", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 10; i++ {
			fmt.Print("hello")
		}
	})
	http.HandleFunc("/zerolog100", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 10; i++ {
			zerolog.Print("hello")
		}
	})
	http.HandleFunc("/zap100", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 10; i++ {
			sugar.Infof("hello")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
