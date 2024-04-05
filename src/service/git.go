package service

import (
	"campfire/auth"
	"campfire/dao"
	"campfire/entity"
	"campfire/log"
	"campfire/storage"
	"campfire/util"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/filemode"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/go-git/go-git/v5/storage/memory"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	RefPrefix = "refs/heads/"
)

type GitService interface {
	CreateRepo(path string) error

	Branches(projID uint, path string) ([]entity.Branch, error)

	CreateBranch(queryID uint, projID uint, branch string) error

	RemoveBranch(queryID uint, projID uint, branch string) error

	CommitFromWeb(queryID uint, projID uint, branch string, description string, files ...GitAction) error

	Clone(queryID uint, projID uint, branch string) ([]byte, error)

	Dir(queryID, projID uint, branch, path string) ([]storage.File, error)

	Read(queryID, projID uint, branch, filePath string) (string, error)
}

func NewGitService() GitService {
	return &gitService{
		access:    auth.SecurityInstance,
		projQuery: dao.ProjectDaoContainer,
		userQuery: dao.UserDaoContainer,
	}
}

type gitService struct {
	access    auth.SecurityGuard
	projQuery dao.ProjectDao
	userQuery dao.UserDao
	repo      *git.Repository
	mutex     sync.Mutex
}

func (g *gitService) CommitFromWeb(queryID uint, projID uint, branch string, description string, files ...GitAction) error {
	if err := g.access.IsUserAProjMember(queryID, projID); err != nil {
		return err
	}
	project, err := g.projQuery.ProjectInfo(projID, "Owner")
	user, err := g.userQuery.UserInfoByID(queryID)

	sig := &object.Signature{
		Name:  user.Name,
		Email: user.Email,
		When:  time.Now(),
	}

	mem := memfs.New()
	pusher := memory.NewStorage()
	push, err := git.Init(pusher, mem)
	if err != nil {
		return err
	}

	w, err := push.Worktree()
	if err != nil {
		return err
	}

	for _, file := range files {
		//mem.Create()
		switch file.Type {
		case "add":
			_, err = w.Add(project.Path + file.Path)
			if err != nil {
				return err
			}

		case "delete":
			_, err = w.Remove(project.Path + file.Path)
			if err != nil {
				return err
			}
		case "update":
			err := os.WriteFile(project.Path+file.Path, []byte(file.Content), 0644)
			if err != nil {
				return err
			}
			_, err = w.Add(project.Path + file.Path)
		}
	}

	_, err = w.Commit(description, &git.CommitOptions{
		Committer: sig,
	})
	if err != nil {
		return err
	}

	//// TODO: 远程推送暂时方法
	//storer := filesystem.NewStorage(osfs.New(project.Path, osfs.WithBoundOS()), nil)
	//repo, err := git.Open(storer, nil)
	//if err != nil {
	//	return err
	//}
	remoteURL := "http://localhost:" + util.CONFIG.Port + "/" + project.Path
	remote, err := push.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{remoteURL},
	})
	if err != nil {
		return err
	}
	err = remote.Push(&git.PushOptions{
		RefSpecs:   []config.RefSpec{(config.RefSpec)(RefPrefix + branch + ":" + RefPrefix + branch)},
		RemoteName: "origin",
	})
	if err == git.NoErrAlreadyUpToDate {
		return util.NewExternalError(err.Error())
	} else if err != nil {
		return err
	}

	err = push.Push(&git.PushOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (g *gitService) CreateRepo(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	storer := filesystem.NewStorage(
		osfs.New(path, osfs.WithBoundOS()),
		nil,
	)

	_, err := git.InitWithOptions(storer, nil, git.InitOptions{
		DefaultBranch: plumbing.NewBranchReferenceName("main"),
	})
	if err != nil {
		return err
	}

	return nil
}

func (g *gitService) Branches(projID uint, path string) ([]entity.Branch, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	refs, err := repo.References()
	if err != nil {
		return nil, err
	}

	var res []entity.Branch
	if err := refs.ForEach(func(reference *plumbing.Reference) error {
		var ref = entity.Branch{
			ProjID: projID,
		}
		if reference.Type() == plumbing.SymbolicReference {
			splits := strings.Split(reference.Target().String(), "/")
			ref.Name = splits[2]
			if ref.Name == "main" {
				ref.IsMain = true
			} else {
				ref.IsMain = false
			}
			res = append(res, ref)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return res, nil
}

func (g *gitService) CreateBranch(queryID uint, projID uint, branch string) error {
	if err := g.access.IsUserAProjMember(queryID, projID); err != nil {
		return err
	}
	project, err := g.projQuery.ProjectInfo(projID)
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.repo, err = git.PlainClone(project.Path, true, &git.CloneOptions{})
	defer g.closeRepo()

	head, err := g.repo.Head()
	if err != nil {
		return err
	}

	ref := plumbing.NewHashReference(plumbing.ReferenceName("refs/heads/"+branch), head.Hash())

	err = g.repo.Storer.SetReference(ref)

	err = g.repo.Storer.RemoveReference(ref.Name())
	return err
}

func (g *gitService) RemoveBranch(queryID uint, projID uint, branchName string) error {
	if err := g.access.IsUserAProjMember(queryID, projID); err != nil {
		return err
	}
	project, err := g.projQuery.ProjectInfo(projID)
	if err != nil {
		return err
	}
	g.repo, err = git.PlainClone(project.Path, true, &git.CloneOptions{})
	defer g.closeRepo()

	refs, err := g.repo.References()
	if err != nil {
		log.Errorf("无法获取仓库引用: %v", err)
	}

	var branchRef *plumbing.Reference
	if err := refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().IsBranch() && ref.Name().Short() == branchName {
			branchRef = ref
		}
		return nil
	}); err != nil {
		return err
	}

	if branchRef == nil {
		return util.NewExternalError("未找到分支")
	}

	err = g.repo.Storer.RemoveReference(branchRef.Name())
	return nil
}

func (g *gitService) Clone(queryID uint, projID uint, branch string) ([]byte, error) {
	if err := g.access.IsUserAProjMember(queryID, projID); err != nil {
		return nil, err
	}
	project, err := g.projQuery.ProjectInfo(projID)
	if err != nil {
		return nil, err
	}

	g.repo, err = git.PlainOpen(project.Path)
	defer g.closeRepo()
	if err != nil {
		return nil, err
	}

	w, err := g.repo.Worktree()
	if err != nil {
		return nil, err
	}
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
	})
	if err != nil {
		return nil, err
	}

	zipData, err := util.CreateZip(project.Path)
	if err != nil {
		return nil, err
	}

	return zipData, nil
}

