// +build generate

package main

import (
	"net/http"

	"github.com/shurcooL/vfsgen"
)

var fs http.FileSystem = http.Dir("/home/lukas/devel/pers/go-godin/godin/templates")

func main() {

	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:     "asset_fs.go",
		PackageName:  "godin",
		VariableName: "Templates",
	})
	if err != nil {
		panic(err)
	}
}
