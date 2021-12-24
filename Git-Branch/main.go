package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {

	r, err := git.PlainClone("./mylcogofile.txt", false, &git.CloneOptions{
		URL:           "https://github.com/Gopikanth/GRPC.git",
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/branch")),
		SingleBranch:  true,
	})

	if err != nil {
		log.Fatal(err)
	}

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()

	if err != nil {
		log.Fatal(err)
	}

	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(commit)
}
  