func (g *gitService) Dir(queryID, projID uint, branch, path string) ([]storage.File, error) {
	if err := g.access.IsUserAProjMember(projID, queryID); err != nil {
		return nil, err
	}
	project, err := g.projQuery.ProjectInfo(projID)
	if err != nil {
		return nil, err
	}

	storer := filesystem.NewStorage(osfs.New(project.Path, osfs.WithBoundOS()), cache.NewObjectLRUDefault())

	repo, err := git.Open(storer, nil)

	if err != nil {
		return nil, err
	}

	ref, err := repo.Reference(plumbing.NewBranchReferenceName(branch), false)
	if err != nil {
		return nil, util.NewExternalError("branch " + branch + " not found")
	}
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}
	if path != "/" && path != "\\" && path != "" {
		tree, err = tree.Tree(path)
		if err != nil {
			return nil, err
		}
	}

	var files []storage.File
	for _, entry := range tree.Entries {
		if entry.Mode == filemode.Dir {
			if err != nil {
				return nil, err
			}
			files = append(files, storage.File{
				Name:        entry.Name,
				IsDirectory: true,
			})
		} else {
			files = append(files, storage.File{
				Name:        entry.Name,
				IsDirectory: false,
			})
		}
	}

	return files, nil
}

func (g *gitService) Read(queryID, projID uint, branch, filePath string) (string, error) {
	if err := g.access.IsUserAProjMember(projID, queryID); err != nil {
		return "", err
	}
	project, err := g.projQuery.ProjectInfo(projID)
	if err != nil {
		return "", err
	}

	storer := filesystem.NewStorage(osfs.New(project.Path, osfs.WithBoundOS()), cache.NewObjectLRUDefault())

	repo, err := git.Open(storer, nil)

	if err != nil {
		return "", err
	}

	ref, err := repo.Reference(plumbing.NewBranchReferenceName(branch), false)
	if err != nil {
		return "", util.NewExternalError("branch " + branch + " not found")
	}
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return "", err
	}

	tree, err := commit.Tree()
	if err != nil {
		return "", err
	}
	file, err := tree.File(filePath)
	if err != nil {
		return "", err
	}
	ok, err := file.IsBinary()
	if err != nil {
		return "", err
	}
	if ok {
		return "", util.NewExternalError("could not open binary")
	}

	content, err := file.Contents()

	return content, err
}

func (g *gitService) closeRepo() {
	g.repo = nil
}

//func (g *gitService) GitBackEnd(body, path string) ([]byte, error) {
//	path = filepath.Join(util.CONFIG.NativeStorageRootPath, path)
//	if _, err := os.Stat(path); os.IsNotExist(err) {
//		return err
//	}
//
//}

type GitAction struct {
	Type    string `json:"type"` // add, delete, update and merge
	Path    string `json:"path"`
	Content string `json:"content"`
}
