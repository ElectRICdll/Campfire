package api

import (
	"campfire/auth"
	"campfire/dao"
	"campfire/service"
	"campfire/storage"
	"campfire/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	gitPattern = `\d+-\w+\.git`
)

type GitController interface {
	Clone(*gin.Context)

	Commit(*gin.Context)

	CreateBranch(*gin.Context)

	RemoveBranch(*gin.Context)

	RepoDir(*gin.Context)

	OpenFile(*gin.Context)

	GitHTTPBackend(*gin.Context)
}

func NewGitController() GitController {
	return gitController{
		service.NewGitService(),
		auth.SecurityInstance,
		dao.ProjectDaoContainer,
		&webdav.Handler{
			FileSystem: webdav.Dir("../repo"),
			LockSystem: webdav.NewMemLS(),
		},
	}
}

type gitController struct {
	gitService service.GitService
	sec        auth.SecurityGuard
	projQuery  dao.ProjectDao
	webdav     *webdav.Handler
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
	err := g.gitService.CommitFromWeb(userID, uri.PID, uri.Branch, body.Description, body.Actions...)
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
	filePath := ctx.Query("filename")
	data, err := g.gitService.Read(userID, uri.PID, uri.Branch, filePath)

	responseJSON(ctx, struct {
		Data string `json:"content"`
	}{data}, err)
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

func (g gitController) GitHTTPBackend(ctx *gin.Context) {
	uri := struct {
		GitPath string `uri:"gitPath"`
	}{}
	if err := ctx.BindUri(&uri); err != nil {
		responseError(ctx, err)
		return
	}

	repoPath := filepath.Join(util.CONFIG.NativeStorageRootPath, uri.GitPath)
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		responseError(ctx, err)
		return
	}

	projID, err := strconv.Atoi(strings.Split(uri.GitPath, "-")[0])
	if err != nil {
		responseError(ctx, err)
		return
	}

	_, err = g.projQuery.ProjectInfo((uint)(projID))
	if err != nil {
		responseError(ctx, err)
		return
	}

	g.webdav.ServeHTTP(ctx.Writer, ctx.Request)

	gitHTTPBackendPath := util.CONFIG.GitPath
	if !util.IsFileExists(gitHTTPBackendPath) {
		responseError(ctx, err)
		return
	}

	cmd := exec.Command(gitHTTPBackendPath)
	cmd.Dir = repoPath
	cmd.Stdout = ctx.Writer
	cmd.Stderr = os.Stderr
	cmd.Stdin = ctx.Request.Body

	if err := cmd.Run(); err != nil {
		responseError(ctx, err)
		return
	}
	responseSuccess(ctx)

	return
}
