package api

import (
	"campfire/service"
	"campfire/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GitController interface {
	Clone(*gin.Context)

	Commit(*gin.Context)

	CreateBranch(*gin.Context)

	RemoveBranch(*gin.Context)

	RepoDir(*gin.Context)

	OpenFile(*gin.Context)
}

func NewGitController() GitController {
	return gitController{
		service.NewGitService(),
	}
}

type gitController struct {
	gitService service.GitService
}

func (g gitController) Commit(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID    uint   `uri:"project_id" binding:"required"`
		Branch string `uri:"branch" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	body := struct {
		Description string              `json:"description"`
		Actions     []service.GitAction `json:"actions"`
	}{}
	if err := ctx.BindJSON(&body); err != nil {
		responseError(ctx, err)
		return
	}
	err := g.gitService.Commit(userID, uri.PID, uri.Branch, body.Description, body.Actions...)
	if err != nil {
		responseError(ctx, err)
		return
	}
	responseSuccess(ctx)
}

func (g gitController) CreateBranch(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID    uint   `uri:"project_id" binding:"required"`
		Branch string `uri:"branch" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	err := g.gitService.CreateBranch(userID, uri.PID, uri.Branch)
	if err != nil {
		responseError(ctx, err)
		return
	}
	responseSuccess(ctx)
}

func (g gitController) RemoveBranch(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID    uint   `uri:"project_id" binding:"required"`
		Branch string `uri:"branch" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	err := g.gitService.RemoveBranch(userID, uri.PID, uri.Branch)
	if err != nil {
		responseError(ctx, err)
		return
	}
	responseSuccess(ctx)
}

func (g gitController) OpenFile(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID    uint   `uri:"project_id" binding:"required"`
		Branch string `uri:"branch" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	filePath := ctx.Query("path")
	data, err := g.gitService.Read(userID, uri.PID, filePath)

	responseJSON(ctx, struct {
		Data []byte `json:"data"`
	}{data}, err)
	responseSuccess(ctx)
}

func (g gitController) Clone(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID    uint   `uri:"project_id" binding:"required"`
		Branch string `uri:"branch" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	zipData, err := g.gitService.Clone(userID, uri.PID, uri.Branch)
	if err != nil {
		responseError(ctx, err)
		return
	}
	ctx.Header("Content-Disposition", "attachment; filename=example.zip")
	ctx.Header("Content-Type", "application/zip")
	ctx.Data(http.StatusOK, "application/zip", zipData)
}

func (g gitController) RepoDir(ctx *gin.Context) {
	userID := (uint)(ctx.Keys["id"].(float64))
	uri := struct {
		PID    uint   `uri:"project_id" binding:"required"`
		Branch string `uri:"branch" binding:"required"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}
	path := ctx.Query("path")
	files, err := g.gitService.Dir(userID, uri.PID, uri.Branch, path)

	responseJSON(ctx, struct {
		Files []storage.File `json:"files"`
	}{files}, err)
}
