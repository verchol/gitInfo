package gitinfo

import (
	"testing"
)

func TestGit(t *testing.T) {

	commits := GetGitInfo("../")
	if len(commits) == 0 {
		panic("commits array can't be empty")
	}
	PrintCommits(commits)
}
