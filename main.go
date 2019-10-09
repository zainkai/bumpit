package main

import (
	"fmt"

	"github.com/coreos/go-semver/semver"
)

func main() {
	v := semver.New("1.2.3")
	fmt.Println(v.String())
}
