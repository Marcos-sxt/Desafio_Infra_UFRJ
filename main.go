package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

const (
	sessionDuration = time.Minute * 15
	adminUsername   = "admin"
	adminPassword   = "123456"
)

var redisClient *redis.Client
var ctx = context.Background()

// Função para criar a instância do aplicativo Fiber
func createApp() *fiber.App {
	app := fiber.New()

	// Middleware para registrar solicitações
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Received request: %s %s", c.Method(), c.Path())
		return c.Next()
	})

	// Rota inicial para exibir uma mensagem de "Hello, world!"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	// Rota para autenticação de login
	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		if username != adminUsername || password != adminPassword {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid credentials")
		}

		// Cria uma nova sessão usando uma chave aleatória
		sessionID := fmt.Sprintf("session:%s", generateRandomString(16))
		err := redisClient.Set(ctx, sessionID, "admin", sessionDuration).Err()
		if err != nil {
			log.Printf("[ERROR] - Error creating session: %s", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error creating session")
		}

		message := fmt.Sprintf("The new session is: %s", sessionID)
		return c.Status(fiber.StatusCreated).SendString(message)
	})

	// Rota para verificar a existência de uma sessão
	app.Get("/check/:session", func(c *fiber.Ctx) error {
		sessionID := c.Params("session")

		// Verifica se a sessão existe no Redis
		_, err := redisClient.Get(ctx, sessionID).Result()
		if err == redis.Nil {
			return c.Status(fiber.StatusNotFound).SendString("Session does not exist")
		} else if err != nil {
			log.Printf("[ERROR] - Error getting session: %s", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error getting session")
		}

		return c.SendString("The session exists")
	})

	// Rota para listar todas as sessões ativas
	app.Get("/list-sessions", func(c *fiber.Ctx) error {
		keys, err := getAllSessionKeys()
		if err != nil {
			log.Printf("[ERROR] - Error listing sessions: %s", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error listing sessions")
		}

		return c.JSON(keys)
	})

	return app
}

// Função para gerar uma string aleatória para o ID da sessão com um comprimento especificado
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Função para listar todas as chaves de sessão ativas
func getAllSessionKeys() ([]string, error) {
	var cursor uint64
	var keys []string
	var err error

	for {
		keys, cursor, err = redisClient.Scan(ctx, cursor, "session:*", 10).Result()
		if err != nil {
			return nil, err
		}

		if cursor == 0 {
			break
		}
	}

	return keys, nil
}

func main() {
	// Recupera o valor da variável de ambiente PORT ou use 3000 como padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Porta padrão se não estiver definida na variável de ambiente
	}

	// Configuração do cliente Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr: "infra-2023-cenario_de_desenvolvimento-redis-1:6379", // Endereço do servidor Redis
	})

	// Cria uma instância do aplicativo Fiber
	app := createApp()

	// Inicia o servidor na porta configurada
	err := app.Listen(":" + port)
	if err != nil {
		log.Panicf("[ERROR] - Error listening on port %s: %v", port, err)
	}
}
