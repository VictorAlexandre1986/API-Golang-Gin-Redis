package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func init() {
	// Configuração do cliente Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Endereço padrão do Redis
	})

	// Verificar conexão
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func main() {
	router := gin.Default()

	// Rotas
	router.POST("/set", setKey)
	router.GET("/get/:key", getKey)

	// Iniciar o servidor
	router.Run(":8080")
}

func setKey(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func getKey(c *gin.Context) {
	key := c.Param("key")

	value, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}
