package main

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {

	r, err := git.PlainClone("./mylcogofile.txt", false, &git.CloneOptions{
		URL: "https://github.com/Gopikanth/GRPC.git",
	})
	if err != nil {
		log.Fatal(err)
	}

	headRef, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	ref := plumbing.NewHashReference("refs/heads/branch", headRef.Hash())
	err = r.Storer.SetReference(ref)
	if err != nil {
		log.Fatal(err)
	}
	err = r.Storer.RemoveReference(ref.Name())
	if err != nil {
		log.Fatal(err)
	}
}
