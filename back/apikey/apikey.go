package apikey

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

var apiKeys = make(map[string]bool)

// generateAPIKey генерирует новый API-ключ.
func generateAPIKey() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func apikey() {
	r := gin.Default()

	// Маршрут для генерации нового API-ключа
	r.POST("/generate", func(c *gin.Context) {
		apiKey, err := generateAPIKey()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate API key"})
			return
		}

		// Сохраняем API-ключ в памяти
		apiKeys[apiKey] = true

		c.JSON(http.StatusOK, gin.H{"api_key": apiKey})
	})

	// Пример защищенного маршрута, который требует API-ключ
	r.GET("/protected", func(c *gin.Context) {
		apiKey := c.GetHeader("API-Key")
		if _, exists := apiKeys[apiKey]; !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected route!"})
	})

	// Запуск сервера
	r.Run(":8080")
}
