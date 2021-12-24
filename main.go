package main

import (
	"context"
	"juriback2/controller"
	"juriback2/reposirory"
	"juriback2/service"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := http.NewServeMux()

	// assets
	assets := http.FileServer((http.Dir("assets")))
	r.Handle("/static/", http.StripPrefix("/static", assets))

	// REPO SERVICE CONTROLLER
	kontenRepo := reposirory.KontenRepository(conn())
	kontenService := service.KontenService(kontenRepo)
	kontenController := controller.KontenController(kontenService)

	r.HandleFunc("/", kontenController.Landing)
	r.HandleFunc("/id", kontenController.ById)
	r.HandleFunc("/tambah", kontenController.Tambah)
	r.HandleFunc("/ubah", kontenController.Ubah)
	r.HandleFunc("/view", kontenController.Views)

	// SERVER
	err := http.ListenAndServe(":8899", r)
	if err != nil {
		log.Println("Error 8899")
		log.Println(err.Error())
	}
}

func conn() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetMonitor(monitor))
	if err != nil {
		return nil
	}

	return client.Database("pebridb")
}
