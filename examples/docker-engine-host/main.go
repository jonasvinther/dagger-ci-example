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

	// Running docker commands in Dagger using a docker socket
	socket := client.Container().
		From("docker").
		WithUnixSocket("/var/run/docker.sock", client.Host().UnixSocket("/var/run/docker.sock")).
		WithExec([]string{"docker", "ps"})
	out, err := socket.Stdout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
