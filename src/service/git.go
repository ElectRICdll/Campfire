package service

import (
	"campfire/auth"
	"campfire/dao"
	"campfire/log"
	"campfire/storage"
	"campfire/util"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/filemode"
	"io/ioutil"
	"os"
	"sync"
)

type GitService interface {
	CreateRepo(path string) error

	CreateBranch(queryID uint, projID uint, branch string) error

	RemoveBranch(queryID uint, projID uint, branch string) error

	Commit(queryID uint, projID uint, branch string, description string, files ...GitAction) error

	Clone(queryID uint, projID uint, branch string) ([]byte, error)

	Dir(queryID, projID uint, branch, path string) ([]storage.File, error)

	Read(queryID, projID uint, filePath string) ([]byte, error)
}

func NewGitService() GitService {
	return &gitService{
		access: auth.SecurityInstance,
		query:  dao.ProjectDaoContainer,
	}
}

type gitService struct {
	access auth.SecurityGuard
	query  dao.ProjectDao
	repo   *git.Repository
	mutex  sync.Mutex
}

func (g *gitService) Commit(queryID uint, projID uint, branch string, description string, files ...GitAction) error {
	if err := g.access.IsUserAProjMember(queryID, projID); err != nil {
		return err
	}
	project, err := g.query.ProjectInfo(projID)
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.repo, err = git.PlainOpen(project.Path)
	defer g.closeRepo()
	if err != nil {
		return err
	}

	w, err := g.repo.Worktree()
	if err != nil {
		return err
	}

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName("refs/heads/" + branch),
		Create: true,
	})
	if err != nil {
		return err
	}

	toRollBack := *w

	for _, file := range files {
		switch file.Type {
		case "add":
			_, err = toRollBack.Add(project.Path + file.Filepath)
			if err != nil {
				return err
			}

		case "delete":
			_, err = toRollBack.Remove(project.Path + file.Filepath)
			if err != nil {
				return err
			}
		case "update":
			err := os.WriteFile(project.Path+file.Filepath, []byte(file.Content), 0644)
			if err != nil {
				return err
			}
			_, err = toRollBack.Add(project.Path + file.Filepath)
		}
	}

	w = &toRollBack
	_, err = w.Commit(description, &git.CommitOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (g *gitService) CreateRepo(path string) error {
	_, err := git.PlainInit(path, true)
	if err != nil {
		return err
	}
	return nil
}

func (g *gitService) CreateBranch(queryID uint, projID uint, branch string) error {
	if err := g.access.IsUserAProjMember(queryID, projID); err != nil {
		return err
	}
	project, err := g.query.ProjectInfo(projID)
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
	project, err := g.query.ProjectInfo(projID)
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
	project, err := g.query.ProjectInfo(projID)
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
	project, err := g.query.ProjectInfo(projID)
	if err != nil {
		return nil, err
	}

	res, err := git.PlainOpen(project.Path + path)
	g.repo = res
	defer g.closeRepo()
	if err != nil {
		return nil, err
	}

	ref, err := res.Reference(plumbing.ReferenceName("/refs/heads/"+branch), true)
	if err != nil {
		return nil, err
	}

	commit, err := res.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}

	var files []storage.File
	for _, entry := range tree.Entries {
		file := storage.File{}
		if entry.Mode == filemode.Regular {
			file.Name = entry.Name
			file.IsDirectory = false
		} else if entry.Mode == filemode.Dir {
			file.Name = entry.Name
			file.IsDirectory = false
			//subTree, err := tree.Tree(entry.Name)
			//if err != nil {
			//	log.Fatalf("Error getting subtree: %v", err)
			//}
		}
		files = append(files, file)
	}

	return files, nil
}

func (g *gitService) Read(queryID, projID uint, filePath string) ([]byte, error) {
	if err := g.access.IsUserAProjMember(queryID, projID); err != nil {
		return nil, err
	}
	project, err := g.query.ProjectInfo(projID)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(project.Path + filePath)
	return content, err
}

func (g *gitService) closeRepo() {
	g.repo = nil
}

type GitAction struct {
	Type     string // add, delete, update and merge
	Filepath string
	Content  string
}
