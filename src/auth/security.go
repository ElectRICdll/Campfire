package auth

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

var SecurityInstance = SecurityGuard{
	dao.CampDaoContainer,
	dao.ProjectDaoContainer,
}

type SecurityGuard struct {
	campQuery dao.CampDao
	query     dao.ProjectDao
}

func (s SecurityGuard) IsUserHavingTitle(projID, userID uint) error {
	if project, found := ProjectCache.Get(fmt.Sprintf("%d", projID)); found {
		for _, value := range project.(*entity.Project).Members {
			if value.UserID == userID && len(value.Title) != 0 {
				return nil
			}
		}
		return util.NewExternalError("access denied")
	}

	project, err := dao.ProjectDao.ProjectInfo(dao.NewProjectDao(), projID)
	if project.ID != 0 {
		ProjectCache.Set(fmt.Sprintf("%d", projID), &project, cache.DefaultExpiration)
		for _, value := range project.Members {
			if value.UserID == userID && len(value.Title) != 0 {
				return nil
			}
		}
		return util.NewExternalError("access denied")
	}
	return err
}

func (s SecurityGuard) IsUserACampMember(campID, userID uint) error {
	if camp, found := CampCache.Get(fmt.Sprintf("%d", campID)); found {
		for _, value := range camp.(*entity.Camp).Members {
			if value.UserID == userID {
				return nil
			}
		}
		return util.NewExternalError("access denied")
	}

	camp, err := dao.ProjectDao.ProjectInfo(dao.NewProjectDao(), campID)
	if camp.ID != 0 {
		CampCache.Set(fmt.Sprintf("%d", campID), &camp, cache.DefaultExpiration)
		for _, value := range camp.Members {
			if value.UserID == userID {
				return nil
			}
		}
		return util.NewExternalError("access denied")
	}
	return err
}

func (s SecurityGuard) IsUserACampLeader(campID, userID uint) error {
	if camp, found := CampCache.Get(fmt.Sprintf("%d", campID)); found {
		if camp.(*entity.Camp).OwnerID == userID {
			return nil
		}
		return util.NewExternalError("access denied")
	}

	camp, err := dao.CampDao.CampInfo(dao.NewCampDao(), campID)
	if camp.ID != 0 {
		ProjectCache.Set(fmt.Sprintf("%d", campID), &camp, cache.DefaultExpiration)
		if camp.OwnerID == userID {
			return nil
		}
		return util.NewExternalError("access denied")
	}
	return err
}

func (s SecurityGuard) IsUserAProjMember(projID, userID uint) error {
	if project, found := ProjectCache.Get(fmt.Sprintf("%d", projID)); found {
		for _, value := range project.(*entity.Project).Members {
			if value.UserID == userID {
				return nil
			}
		}
		return util.NewExternalError("access denied")
	}

	project, err := dao.ProjectDao.ProjectInfo(dao.NewProjectDao(), projID)
	if project.ID != 0 {
		ProjectCache.Set(fmt.Sprintf("%d", projID), &project, cache.DefaultExpiration)
		for _, value := range project.Members {
			if value.UserID == userID {
				return nil
			}
		}
		return util.NewExternalError("access denied")
	}
	return err
}

func (s SecurityGuard) IsUserAProjLeader(projID, userID uint) error {
	if project, found := ProjectCache.Get(fmt.Sprintf("%d", projID)); found {
		if project.(*entity.Project).OwnerID == userID {
			return nil
		}
		return util.NewExternalError("access denied")
	}

	project, err := dao.ProjectDao.ProjectInfo(dao.NewProjectDao(), projID)
	if project.ID != 0 {
		ProjectCache.Set(fmt.Sprintf("%d", projID), &project, cache.DefaultExpiration)
		if project.OwnerID == userID {
			return nil
		}
		return util.NewExternalError("access denied")
	}
	return err
}

func (s SecurityGuard) IsUserATaskOwner(projID, taskID, userID uint) error {
	task, err := GetTaskFromCache(projID, taskID)
	if err != nil {
		task, err := s.query.TaskInfo(taskID)
		if err != nil {
			return err
		}
		if err := StoreTaskToProject(projID, task); err != nil {
			return err
		}
	}
	if task.OwnerID != userID {
		return util.NewExternalError("access denied")
	}
	return nil
}

func (s SecurityGuard) AuthMiddleware() gin.HandlerFunc {
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

func (s SecurityGuard) EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(([]byte)(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return (string)(hashedPassword), nil
}

func (s SecurityGuard) TokenGenerate(user entity.User) (string, error) {
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

func (s SecurityGuard) WSTokenVerify(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return util.CONFIG.SecretKey, nil
	})

	if err != nil || !token.Valid {
		return 0, util.NewExternalError("illegal token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, util.NewExternalError("illegal token")
	}

	return (uint)(claims["id"].(float64)), nil
}
