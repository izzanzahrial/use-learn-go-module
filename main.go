package main

// go mod init [link of the repository, example: github.com/izzanzahrial/use-learn-go-module] <-- to initialize module
// go get [name of the library, example: github.com/izzanzahrial/learn-go-module] <-- to get the module
// to upgrade version of the library you can change directly in go.mod or use go get again
// if there is a major upgrade, the solution is to change or create new module, rather than upgrade the version of the module
// check go.mod for major update

import (
	"fmt"

	learn_go_module "github.com/izzanzahrial/learn-go-module/v2"
)

func main() {
	fmt.Println(learn_go_module.TestGoModule("izzan"))
}
