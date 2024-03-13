package api

import (
	"campfire/log"
	"campfire/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	RES_SUCCESS string = "success"
	RES_FAILURE string = "failed"
)

func responseJSON(ctx *gin.Context, res interface{}, err error) {
	if err != nil {
		responseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"res":  RES_SUCCESS,
		"data": res,
	})
	return
}

func responseSuccess(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"res": RES_SUCCESS})
	return
}

func responseError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}
	if _, ok := err.(util.ExternalError); ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"res": RES_FAILURE,
			"e":   err.Error(),
		})
		return
	}
	ctx.AbortWithStatus(http.StatusInternalServerError)
	log.Error(err.Error())
	return
}
