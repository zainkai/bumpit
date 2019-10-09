package main

import (
	"fmt"

	"github.com/coreos/go-semver/semver"
	"github.com/zainkai/bumpit/pkg/utils"
)

func main() {
	v := semver.New("1.2.3")
	fmt.Println(v.String())
	fmt.Println(utils.GetFilesInDir("."))
}
