

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func main() {

	r, err := git.PlainClone("./mylcogofile.txt", false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "Gopikanth",
			Password: "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", //authentication using access token
		},
		URL:      "https://github.com/Gopikanth/go_git_authentication.git",
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}
	//file input and output
	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		log.Fatal(err)
	}
	length, err := io.WriteString(file, r.Storer.NewEncodedObject().Hash().String())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}
func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
