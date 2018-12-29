package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

var (
	router *chi.Mux
	conf   *viper.Viper
)

func main() {
	conf = viper.New()
	conf.SetConfigName("dev")       // no need to include file extension
	conf.AddConfigPath("../config") // set the path of your config file

	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}

	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	logger := setupLogger()
	handler := setupHandler(db, logger)
	router = chi.NewRouter()
	handler.Log.Info(fmt.Sprint("starting application..."))

	routers(handler)

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	http.ListenAndServe(":3000", router)
}
