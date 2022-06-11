package version

import (
	"errors"
	"io"
	"log"
	"regexp"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const (
	gitBranchMatchPattern = `^refs\/heads\/(?P<branch>[\S]*)$`
)

var (
	gitBranchMatch = regexp.MustCompile(gitBranchMatchPattern)
)

// Git is a struct that interacts with a git repo.
type Git struct {
	repo     *git.Repository
	branches map[plumbing.Hash][]*plumbing.Reference
	tags     map[plumbing.Hash]*plumbing.Reference
}

func (g *Git) mostRecentBranchReference() (*plumbing.Reference, error) {
	ref, err := g.repo.Head()
	if err != nil {
		return nil, err
	}

	com, err := g.repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	ci := object.NewCommitPreorderIter(com, nil, nil)
	for {
		com, err := ci.Next()
		if err != nil {
			return nil, err
		}
		log.Println(com.Hash)
	}
}

// Branches returns the current branch.
func (g *Git) Branches() ([]string, error) {
	h, err := g.repo.Head()
	if err != nil {
		return nil, err
	}

	if h.Name() != plumbing.HEAD {
		return []string{h.Name().Short()}, nil
	}

	branches := []string{}
	for _, ref := range g.branches[h.Hash()] {
		matches := gitBranchMatch.FindAllStringSubmatch(ref.Name().String(), -1)
		if matches != nil {
			for _, match := range matches {
				for i, name := range gitBranchMatch.SubexpNames() {
					if name == "branch" {
						branches = append(branches, match[i])
					}
				}
			}
		}
	}

	return branches, nil
}

// Commit returns the commit of the HEAD.
func (g *Git) Commit() (string, error) {
	h, err := g.repo.Head()
	if err != nil {
		return "", err
	}
	return h.String(), nil
}

// Committish returns the shortened committish of the HEAD.
func (g *Git) Committish() (string, error) {
	h, err := g.repo.Head()
	if err != nil {
		return "", err
	}
	return h.String()[:7], nil
}

// Tag returns the closest tag to the HEAD.
func (g *Git) Tag() (string, error) {
	h, err := g.repo.Head()
	if err != nil {
		return "", err
	}

	c, err := g.repo.CommitObject(h.Hash())
	if err != nil {
		return "", err
	}

	ci, err := g.repo.Log(&git.LogOptions{From: c.Hash})
	if err != nil {
		return "", err
	}

	var t string
	for {
		c, err := ci.Next()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return "", err
		}

		if ref, ok := g.tags[c.Hash]; ok {
			t = ref.Name().Short()
			break
		}
	}

	return t, nil
}

// NewGit returns a new Git reader.
func NewGit(path string) (*Git, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err
	}

	branches := map[plumbing.Hash][]*plumbing.Reference{}
	bi, err := repo.Branches()
	if err != nil {
		return nil, err
	}

	if err := bi.ForEach(func(r *plumbing.Reference) error {
		brs := branches[r.Hash()]
		branches[r.Hash()] = append(brs, r)
		return nil
	}); err != nil {
		return nil, err
	}

	tags := map[plumbing.Hash]*plumbing.Reference{}
	ti, err := repo.Tags()
	if err != nil {
		return nil, err
	}

	if err := ti.ForEach(func(r *plumbing.Reference) error {
		tags[r.Hash()] = r
		return nil
	}); err != nil {
		return nil, err
	}

	return &Git{
		repo:     repo,
		branches: branches,
		tags:     tags,
	}, nil
}
