package main

import (
	"fmt"
	"net/http"

	"github.com/mikerybka/util"
)

func main() {
	root := util.RequireEnvVar("ROOT")
	port := util.RequireEnvVar("PORT")
	s := &util.FileServer{
		Root: root,
	}
	err := http.ListenAndServe(":"+port, s)
	fmt.Println(err)
}
