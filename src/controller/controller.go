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
		if _, ok := err.(entity.ExternalError); ok {
			responseBadRequest(ctx, err.Error())
			return
		}
		responseInternalError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"res":  RES_SUCCESS,
		"data": res,
	})
	return
}

func responseInternalError(ctx *gin.Context, err error) {
	ctx.AbortWithStatus(http.StatusInternalServerError)
	Log.Error(err.Error())
}

func responseBadRequest(ctx *gin.Context, errMessage string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"res": RES_FAILURE,
		"e":   errMessage,
	})
	return
}
