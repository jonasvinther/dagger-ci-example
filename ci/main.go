package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jonasvinther/dagger-ci-example/internal/ci"
)

func main() {

	var (
		ctx = context.Background()
	)

	client, err := ci.NewDaggerClient(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	build := ci.Build(client)
	exitcode, err := build.ExitCode(ctx)
	if exitcode != 0 {
		fmt.Println(exitcode)
		fmt.Println(err)
	}

	test := ci.Test(client)
	exitcode, err = test.ExitCode(ctx)
	if exitcode != 0 {
		fmt.Println(exitcode)
		fmt.Println(err)
	}

}
