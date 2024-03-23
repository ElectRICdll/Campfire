package service

import (
	. "campfire/cache"
	"campfire/dao"
	"campfire/entity"
	"campfire/util"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"golang.org/x/crypto/bcrypt"
)

type SecurityService interface {
	AuthMiddleware() gin.HandlerFunc

	encryptPassword(password string) (string, error)

	tokenGenerate(entity.User) (string, error)

	IsUserACampMember(campID, userID uint) (bool, error)

	IsUserAProjMember(projID, userID uint) (bool, error)

	IsUserACampLeader(campID, userID uint) (bool, error)

	IsUserAProjLeader(projID, userID uint) (bool, error)

	IsUserHavingTitle(projID, userID uint) (bool, error)
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

func (s securityService) IsUserACampLeader(campID, userID uint) (bool, error) {
	if camp, found := CampCache.Get(fmt.Sprintf("%d", campID)); found {
		if camp.(entity.Camp).OwnerID == userID {
			return true, nil
		}
	}

	camp, err := dao.CampDao.CampInfo(dao.NewCampDao(), campID)
	if camp.ID != 0 {
		ProjectCache.Set(fmt.Sprintf("%d", campID), &camp, cache.DefaultExpiration)
		if camp.OwnerID == userID {
			return true, nil
		}
	}
	return false, err
}

func (s securityService) IsUserAProjMember(projID, userID uint) (bool, error) {
	if project, found := ProjectCache.Get(fmt.Sprintf("%d", projID)); found {
		for _, value := range project.(entity.Project).Members {
			if value.UserID == userID {
				return true, nil
			}
		}
	}

	project, err := dao.ProjectDao.ProjectInfo(dao.NewProjectDao(), projID)
	if project.ID != 0 {
		ProjectCache.Set(fmt.Sprintf("%d", projID), &project, cache.DefaultExpiration)
		for _, value := range project.Members {
			if value.UserID == userID {
				return true, nil
			}
		}
	}
	return false, err
}

func (s securityService) IsUserAProjLeader(projID, userID uint) (bool, error) {
	if project, found := ProjectCache.Get(fmt.Sprintf("%d", projID)); found {
		if project.(entity.Project).OwnerID == userID {
			return true, nil
		}
	}

	project, err := dao.ProjectDao.ProjectInfo(dao.NewProjectDao(), projID)
	if project.ID != 0 {
		ProjectCache.Set(fmt.Sprintf("%d", projID), &project, cache.DefaultExpiration)
		if project.OwnerID == userID {
			return true, nil
		}
	}
	return false, err
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
