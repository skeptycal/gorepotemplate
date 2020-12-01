package main

import (
	"fmt"

	"github.com/skeptycal/anansi"
	"github.com/skeptycal/gorepo"
)

func main() {

	command := "echo -e Hello world \n\tfrom golang"

	fmt.Println(gorepo.Shell(command))

	anansi.Sample()
}
