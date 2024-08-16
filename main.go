package main

import (
	"context"
	"fmt"
	"os"
	"task/internal/api"
	"task/internal/config"
	"task/internal/database"
	"task/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"gopkg.in/yaml.v3"
)

func main() {

	config := initConfig("config.yaml")

	conn := database.Connect(&config)

	defer conn.Close(context.Background())

	initRouter(conn, &config)
}

func initConfig(path string) config.Config {

	fileConfig, err := os.Open(path)
	var config config.Config

	handlers.HasError(err, "Ошибка чтения файла конфигурации! %s", err)

	defer fileConfig.Close()

	decoder := yaml.NewDecoder(fileConfig)

	err = decoder.Decode(&config)

	handlers.HasError(err, "Ошибка десериализации файла конфигурации! %s", err)

	return config
}

/*
Инициализирует роутер и маршруты
*/
func initRouter(conn *pgx.Conn, cfg *config.Config) {

	router := gin.Default()

	router.GET("/api/v1/inventory/", api.GetInventoryWithOperations(conn))
	router.GET("/api/v1/inventory/all/", api.GetAllInventory(conn))
	router.GET("/api/v1/operations/", api.GetOperationWithInventory(conn))

	router.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
}
