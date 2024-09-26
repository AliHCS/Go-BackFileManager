package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"FileManager/src/config"
	"FileManager/src/domain/models" // Importa tu modelo de usuario

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // Asegúrate de importar este paquete
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifica el token JWT
func AuthMiddleware(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		claims, err := config.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		// Convertir el ID del claim a ObjectID
		userID, err := primitive.ObjectIDFromHex(claims.ID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID"})
			c.Abort()
			return
		}

		// Verificar si el usuario existe en la base de datos
		collection := client.Database(config.LoadEnv().MONGO_BD_NAME).Collection("users")
		var existingUser models.User
		err = collection.FindOne(c, bson.M{"_id": userID}).Decode(&existingUser)

		// Debug para ver el usuario existente
		fmt.Println("Existing User:", existingUser)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error querying the database"})
			c.Abort()
			return
		}

		// Aquí puedes guardar los claims en el contexto si lo necesitas
		c.Set("userID", claims.ID)
		c.Set("userEmail", claims.Email)

		c.Next()
	}
}
