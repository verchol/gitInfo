package gitinfo

import (
	"fmt"
	"path"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// Example how to resolve a revision into its commit counterpart
func GetGitInfo(path string) []*object.Commit {
	//CheckArgs("<path>", "<revision>")

	//revision := os.Args[2]

	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	remotes, err := r.Remotes()
	PrintArray(remotes)
	//CheckIfError(err)
	ref, err := r.Head()
	// Resolve revision into a sha1 commit, only some revisions are resolved
	// look at the doc to get more details
	//Info("git rev-parse %s", revision)
	commits := []*object.Commit{}
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	cIter.ForEach(func(c *object.Commit) error {
		commits = append(commits, c)
		//PrintCommits(commits)
		return nil
	})

	return commits

}
func PrintArray(array []*git.Remote) {
	fmt.Printf("Print Array, %v", len(array))
	defer fmt.Println("After PrintArray")
	for _, a := range array {
		fmt.Println(a)
		_, file := path.Split(a.String())
		fmt.Printf("path : %s", file)

	}
}
func PrintCommits(commits []*object.Commit) {
	for _, commit := range commits {
		fmt.Println(commit)
	}
}
func main() {
	commits := GetGitInfo("./")
	if commits == nil {
		panic("can't be empty")
	}
}
