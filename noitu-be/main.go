package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/khalt00/noitu/internal/controller"
	"github.com/khalt00/noitu/internal/dict"
	"github.com/khalt00/noitu/internal/ws"
	"github.com/khalt00/noitu/pkg/config"
	"github.com/lmittmann/tint"
)

func main() {
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("invalid configuration", err)
	}
	engine := gin.Default()
	programLevel := new(slog.LevelVar)

	programLevel.Set(slog.LevelDebug)
	w := os.Stderr

	// create a new logger
	logger := slog.New(tint.NewHandler(w, &tint.Options{
		Level:     programLevel,
		AddSource: true,
	}))
	slog.SetDefault(logger)

	dict.InitDict()

	gm := ws.NewGameManager()
	hub := ws.NewHub(gm)
	go hub.Run()

	controller := controller.NewGameController(hub, gm)

	engine.GET("/api/v1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/ws", controller.Register)

	slog.Info("Starting server", "port", config.PORT)
	engine.Run(fmt.Sprintf(":%d", config.PORT))
}
