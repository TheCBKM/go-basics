package main

import "github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)

func main() {
	repo, err := git.PlainOpen("./repository")
	commit, err := repo.CommitObject(ref.Hash())
	tree, err := commit.Tree()
	var files []string
	tree.Files().ForEach(func(f *object.File) error {
		files = append(files, f.Name)
		return nil
	})
}
