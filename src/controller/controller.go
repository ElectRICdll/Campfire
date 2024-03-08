package controller

import (
	"campfire/entity"
	. "campfire/log"
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
	if _, ok := err.(entity.ExternalError); ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"res": RES_FAILURE,
			"e":   err.Error(),
		})
	}
	ctx.AbortWithStatus(http.StatusInternalServerError)
	Log.Error(err.Error())
	return
}
