package service

import (
	"campfire/entity"
	"campfire/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SecurityService interface {
	AuthMiddleware() gin.HandlerFunc

	encryptPassword(password string) string

	tokenGenerate(entity.User) (string, error)
}

func NewSecurityService() SecurityService {
	return securityService{}
}

type securityService struct{}

func (s securityService) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return util.CONFIG.SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("user", claims["name"])
		c.Set("id", claims["id"])

		c.Next()
	}
}

func (s securityService) encryptPassword(password string) string {
	return "crypto.RegisterHash()"
}

func (s securityService) tokenGenerate(user entity.User) (string, error) {
	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"exp":  util.CONFIG.AuthDuration.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(util.CONFIG.SecretKey)
	if err != nil {
		return "", err
	}

	return signed, nil
}
