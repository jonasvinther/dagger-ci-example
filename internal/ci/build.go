package ci

import "dagger.io/dagger"

func Build(client *dagger.Client) *dagger.Container {
	src := LoadSource(client, "./")

	container := NewBaseContainer(client.Pipeline("build"), src)

	container = container.WithExec([]string{"go", "build", "-o", "example"}).
		WithExec([]string{"./example"})

	return container
}
