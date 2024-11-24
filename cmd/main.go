package main

import (
	"log"
	"login-api/internal/config"
	"login-api/pkg/handlers"
	"login-api/pkg/repositories"
	"login-api/pkg/services"
	"github.com/gin-gonic/gin"
)

func init() {
	// Conectar ao banco de dados e sincronizar as migrações
	config.ConnectDatabase() // Assume que isso cria uma conexão com o banco
	config.SyncDatabase()    // Assume que isso aplica as migrações do banco
}

func main() {
	// Carregar a porta a partir da configuração
	port := config.LoadPort()

	// Criar o servidor Gin
	r := gin.Default()

	// Criar instâncias de repositório, serviço e handler
	db := config.DB // Supondo que config.GetDB() retorna a instância do banco
	userRepo := repositories.NewUserRepository(db) // Inicializa o repositório de usuários
	userService := services.NewUserService(userRepo) // Inicializa o serviço de usuários
	usersHandler := handlers.NewUserHandler(userService) // Inicializa o handler de usuários

	// Definir as rotas
	v1 := r.Group("/v1")
	{
		// Endpoint de ping para testar o servidor
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// Endpoint para criação de usuários
		v1.POST("/users", usersHandler.CreateUser)
	}

	// Iniciar o servidor
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run the server: ", err)
	}
}
