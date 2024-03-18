package service

import (
	"campfire/dao"
	"campfire/entity"
	"campfire/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type SecurityService interface {
	AuthMiddleware() gin.HandlerFunc

	encryptPassword(password string) (string, error)

	tokenGenerate(entity.User) (string, error)

	IsUserACampMember(campID, userID uint) (bool, error)

	IsUserAProjMember(projID, userID uint) (bool, error)
}

func NewSecurityService() SecurityService {
	return securityService{}
}

type securityService struct {
	campQuery dao.CampDao
	query     dao.ProjectDao
}

func (s securityService) IsUserACampMember(campID, userID uint) (bool, error) {
	// TODO

	res, err := s.campQuery.IsUserACampMember(campID, userID)
	if err != nil {
		return false, err
	}

	return res, nil
}

func (s securityService) IsUserAProjMember(projID, userID uint) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s securityService) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return util.CONFIG.SecretKey, nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", claims["name"])
		ctx.Set("id", claims["id"])

		ctx.Next()
	}
}

func (s securityService) encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(([]byte)(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return (string)(hashedPassword), nil
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
