package git

import (
	"errors"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// FetchCommits fetches all commits after the specified SHA for the git repository
// in the path specified
func FetchCommits(path string, afterSha string) ([]Commit, error) {
	repo, repoErr := gogit.PlainOpen(path)
	if repoErr != nil {
		return nil, repoErr
	}

	head, headErr := repo.Head()
	if headErr != nil {
		return nil, headErr
	}

	iterator, logErr := repo.Log(&gogit.LogOptions{From: head.Hash()})
	if logErr != nil {
		return nil, logErr
	}

	commits := make([]Commit, 0)
	doneErr := errors.New("NOT AN ERROR")

	iterateErr := iterator.ForEach(func(commit *object.Commit) error {
		if commit.Hash.String() == afterSha {
			return doneErr
		}

		commits = append(commits, Commit{
			Message: commit.Message,
			SHA:     commit.Hash.String(),
		})

		return nil
	})

	if iterateErr != doneErr && iterateErr != nil {
		return nil, iterateErr
	}

	return commits, nil
}
