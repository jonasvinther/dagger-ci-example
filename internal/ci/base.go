package ci

import (
	"context"
	"os"

	"dagger.io/dagger"
)

func NewDaggerClient(ctx context.Context) (*dagger.Client, error) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewBaseContainer(client *dagger.Client, src *dagger.Directory) *dagger.Container {
	container := client.Container().
		From("golang:1.20.5").
		WithWorkdir("/src").
		WithDirectory("/src", src).
		WithEnvVariable("CGO_ENABLED", "0").
		WithEnvVariable("GOOS", "linux").
		WithExec([]string{"go", "mod", "download"})

	return container
}

func LoadSource(client *dagger.Client, path string) *dagger.Directory {
	src := client.Host().Directory(path, dagger.HostDirectoryOpts{
		Exclude: []string{
			".git/",
		},
	})

	return src
}
